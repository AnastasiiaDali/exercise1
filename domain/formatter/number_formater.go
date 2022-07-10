package formatter

import (
	"fmt"
	"regexp"
	"strconv"
)

type NumberFormatter struct {
}

func New() *NumberFormatter {
	return &NumberFormatter{}
}

//this function formats the number to thousands separated by comma

func (receiver NumberFormatter) FormatNumbers(num int) string {
	if num > 9999 {
		str := fmt.Sprintf("%d", num)
		re := regexp.MustCompile("(\\d+)(\\d{3})")

		for n := ""; n != str; {
			n = str
			str = re.ReplaceAllString(str, "$1,$2")
		}
		return str
	}
	return strconv.Itoa(num)
}
