/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/10 15:43
* “通知”对象，是一条推送的实体内容对象之一（另一个是“消息”），是会作为“通知”推送到客户端的。
* 其下属属性包含 4 种，3 个平台属性，以及一个 "alert" 属性。
 */

package pushapi

type Notification struct {
	JpushResult *JpushResult        `json:"-"`
	Alert       string              `json:"alert"` //通知内容
	Android     AndroidNotification `json:"android,omitempty"`
	Ios         IosNotification     `json:"ios,omitempty"`
}

type AndroidNotification struct {
	Alert             string                 `json:"alert,omitempty"`              //	这里指定了，则会覆盖上级统一指定的 alert 信息；内容可以为空字符串，则表示不展示到通知栏。
	Title             string                 `json:"title,omitempty"`              //	如果指定了，则通知里原来展示 App 名称的地方，将展示成这个字段。
	BuilderId         int                    `json:"builder_id,omitempty"`         // 通知栏样式 ID Android SDK 可设置通知栏样式，这里根据样式 ID 来指定该使用哪套样式，android 8.0 开始建议采用NotificationChannel配置。
	ChannelId         string                 `json:"channel_id,omitempty"`         //	android通知channel_id 不超过1000字节，Android 8.0开始可以进行NotificationChannel配置，这里根据channel ID 来指定通知栏展示效果。
	Priority          int                    `json:"priority,omitempty"`           //通知栏展示优先级	默认为 0，范围为 -2～2。
	Category          string                 `json:"category,omitempty"`           //通知栏条目过滤或排序 完全依赖 rom 厂商对 category 的处理策略
	Style             int                    `json:"style,omitempty"`              //	通知栏样式类型 bigText=1，Inbox=2，bigPicture=3。
	AlertType         int                    `json:"alert_type,omitempty"`         //	通知提醒方式
	BigText           string                 `json:"big_text,omitempty"`           //	大文本通知栏样式
	Inbox             string                 `json:"inbox,omitempty"`              //	文本条目通知栏样式 JSONObject (当 style = 2 时可用， json 的每个 key 对应的 value 会被当作文本条目逐条展示。支持 api 16 以上的 rom。)
	BigPicPath        string                 `json:"big_pic_path,omitempty"`       //大图片通知栏样式 (当 style = 3 时可用，可以是网络图片 url，或本地图片的 path，目前支持 .jpg 和 .png 后缀的图片。)
	Extras            map[string]interface{} `json:"extras,omitempty"`             //这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。
	LargeIcon         string                 `json:"large_icon,omitempty"`         //通知栏大图标 (图标路径可以是以http或https开头的网络图片)
	Sound             string                 `json:"sound,omitempty"`              //填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`    //定时展示开始时间（yyyy-MM-dd HH:mm:ss）
	ShowEndTime       string                 `json:"show_end_time,omitempty"`      //定时展示结束时间（yyyy-MM-dd HH:mm:ss）
	DisplayForeground string                 `json:"display_foreground,omitempty"` //APP在前台，通知是否展示  值为 "1" 时，APP 在前台会弹出通知栏消息；值为 "0" 时，APP 在前台不会弹出通知栏消息。注：默认情况下 APP 在前台会弹出通知栏消息。
}

type IosNotification struct {
	Alert            string                 `json:"alert,omitempty"`             //	这里指定了，则会覆盖上级统一指定的 alert 信息；内容可以为空字符串，则表示不展示到通知栏。
	Sound            string                 `json:"sound,omitempty"`             //通知提示声音或警告通知
	Badge            int                    `json:"badge,omitempty"`             //应用角标 如果不填，表示不改变角标数字，否则把角标数字改为指定的数字；为 0 表示清除。JPush 官方 SDK 会默认填充 badge 值为 "+1"
	ContentAvailable bool                   `json:"content-available,omitempty"` //推送唤醒
	MutableContent   bool                   `json:"mutable-content,omitempty"`   //通知扩展 推送的时候携带 ”mutable-content":true 说明是支持iOS10的UNNotificationServiceExtension
	Extras           map[string]interface{} `json:"extras,omitempty"`            //这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。
	ThreadId         string                 `json:"thread_id,omitempty"`         //通知分组
}

//初始化
func InitNotification(title, alert string, JpushResult *JpushResult) *Notification {
	object := &Notification{
		Alert: alert,
		Android: AndroidNotification{
			Title:     title,
			Alert:     alert,
			BuilderId: 1,
		},
		Ios: IosNotification{
			Alert:            alert,
			ContentAvailable: true,
			MutableContent:   true,
			Badge:            1,
			Sound:            "sound.caf",
		},
	}
	return object
}

//设置声音
func (b *Notification) SetSound() {

}

//设置声音
func (b *Notification) SetContentAvailable() {

}

//添加额外字段
func (b *Notification) AddExtra(key string, value interface{}) {
	if b.Android.Extras == nil {
		b.Android.Extras = make(map[string]interface{})
	}
	b.Android.Extras[key] = value

	if b.Ios.Extras == nil {
		b.Ios.Extras = make(map[string]interface{})
	}
	b.Ios.Extras[key] = value
}
