/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/11 14:14
 */

package pushapi

import "github.com/xgh2012/jpush/jpushcommon"

type JpushResult struct {
	Flag     bool
	Message  string
	Result   interface{}
	HttpCode int
	Limit    *jpushcommon.RateLimitInfo
}
