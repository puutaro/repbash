package maker

import (
	"fmt"
	"github.com/puutaro/repbash/pkg/util/stringMethod"
	"os"
	"strings"
)

type TsvReader interface {
	ReadTsv(tsvPath string) (string, error)
}

func NewTsvReader() TsvReader {
	return &tsvReader{}
}

type tsvReader struct{}

func (tr tsvReader) ReadTsv(tsvPath string) (string, error) {
	return readFile(
		tsvPath,
	)
}

type tsvReadService struct {
	reader TsvReader
}

func (trs tsvReadService) ReadTsv(tsvPath string) (string, error) {
	return trs.reader.ReadTsv(tsvPath)
}

func TsvToRepVarMap(
	tsvPath string,
	tsvReadService tsvReadService,
) (map[string]string, error) {
	repValMap := make(map[string]string)
	tsvCon, err := tsvReadService.ReadTsv(
		tsvPath,
	)
	if err != nil {
		return repValMap, err
	}
	repValConList := convertTsvToList(
		tsvCon,
	)
	lastRepValConList := replaceRepValRecursively(
		repValConList,
	)
	repValMap = convertRepValListToMap(
		lastRepValConList,
	)
	return repValMap, nil
}

func readFile(
	tsvPath string,
) (
	tsvCon string,
	err error,
) {
	_, err = os.Stat(tsvPath)
	if err != nil {
		return tsvCon, fmt.Errorf(
			"not found %s\n",
			tsvPath,
		)
	}
	tsvByteCon, err := os.ReadFile(tsvPath)
	if err != nil {
		return tsvCon, fmt.Errorf(
			"read err %s\n",
			tsvPath,
		)
	}
	tsvCon = string(tsvByteCon)
	return
}

func convertRepValListToMap(
	repValConList []string,
) (repValMap map[string]string) {
	repValMap = make(map[string]string)
	for _, line := range repValConList {
		trimLine := strings.TrimSpace(line)
		if stringMethod.IsEmpty(trimLine) ||
			strings.HasPrefix(trimLine, "//") {
			continue
		}
		trimLineKeyValueList := strings.Split(trimLine, "\t")
		if len(trimLineKeyValueList) < 2 {
			continue
		}
		valName := trimLineKeyValueList[0]
		stringMethod.TrimBothQuote(
			&valName,
		)
		valValue := strings.Join(trimLineKeyValueList[1:], "\t")
		stringMethod.TrimBothQuote(
			&valValue,
		)
		repValMap[valName] = valValue
	}
	return
}

func convertTsvToList(
	tsvCon string,
) (tsvConList []string) {
	tsvConListSrc := strings.Split(
		tsvCon,
		"\n",
	)
	for _, line := range tsvConListSrc {
		trimLine := strings.TrimSpace(line)
		if stringMethod.IsEmpty(trimLine) ||
			strings.HasPrefix(trimLine, "//") {
			continue
		}
		trimLineKeyValueList := strings.Split(trimLine, "\t")
		if len(trimLineKeyValueList) < 2 {
			continue
		}
		valName := trimLineKeyValueList[0]
		stringMethod.TrimBothQuote(
			&valName,
		)
		valValue := strings.Join(trimLineKeyValueList[1:], "\t")
		stringMethod.TrimBothQuote(
			&valValue,
		)
		tsvConList = append(
			tsvConList, fmt.Sprintf(
				"%s\t%s",
				valName,
				valValue,
			),
		)
	}
	return
}

func replaceRepValRecursively(
	repValConList []string,
) (lastRepValConList []string) {
	lastRepValConList = repValConList
	lastRepValConListLength := len(lastRepValConList)
	for i := 0; i < lastRepValConListLength; i++ {
		line := lastRepValConList[i]
		trimLineKeyValueList := strings.Split(line, "\t")
		if len(trimLineKeyValueList) < 2 {
			continue
		}
		valName := fmt.Sprintf("${%s}", trimLineKeyValueList[0])
		valValue := strings.Join(trimLineKeyValueList[1:], "\t")
		var tempRepValConList []string
		for _, lastRepValLine := range lastRepValConList {
			tempRepValConList = append(
				tempRepValConList,
				strings.Replace(
					lastRepValLine,
					valName,
					valValue,
					-1,
				),
			)
		}
		lastRepValConList = tempRepValConList
	}
	return
}
