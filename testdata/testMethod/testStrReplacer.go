package testMethod

import (
	"fmt"
	"regexp"
)

func ReplaceTestDataDirPath(con string) string {
	testDataDirPath := GetTestDataDirPath()
	testDataDirPathRegex := regexp.MustCompile("\t.*/testdata/")
	return testDataDirPathRegex.ReplaceAllString(
		con, fmt.Sprintf("\t%s/", testDataDirPath),
	)
}
