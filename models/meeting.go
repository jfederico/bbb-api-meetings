package models

import (
	"bbb-api-meetings/lib/messenger"
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	//"github.com/astaxie/beego/logs"
)


const CREATE_PROPS = "" +
"{" + "\n" +
"   \"props\":{" + "\n" +
"      \"meetingProp\":{" + "\n" +
"         \"name\":\"random-773640\"," + "\n" +
"         \"extId\":\"random-773640\"," + "\n" +
"         \"intId\":\"ec9d5087d9a52497ddfcb5c73fc6d1d4328547b6-1524768169803\"," + "\n" +
"         \"isBreakout\":false" + "\n" +
"      }," + "\n" +
"      \"breakoutProps\":{" + "\n" +
"         \"parentId\":\"bbb-none\"," + "\n" +
"         \"sequence\":0," + "\n" +
"         \"breakoutRooms\":[" + "\n" +
"         ]" + "\n" +
"      }," + "\n" +
"      \"durationProps\":{" + "\n" +
"         \"duration\":0," + "\n" +
"         \"createdTime\":1524768169803," + "\n" +
"         \"createdDate\":\"Thu Apr 26 18:42:49 UTC 2018\"," + "\n" +
"         \"maxInactivityTimeoutMinutes\":120," + "\n" +
"         \"warnMinutesBeforeMax\":5," + "\n" +
"         \"meetingExpireIfNoUserJoinedInMinutes\":5," + "\n" +
"         \"meetingExpireWhenLastUserLeftInMinutes\":1" + "\n" +
"      }," + "\n" +
"      \"password\":{" + "\n" +
"         \"moderatorPass\":\"mp\"," + "\n" +
"         \"viewerPass\":\"ap\"" + "\n" +
"      }," + "\n" +
"      \"recordProp\":{" + "\n" +
"         \"record\":false," + "\n" +
"         \"autoStartRecording\":false," + "\n" +
"         \"allowStartStopRecording\":true" + "\n" +
"      }," + "\n" +
"      \"welcomeProp\":{" + "\n" +
"         \"welcomeMsgTemplate\":\"<br>Welcome to <b>\"," + "\n" +
"         \"welcomeMsg\":\"<br>Welcome to <b>random-773640</b>!\"," + "\n" +
"         \"modOnlyMessage\":\"\"" + "\n" +
"      }," + "\n" +
"      \"voiceProp\":{" + "\n" +
"         \"telVoice\":\"79922\"," + "\n" +
"         \"voiceConf\":\"79922\"," + "\n" +
"         \"dialNumber\":\"613-555-1234\"," + "\n" +
"         \"muteOnStart\":false" + "\n" +
"      }," + "\n" +
"      \"usersProp\":{" + "\n" +
"         \"maxUsers\":0," + "\n" +
"         \"webcamsOnlyForModerator\":false," + "\n" +
"         \"guestPolicy\":\"ASK_MODERATOR\"" + "\n" +
"      }," + "\n" +
"      \"metadataProp\":{" + "\n" +
"         \"metadata\":{" + "\n" +
"         }" + "\n" +
"      }," + "\n" +
"      \"screenshareProps\":{" + "\n" +
"         \"screenshareConf\":\"79922-SCREENSHARE\"," + "\n" +
"         \"red5ScreenshareIp\":\"10.159.224.245\"," + "\n" +
"         \"red5ScreenshareApp\":\"video-broadcast\"" + "\n" +
"      }" + "\n" +
"   }" + "\n" +
"}"



var (
	MeetingList map[string]*Meeting
)

type Meeting struct {
	Id       string
	Name     string
}

func init() {
	MeetingList = make(map[string]*Meeting)
	m := Meeting{"meeting_11111", "astaxie"}
	MeetingList["meeting_11111"] = &m
	bm, err := cache.NewCache("memory", `{"interval":60}`)
	if err == nil {
		beego.Debug("Testing cache")
		bm.Put("astaxie", 1, 10*time.Second)
		bm.Get("astaxie")
		bm.IsExist("astaxie")
		bm.Delete("astaxie")
	}
}

func AddMeeting(meeting Meeting) (MeetingId string) {
	messenger.SendMessage("CreateMeetingReqMsg", CREATE_PROPS)
	meeting.Id = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	MeetingList[meeting.Id] = &meeting
	return meeting.Id
}

func GetMeeting(MeetingId string) (meeting *Meeting, err error) {
	if v, ok := MeetingList[MeetingId]; ok {
		return v, nil
	}
	return nil, errors.New("MeetingId Not Exist")
}

func GetAllMeetings() map[string]*Meeting {
	return MeetingList
}

func UpdateMeeting(MeetingId string, Name string) (err error) {
	if v, ok := MeetingList[MeetingId]; ok {
		v.Name = Name
		return nil
	}
	return errors.New("MeetingId Not Exist")
}

func DeleteMeeting(MeetingId string) {
	delete(MeetingList, MeetingId)
}
