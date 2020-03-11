##使用demo
第一版本只实现 push接口
部分借鉴 github.com/DeanThompson/jpush-api-go-client
```
client := new(jpush.Client).Init("your key", "your secret")
push := new(apis.Push).
    Init("测试标题", "c测试内容").
    PlatformSet("all").
    OptionsSetEnv(false).
    AudienceAddRegistrationId("1517bfd3f7a8bb972ec").
    //AudienceAddAlias("alias1,alias2,alias3,alias4").
    //AudienceAddTag("tag1,tag2,tag3,tag4").
    NotificationAddExtra("notification1", "notification_内容1").
    NotificationAddExtra("notification2", "notification_内容2")
push.DoPush(client)

fmt.Printf("%+v\n", push.JpushResult)
fmt.Printf("%+v\n", push.JpushResult.Result)
```