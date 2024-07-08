package searcher

import (
	"errors"
	"github.com/puutaro/repbash/pkg/util/stringMethod"
	"io/fs"
	"os"
	"path/filepath"
)

func Search(
	launchShellPath string,
) (searchReplaceVariablesTsvPath string, err error) {
	replaceVariableTsvRelativePath := os.Getenv("REPLACE_VARIABLES_TSV_RELATIVE_PATH")
	if stringMethod.IsEmpty(replaceVariableTsvRelativePath) {
		return "",
			errors.New("not found env: REPLACE_VARIABLES_TSV_RELATIVE_PATH")
	}
	replaceVariableTsvName := filepath.Base(
		replaceVariableTsvRelativePath,
	)
	parentDirPath := filepath.Dir(launchShellPath)

	_, err = os.Stat(launchShellPath)
	if err != nil {
		return "", err
	}
	var stopWalk = errors.New("stop walking")
	for i := 0; i < 5; i++ {
		err = filepath.WalkDir(parentDirPath, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			if d.Name() != replaceVariableTsvName {
				return nil
			}
			//fmt.Printf("Base: %s\n", path)
			searchReplaceVariablesTsvPath = path
			return filepath.SkipAll
		})
		if !stringMethod.IsEmpty(searchReplaceVariablesTsvPath) {
			break
		}
		parentDirPath = filepath.Dir(parentDirPath)
	}
	if errors.Is(stopWalk, err) {
		return searchReplaceVariablesTsvPath, nil
	}
	return

}
