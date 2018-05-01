package messenger

import (
	"context"
	//"errors"
  _ "fmt"
  "os"
	"time"

  "encoding/json"

	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
  "github.com/joho/godotenv"
  "github.com/tidwall/gjson"
)

/*
{
    "envelope":{
        "name":"CreateMeetingReqMsg",
        "routing":{
            "sender":"bbb-web"
        }
    },
    "core":{
        "header":{
            "name":"CreateMeetingReqMsg"
        },
        "body":{
          .......
        }
    }
}
*/

type Message struct {
    Envelope MessageEnvelope `json:"envelope"`
	  Core MessageCore `json:"core"`
}

type MessageEnvelope struct {
    Name string `json:"name"`
    Routing MessageEnvelopeRouting `json:"routing"`
}

type MessageEnvelopeRouting struct {
    Sender string `json:"sender"`
}

type MessageCore struct {
    Header MessageCoreHeader `json:"header"`
    Body MessageCoreBody `json:"body"`
}

type MessageCoreHeader struct {
    Name string `json:"name"`
}

type MessageCoreBody struct {
    Props interface {} `json:"props"`
}

type MessageCoreBodyPropsMeetingProp struct {
    Name string `json:"name"`
    ExtId string `json:"extId"`
    IntId string `json:"intId"`
    IsBreakout string `json:"isBreakout"`
}

var messageMapping = make(map[string]string)

func init() {
    err := godotenv.Load()
    if err != nil {
        beego.Error("Error loading .env file")
    }
    messageMapping["CreateMeetingReqMsg"] = "MeetingCreatedEvtMsg"
    //messageMapping["CreateMeetingReqMsg"] = "CheckAlivePongSysMsg"
}

func SendMessage(name, props string) {
    messageName := name
    var result map[string]interface{}
    json.Unmarshal([]byte(props), &result)
    messageBodyProps := result["props"].(map[string]interface{})

    message := Message {
        Envelope: MessageEnvelope {
            Name: messageName,
            Routing: MessageEnvelopeRouting {
                Sender: "bbb-api-meetings",
            },
        },
        Core: MessageCore {
            Header: MessageCoreHeader {
                Name: messageName,
            },
            Body: MessageCoreBody {
                Props: messageBodyProps,
            },
        },
    }

    // This could be global
    redisServerAddr := os.Getenv("REDIS_SERVER_ADDRESS")

  	ctx, cancel := context.WithCancel(context.Background())

    go func() {
        err := listenPubSubChannels(ctx,
          redisServerAddr,
          func() error {
            beego.Info("Listener started push the message...")
            publish(message)
            return nil
          },
          func(channel string, message []byte) error {
            // This listener runs forever and triggers the response to the requester
            // but for the moment we just see what is going on
            beego.Info("Channel: ", channel, ", Message: ", string(message[:]))

            m, ok := gjson.Parse(string(message)).Value().(map[string]interface{})
            if !ok {
              beego.Error("Error")
            }
            mEnvelope := m["envelope"].(map[string]interface{})

            if mEnvelope["name"] != messageMapping[messageName] {
              // Not the message we are looking for, continue
              return nil
            }

            mCore := m["core"].(map[string]interface{})
            mCoreBody := mCore["body"].(map[string]interface{})
            mCoreBodyProps := mCoreBody["props"].(map[string]interface{})
            if mCoreBodyProps == nil {
              // It doesn't have props, continue
              return nil
            }
            mCoreBodyPropsMeetingProp := mCoreBodyProps["meetingProp"].(map[string]interface{})
            if mCoreBodyPropsMeetingProp == nil {
              // It doesn't have meetingProp, continue
              return nil
            }

            if mCoreBodyPropsMeetingProp["intId"] != "ec9d5087d9a52497ddfcb5c73fc6d1d4328547b6-1524768169803" {
              // The intId doesn't match, continue
              return nil
            }

            //all right, we found it, return to the requester
            beego.Critical(mCoreBodyProps)
            beego.Debug(">>>>>>>>>>>>>>>>>>>>> Done")
            cancel()
            return nil
          },
          "from-akka-apps-redis-channel")

        if err != nil {
          beego.Error(err)
          return
        }
    }()
}

// listenPubSubChannels listens for messages on Redis pubsub channels. The
// onStart function is called after the channels are subscribed. The onMessage
// function is called for each message.
func listenPubSubChannels(ctx context.Context, redisServerAddr string,
    onStart func() error,
    onMessage func(channel string, data []byte) error,
    channels ...string) error {
    // A ping is set to the server with this period to test for the health of
    // the connection and server.
    const healthCheckPeriod = time.Minute

    c, err := redis.Dial("tcp", redisServerAddr,
        // Read timeout on server should be greater than ping period.
        redis.DialReadTimeout(healthCheckPeriod+10*time.Second),
        redis.DialWriteTimeout(10*time.Second))
    if err != nil {
        return err
    }
    defer c.Close()

    psc := redis.PubSubConn{Conn: c}

    if err := psc.Subscribe(redis.Args{}.AddFlat(channels)...); err != nil {
        return err
    }

    done := make(chan error, 1)

    // Start a goroutine to receive notifications from the server.
    go func() {
        for {
            switch n := psc.Receive().(type) {
            case error:
                done <- n
                return
            case redis.Message:
                if err := onMessage(n.Channel, n.Data); err != nil {
                    done <- err
                    return
                }
            case redis.Subscription:
                switch n.Count {
                case len(channels):
                    // Notify application when all channels are subscribed.
                    if err := onStart(); err != nil {
                        done <- err
                        return
                    }
                case 0:
                    // Return from the goroutine when all channels are unsubscribed.
                    done <- nil
                    return
                }
            }
        }
    }()

    ticker := time.NewTicker(healthCheckPeriod)
    defer ticker.Stop()
loop:
    for err == nil {
        select {
        case <-ticker.C:
            // Send ping to test health of connection and server. If
            // corresponding pong is not received, then receive on the
            // connection will timeout and the receive goroutine will exit.
            if err = psc.Ping(""); err != nil {
                break loop
            }
        case <-ctx.Done():
            break loop
        case err := <-done:
            // Return error from the receive goroutine.
            return err
        }
    }

    // Signal the receiving goroutine to exit by unsubscribing from all channels.
    psc.Unsubscribe()

    // Wait for goroutine to complete.
    return <-done
}

func publish(message Message) {
    b, err := json.Marshal(&message)
    if err != nil {
        beego.Debug(err)
        return
    }
    beego.Debug(string(b))

    // This could be global
    redisServerAddr := os.Getenv("REDIS_SERVER_ADDRESS")

		c, err := redis.Dial("tcp", redisServerAddr)
    if err != nil {
        beego.Error(err)
        return
    }
    defer c.Close()

    c.Do("PUBLISH", "to-akka-apps-redis-channel", string(b))
}
