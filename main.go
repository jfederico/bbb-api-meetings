package main

import (
	_ "bbb-api-meetings/routers"

	"time"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
	//"github.com/astaxie/beego/logs"
	//"github.com/beego/redigo/redis"
	//"github.com/gomodule/redigo/redis"
	//"github.com/go-redis/redis"
)


var redisToAkka, redisFromAkka, bm, err string

func main() {
	//log := logs.NewLogger(10000)
	//log.SetLogger("console")
	//log.Debug("************************* Main ****************************")
	beego.Debug("this is debug")
	beego.Info("this is info")
	beego.Notice("this is notice")
	beego.Warn("this is warn")
	beego.Error("this is error")
	beego.Critical("this is critical")
	beego.Alert("this is alert")
	beego.Emergency("this is emergency")

	bm, err := cache.NewCache("memory", `{"interval":60}`)
	if err == nil {
		beego.Debug("Testing cache...")
		bm.Put("astaxie", 1, 10*time.Second)
		var astaxie = bm.Get("astaxie")
		beego.Debug("Data from memory ", astaxie)
		bm.IsExist("astaxie")
		bm.Delete("astaxie")
	}
	//redisToAkka, err := cache.NewCache("redis", `{"conn":":6039","dbNum":"0"}`)
	redisToAkka, err := cache.NewCache("redis", `{"key":"to-akka-apps-redis-channel","conn":":6379"}`)
	if err == nil {
		beego.Debug("Redis ToAkka initialized...")
		var astaxie = redisToAkka.Get("astaxie")
		beego.Debug("Data from redis ", astaxie)
	} else {
		beego.Debug(err)
	}

	redisFromAkka, err := cache.NewCache("redis", `{"key":"from-akka-apps-redis-channel","conn":":6379","dbNum":"0"}`)
	if err == nil {
		beego.Debug("Redis redisFromAkka initialized...")
		var astaxie = redisFromAkka.Get("astaxie")
		beego.Debug("Data from redis ", astaxie)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
