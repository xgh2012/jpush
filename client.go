/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/10 15:42
 */

package jpush

import "encoding/base64"

type Client struct {
	AppKey  string
	Secret  string
	Headers map[string]string
}

//初始化
func (b *Client) Init(key, secret string) *Client {
	headers := make(map[string]string)
	headers["User-Agent"] = "jpush-xgh2012"
	headers["Connection"] = "keep-alive"

	b.AppKey = key
	b.Secret = secret
	headers["Authorization"] = b.BasicAuth()
	b.Headers = headers
	return b
}

//授权 签名
func (b *Client) BasicAuth() string {
	auth := b.AppKey + ":" + b.Secret
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
