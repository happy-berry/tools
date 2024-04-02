package phone

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"unicode"
)

// SmsCode 生成6位验证码
func SmsCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	smsCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return smsCode
}

//

func CheckPassword(ps string) error {
	if len(ps) < 8 {
		return fmt.Errorf("密码长度不能小于8位")
	}
	num := `[0-9]{1}`
	small := `[a-z]{1}`
	big := `[A-Z]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("至少包含一个数字")
	}
	if b, err := regexp.MatchString(small, ps); !b || err != nil {
		return fmt.Errorf("至少包含一个小写字母")
	}
	if b, err := regexp.MatchString(big, ps); !b || err != nil {
		return fmt.Errorf("至少包含一个大写字母")
	}
	if IsChineseChar(ps) {
		return fmt.Errorf("不能包含中文字符")
	}
	return nil
}

// IsChineseChar 是否包含中文
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
