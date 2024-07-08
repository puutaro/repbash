package stringMethod

import "strings"

func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func TrimBothQuote(str *string) {
	*str = strings.TrimSpace(*str)
	*str = strings.Trim(*str, "'")
	*str = strings.Trim(*str, `"`)
	*str = strings.Trim(*str, "`")
	return
}
