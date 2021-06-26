package isbn

import (
	"regexp"
	"strings"
)

func parsing(str string) (isbn string, valid bool) {
	if len(str) > 0 {
		ss := strings.Join(strings.Split(str, "-"), "")
		if len(ss) == 10 {
			isbn, valid = ss, true
		}
		return
	}
	return
}

func isbnProcessing(isbn string) bool {
	total := 0
	n := 0
	for i := 0; i < 10; i++ {
		if isbn[i] == 'X' {
			n = 10
		} else {
			n = int(isbn[i] - '0')
		}
		total += n * (10 - i)
	}
	return total%11 == 0
}

func IsValidISBN(isbn string) bool {
	isbnStr, valid := parsing(isbn)
	if valid {
		var reg = regexp.MustCompile("^[0-9]{9}[0-9X]$")
		if reg.MatchString(isbnStr) {
			return isbnProcessing(isbnStr)
		}
		return false
	}
	return false
}
