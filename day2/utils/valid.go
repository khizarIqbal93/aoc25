package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func IsValid(id int) bool {
	str := strconv.Itoa(id)

	if len(str)%2 != 0 {
		return true
	}

	middle := (len(str) / 2)

	firstHalf := str[:middle]
	secondHalf := str[middle:]

	return firstHalf != secondHalf
}

func IsValidV2(id int) bool {
	str := strconv.Itoa(id)
	for i := 1; i <= len(str)/2; i++ {
		subStr := str[:i]
		r, _ := regexp.Compile(subStr)
		matches := r.FindAllString(str, -1)
		final := strings.Join(matches, "")
		if str == final {
			return false
		}
	}
	return true
}
