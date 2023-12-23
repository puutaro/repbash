package mapMethod

import (
	"fmt"
	"strings"
)

func Concat(
	mainMap map[string]string,
	addMap map[string]string,
) {
	if mainMap == nil &&
		addMap == nil {
		return
	}
	if addMap == nil {
		return
	}

	for k, v := range addMap {
		mainMap[k] = v
	}
	return
}

func ReplaceConByMap(
	con *string,
	mainRepValMap map[string]string,
) {
	if mainRepValMap == nil {
		return
	}
	for valName, value := range mainRepValMap {
		*con = strings.Replace(
			*con,
			fmt.Sprintf("${%s}",
				valName,
			),
			value,
			-1,
		)
	}
	return
}
