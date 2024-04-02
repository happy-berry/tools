package idcard

import (
	"regexp"
	"strconv"
	"time"
)

// Replace 身份证敏感信息替换
func Replace(str string) string {
	return str[0:4] + "**********" + str[13:17]
}

// GetAge 获取年龄
func GetAge(idNum string) int {
	reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)(\d{2})((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	params := reg.FindStringSubmatch(idNum)
	birYear, _ := strconv.Atoi(params[1] + params[2])
	birMonth, _ := strconv.Atoi(params[3])
	age := time.Now().Year() - birYear
	if int(time.Now().Month()) < birMonth {
		age--
	}
	return age
}
