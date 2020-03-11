/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/10 15:44
* 可选参数
 */

package pushapi

type Options struct {
	JpushResult       *JpushResult           `json:"-"`
	SendNo            int                    `json:"sendno,omitempty"`              //推送序号
	TimeToLive        int                    `json:"time_to_live,omitempty"`        //离线消息保留时长(秒)
	OverrideMsgId     int64                  `json:"override_msg_id,omitempty"`     //要覆盖的消息 ID
	ApnsProduction    bool                   `json:"apns_production"`               //APNs 是否生产环境
	ApnsCollapseId    string                 `json:"apns_collapse_id,omitempty"`    //更新 iOS 通知的标识符
	BigPushDuration   int                    `json:"big_push_duration,omitempty"`   //定速推送时长(分钟)
	ThirdPartyChannel map[string]interface{} `json:"third_party_channel,omitempty"` //推送请求下发通道
}

//初始化
func InitOptions(env string) *Options {
	object := &Options{
		ApnsProduction: false,
	}

	if env == "prd" {
		object.ApnsProduction = true
	}

	return object
}
