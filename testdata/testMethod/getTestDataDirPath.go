package testMethod

import (
	"os"
	"path/filepath"
	"strings"
)

func GetTestDataDirPath() string {
	workingDirPath, err := os.Getwd()
	if err != nil {
		panic("cannot get wd path")
	}
	workingDirPathList := strings.Split(workingDirPath, "/")
	repbashIndex := indexOf("repbash", workingDirPathList)
	return filepath.Join(
		strings.Join(workingDirPathList[:repbashIndex+1], "/"),
		"testdata",
	)
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
