package idcard

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var tail = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
var province = map[string]string{
	"11": "北京市",
	"12": "天津市",
	"13": "河北省",
	"14": "山西省",
	"15": "内蒙古自治区",
	"21": "辽宁省",
	"22": "吉林省",
	"23": "黑龙江省",
	"31": "上海市",
	"32": "江苏省",
	"33": "浙江省",
	"34": "安徽省",
	"35": "福建省",
	"36": "江西省",
	"37": "山东省",
	"41": "河南省",
	"42": "湖北省",
	"43": "湖南省",
	"44": "广东省",
	"45": "广西壮族自治区",
	"46": "海南省",
	"50": "重庆市",
	"51": "四川省",
	"52": "贵州省",
	"53": "云南省",
	"54": "西藏自治区",
	"61": "陕西省",
	"62": "甘肃省",
	"63": "青海省",
	"64": "宁夏回族自治区",
	"65": "新疆维吾尔自治区",
	"71": "台湾省",
	"81": "香港特别行政区",
	"91": "澳门特别行政区",
}

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

// IsValidCitizenNo 有效身份证
func IsValidCitizenNo(citizenNo *[]byte) bool {
	if !IsValidCitizenNo18(citizenNo) {
		return false
	}

	for i, v := range *citizenNo {
		n, _ := strconv.Atoi(string(v))
		if n >= 0 && n <= 9 {
			continue
		}

		if v == 'X' && i == 16 {
			continue
		}

		return false
	}

	if !checkProvince(*citizenNo) {
		return false
	}

	nYear, _ := strconv.Atoi(string((*citizenNo)[6:10]))
	nMonth, _ := strconv.Atoi(string((*citizenNo)[10:12]))
	nDay, _ := strconv.Atoi(string((*citizenNo)[12:14]))
	if !checkBirthday(nYear, nMonth, nDay) {
		return false
	}

	return true
}

func IsValidCitizenNo18(citizenNo18 *[]byte) bool {
	nLen := len(*citizenNo18)
	if nLen != 18 {
		return false
	}

	nSum := 0
	for i := 0; i < nLen-1; i++ {
		n, _ := strconv.Atoi(string((*citizenNo18)[i]))
		//fmt.Println(n)
		nSum += n * weight[i]
		fmt.Println(nSum)
	}
	mod := nSum % 11
	if tail[mod] == (*citizenNo18)[17] {
		return true
	}

	return false
}

func checkProvince(citizenNo []byte) bool {
	provinceCode := make([]byte, 0)
	provinceCode = append(provinceCode, citizenNo[:2]...)
	provinceStr := string(provinceCode)

	for i := range province {
		if provinceStr == i {
			return true
		}
	}

	return false
}

func checkBirthday(nYear, nMonth, nDay int) bool {
	if nYear < 1900 || nMonth <= 0 || nMonth > 12 || nDay <= 0 || nDay > 31 {
		return false
	}

	curYear, curMonth, curDay := time.Now().Date()
	if nYear == curYear {
		if nMonth > int(curMonth) {
			return false
		} else if nMonth == int(curMonth) && nDay > curDay {
			return false
		}
	}

	if 2 == nMonth {
		if isLeapYear(nYear) {
			if nDay > 29 {
				return false
			}
		} else if nDay > 28 {
			return false
		}
	} else if 4 == nMonth || 6 == nMonth || 9 == nMonth || 11 == nMonth {
		if nDay > 30 {
			return false
		}
	}

	return true
}

func isLeapYear(nYear int) bool {
	if nYear <= 0 {
		return false
	}

	if (nYear%4 == 0 && nYear%100 != 0) || nYear%400 == 0 {
		return true
	}

	return false
}
