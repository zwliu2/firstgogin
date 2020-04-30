package main

import (
	"fmt"

	"regexp"
)

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

func main() {
	str := "问11"
	StrFilterNonChinese(&str)
	fmt.Println(str)
}

func StrFilterNonChinese(src *string) {
	strn := ""
	for _, c := range *src {
		if hzRegexp.MatchString(string(c)) {
			strn += string(c)
		}
	}

	*src = strn
}