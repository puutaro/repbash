package maker

import (
	"fmt"
	"github.com/puutaro/repbash/pkg/util/mapMethod"
)

func MakeMainRepValMap(
	srcTsvPaths []string,
	tsvReader TsvReader,
) (mainRepValMap map[string]string, err error) {
	mainRepValMap = make(map[string]string)
	for _, srcTsvPath := range srcTsvPaths {
		mapMethod.ReplaceConByMap(
			&srcTsvPath,
			mainRepValMap,
		)
		repValMap, err := TsvToRepVarMap(
			srcTsvPath,
			tsvReadService{reader: tsvReader},
		)
		if err != nil {
			return mainRepValMap,
				fmt.Errorf(err.Error())
		}
		mapMethod.Concat(
			mainRepValMap,
			repValMap,
		)
	}
	return
}
