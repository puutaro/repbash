package testMethod

import "os"

func ReadFileForTest(path string) string {
	con, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(con)
}
