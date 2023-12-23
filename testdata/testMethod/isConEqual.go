package testMethod

import (
	"strings"
)

func IsConEqual(
	getCon string,
	wantsCon string,
) bool {
	return trim(getCon) == trim(wantsCon)

}

func trim(con string) (resultCon string) {
	trimCon := strings.Trim(
		strings.TrimSpace(con),
		"\n",
	)
	for _, line := range strings.Split(trimCon, "\n") {
		resultCon = strings.TrimSpace(line)
	}
	return
}
