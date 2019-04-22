package main

import (
	"fmt"

	"github.com/deaswang/jpush-api-golang"
)

// AppKey set value from JPush web
const (
	Appkey       = "appkey"
	masterSecret = "secret"
)

func main() {
	j := jpush.NewJPush(Appkey, masterSecret)
	aud := &jpush.PushAudience{}
	aud.SetAll(true)
	aud.Aud = &jpush.Audience{}
	aud.Aud.Tag = []string{"tag"}
	req := &jpush.PushRequest{
		Platform: &jpush.Platform{Platforms: []string{"android", "ios"}},
		Audience: aud,
		Notification: &jpush.PushNotification{
			Alert: "test alert",
			Android: &jpush.NotificationAndroid{
				Alert:     "alert",
				Title:     "title",
				BuilderID: 0,
				Priority:  1,
				AlertType: 7,
			},
		},
		Options: &jpush.PushOptions{
			TimeToLive: 0,
		},
	}
	ret, err := j.Push(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("result:", ret)
	return
}
