/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/10 15:43
* message：自定义消息 应用内消息。或者称作：自定义消息，透传消息。
 */

package pushapi

type Message struct {
	JpushResult *JpushResult           `json:"-"`
	MsgContent  string                 `json:"msg_content"` //消息内容本身
	Title       string                 `json:"title"`
	ContentType string                 `json:"content_type"`
	Extras      map[string]interface{} `json:"extras"` //这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。
}

//初始化
func InitMessage(title, alert, contentType string) *Message {
	object := &Message{
		Title:       title,
		MsgContent:  alert,
		ContentType: contentType,
	}
	return object
}

//添加额外字段
func (b *Message) AddExtra(key string, value interface{}) {
	if b.Extras == nil {
		b.Extras = make(map[string]interface{})
	}
	b.Extras[key] = value
}
