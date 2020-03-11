/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/3/10 15:43
* audience：推送目标
 */

package pushapi

import (
	"strconv"
	"strings"
)

type Audience struct {
	JpushResult *JpushResult `json:"-"`
	All         bool
	Value       map[string][]string
}

func (b *Audience) SetAll() {
	b.All = true
}

// 标签OR 为方便调用，直接使用 英文 ,分割
func (b *Audience) AddTag(tag string) {
	b.setValues("tag", tag, 20)
}

// 标签OR 为方便调用，直接使用 英文 ,分割
func (b *Audience) AddTagAnd(tagAnd string) {
	b.setValues("tag_and", tagAnd, 20)
}

// 标签OR 为方便调用，直接使用 英文 ,分割
func (b *Audience) AddTagNot(tagNot string) {
	b.setValues("tag_not", tagNot, 20)
}

// 多个别名之间是 OR 关系，即取并集 用别名来标识一个用户 英文 ,分割
func (b *Audience) AddAlias(alias string) {
	b.setValues("alias", alias, 1000)
}

// 多个别名之间是 OR 关系，即取并集 用别名来标识一个用户 英文 ,分割
func (b *Audience) AddRegistrationId(registrationId string) {
	b.setValues("registration_id", registrationId, 1000)
}

func (b *Audience) setValues(key, value string, maxLen int) {
	if b.All {
		return
	}

	if b.Value == nil {
		b.Value = make(map[string][]string)
	}

	if b.Value[key] == nil {
		b.Value[key] = make([]string, 0)
	}

	values := strings.Split(strings.ToLower(value), ",")
	lenValues := len(values)
	if lenValues <= 0 {
		b.JpushResult.Flag = false
		b.JpushResult.Message = key + " 设置错误"
		return
	} else if maxLen > 0 && lenValues > maxLen {
		b.JpushResult.Flag = false
		b.JpushResult.Message = key + " 最多" + strconv.Itoa(maxLen) + "个"
		return
	}

	//初始化
	if len(b.Value[key]) == 0 {
		b.Value[key] = values
	} else {
		for _, v := range values {
			b.Value[key] = append(b.Value[key], v)
		}
		if maxLen > 0 && len(b.Value[key]) > maxLen {
			b.JpushResult.Flag = false
			b.JpushResult.Message = key + " 最多" + strconv.Itoa(maxLen) + "个"
			return
		}
	}
}
