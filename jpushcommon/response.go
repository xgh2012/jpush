/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/11 10:34
 */

package jpushcommon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Response struct {
	Resp          *http.Response
	Err           error
	StatusCode    int // HTTP状态
	RateLimitInfo *RateLimitInfo
}

// 频率限制
type RateLimitInfo struct {
	RateLimitLimit     int
	RateLimitRemaining int
	RateLimitReset     int
}

func (r *Response) GetResult(dest interface{}) error {
	if r.Err != nil {
		r.StatusCode = 0
		return r.Err
	}

	r.StatusCode = r.Resp.StatusCode

	r.ResultHeader()
	return r.ResultJson(dest)
}

func (r *Response) ResultHeader() {
	r.RateLimitInfo = &RateLimitInfo{}

	r.RateLimitInfo.RateLimitLimit, _ = strconv.Atoi(r.Resp.Header.Get("X-Rate-Limit-Limit"))
	r.RateLimitInfo.RateLimitRemaining, _ = strconv.Atoi(r.Resp.Header.Get("X-Rate-Limit-Remaining"))
	r.RateLimitInfo.RateLimitReset, _ = strconv.Atoi(r.Resp.Header.Get("X-Rate-Limit-Reset"))
}

//返回结果 是json
func (r *Response) ResultJson(dest interface{}) error {
	body, err := ioutil.ReadAll(r.Resp.Body)
	if err != nil {
		return err
	}
	defer r.Resp.Body.Close()
	return json.Unmarshal(body, &dest)
}
