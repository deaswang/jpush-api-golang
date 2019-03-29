package main

import (
	"encoding/json"
	"fmt"

	"github.com/cochainio/jpush-api-golang"
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
	result, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(string(result))
	}
	return
}
