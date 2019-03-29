# JPush API Golang

## 概述
这是 JPush REST API 的 Golang 版本封装开发包，不是由极光推送官方提供的，一般支持最新的 API 功能。

对应的 REST API 文档：<https://docs.jiguang.cn/jpush/server/push/server_overview/>

## 兼容版本
+  Golang 1.10

## 环境配置

```bash
go get github.com/cochain/jpush-api-golang
```

## 代码样例

>   代码样例在 jpush-api-golang 中的 example 文件夹中，[点击查看所有 example ](https://github.com) 。

>   这个样例演示了消息推送。

```golang
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
```

## HTTP 状态码

参考文档：<http://docs.jiguang.cn/jpush/server/push/http_status_code/>

Push v3 API 状态码 参考文档：<http://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/>

Report API  状态码 参考文档：<http://docs.jiguang.cn/jpush/server/push/rest_api_v3_report/>

Device API 状态码 参考文档：<http://docs.jiguang.cn/jpush/server/push/rest_api_v3_device/>

Push Schedule API 状态码 参考文档：<http://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule/>

[Release页面](https://github.com) 有详细的版本发布记录与下载。
