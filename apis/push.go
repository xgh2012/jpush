/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/11 10:18
 */

package apis

import (
	"github.com/xgh2012/jpush"
	"github.com/xgh2012/jpush/apis/pushapi"
	"github.com/xgh2012/jpush/jpushcommon"
)

type Push struct {
	Platform     *pushapi.Platform
	Audience     *pushapi.Audience
	Notification *pushapi.Notification
	Message      *pushapi.Message
	Options      *pushapi.Options
	JpushResult  *pushapi.JpushResult
}
type Body struct {
	Platform        interface{} `json:"platform"`
	Audience        interface{} `json:"audience"`
	Notification    interface{} `json:"notification"`
	Message         interface{} `json:"message,omitempty"`
	SmsMessage      interface{} `json:"sms_message,omitempty"`
	Options         interface{} `json:"options,omitempty"`
	Callback        interface{} `json:"callback,omitempty"`
	Notification3rd interface{} `json:"notification_3rd,omitempty"`
	Cid             string      `json:"cid,omitempty"`
}

type PushResult struct {
	Error  ErrorInfo
	SendNo string `json:"sendno"`
	MsgId  string `json:"msg_id"`
}
type ErrorInfo struct {
	Code    int
	Message string
}

func (api *Push) Init(title, alert string) *Push {
	TmpRes := &pushapi.JpushResult{Flag: true, Message: "成功"}
	api.Platform = &pushapi.Platform{JpushResult: TmpRes}
	api.Audience = &pushapi.Audience{JpushResult: TmpRes}
	api.Notification = pushapi.InitNotification(title, alert, TmpRes)
	api.Message = &pushapi.Message{JpushResult: TmpRes}
	api.Options = &pushapi.Options{JpushResult: TmpRes, ApnsProduction: false}
	api.JpushResult = TmpRes

	return api
}

//设置发送平台 all ios android 可选
func (api *Push) PlatformSet(plt string) *Push {
	if plt == "all" {
		api.Platform.SetAll()
	} else {
		api.Platform.Add(plt)
	}
	return api
}

//全部发送
func (api *Push) AudienceAll() *Push {
	api.Audience.SetAll()
	return api
}

//AudienceAddTag
func (api *Push) AudienceAddTag(tag string) *Push {
	api.Audience.AddTag(tag)
	return api
}

//AudienceAddTagAnd
func (api *Push) AudienceAddTagAnd(tag string) *Push {
	api.Audience.AddTagAnd(tag)
	return api
}

//AudienceAddTagNot
func (api *Push) AudienceAddTagNot(tag string) *Push {
	api.Audience.AddTagNot(tag)
	return api
}

//AudienceAddTagNot
func (api *Push) AudienceAddAlias(alias string) *Push {
	api.Audience.AddAlias(alias)
	return api
}

//AudienceAddRegistrationId
func (api *Push) AudienceAddRegistrationId(registrationId string) *Push {
	api.Audience.AddRegistrationId(registrationId)
	return api
}

//通知 扩展字段
func (api *Push) NotificationAddExtra(key string, value interface{}) *Push {
	api.Notification.AddExtra(key, value)
	return api
}

//自定义消息 主体
func (api *Push) MessageInit(title, content, msgType string) *Push {
	api.Message.Title = title
	api.Message.MsgContent = content
	api.Message.ContentType = msgType
	return api
}

//自定义消息 扩展字段
func (api *Push) MessageAddExtra(key string, value interface{}) *Push {
	api.Message.AddExtra(key, value)
	return api
}

//可选参数 设置环境
func (api *Push) OptionsSetEnv(env bool) *Push {
	api.Options.ApnsProduction = env
	return api
}

//取得最终数据
func (api *Push) GetBody() *Body {
	body := &Body{}
	if api.Platform.All {
		body.Platform = "all"
	} else {
		body.Platform = api.Platform.Value
	}

	if api.Audience.All == true {
		body.Audience = "all"
	} else if len(api.Audience.Value) > 0 {
		body.Audience = api.Audience.Value
	}

	body.Notification = api.Notification

	if api.Message.MsgContent != "" {
		body.Message = api.Message
	}
	body.Options = api.Options
	return body
}

//发送请求
func (api *Push) DoPush(client *jpush.Client) {
	if api.JpushResult.Flag == false {
		return
	}
	body := api.GetBody()

	c := jpushcommon.NewClient(true)
	resp, err := c.PostJson(jpushcommon.PUSHAPI, body, client.Headers)

	response := jpushcommon.Response{
		Resp: resp,
		Err:  err,
	}

	rst := &PushResult{}
	err = response.GetResult(rst)

	api.JpushResult.HttpCode = response.StatusCode

	if err != nil {
		api.JpushResult.Flag = false
		api.JpushResult.Message = err.Error()
		return
	}

	api.JpushResult.Limit = response.RateLimitInfo
	api.JpushResult.Result = rst
}
