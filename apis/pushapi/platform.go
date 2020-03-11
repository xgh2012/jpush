/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/10 15:42
* 设置发送平台
 */

package pushapi

import (
	"strings"
)

type Platform struct {
	JpushResult *JpushResult `json:"-"`
	All         bool
	Value       []string
}

//设置所有
func (b *Platform) SetAll() {
	b.All = true
}

// 添加 platform "ios", "android", "winphone"
// 为方便调用，直接使用 英文 ,分割
func (b *Platform) Add(plt string) {
	if b.All {
		return
	}
	platforms := strings.Split(strings.ToLower(plt), ",")
	if len(platforms) <= 0 {
		b.JpushResult.Flag = false
		b.JpushResult.Message = "platform 设置错误"
		return
	}

	//初始化
	if b.Value == nil {
		b.Value = []string{}
	}

	//去重
	for _, v := range platforms {
		if b.inArray(v, b.Values()) {
			continue
		}

		if v != "ios" && v != "android" && v != "winphone" {
			b.JpushResult.Flag = false
			b.JpushResult.Message = "platform 可选【ios、android、winphone】"
			return
		}
		b.Value = append(b.Value, v)
	}

	return
}

func (b *Platform) Values() []string {
	return b.Value
}

//去重
func (b *Platform) inArray(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}
