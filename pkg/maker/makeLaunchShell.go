package maker

import (
	"fmt"
	"github.com/puutaro/repbash/pkg/util/mapMethod"
	"github.com/puutaro/repbash/pkg/util/stringMethod"
	"github.com/puutaro/repbash/pkg/val"
	"io"
	"net/http"
	"os"
	"strings"
)

func MakeLaunchShell(
	ioGetter IoGetter,
	argsStruct val.ArgumentsStruct,
	mainRepValMap map[string]string,
) (resultShellCon string, err error) {
	launchShellPath := argsStruct.LaunchShellPath
	shellConByte, err := os.ReadFile(launchShellPath)
	if err != nil {
		return "", err
	}

	resultShellCon = string(shellConByte)
	err = concatImportCon(
		ioGetter,
		&resultShellCon,
		argsStruct,
		mainRepValMap,
	)
	if err != nil {
		return "", err
	}
	mapMethod.ReplaceConByMap(
		&resultShellCon,
		mainRepValMap,
	)
	replaceByAppVal(
		ioGetter,
		&resultShellCon,
		argsStruct,
		mainRepValMap,
	)
	return
}

func concatImportCon(
	ioGetter IoGetter,
	shellCon *string,
	argsStruct val.ArgumentsStruct,
	mainRepValMap map[string]string,
) (err error) {
	importPaths := argsStruct.ImportPaths
	maxProcessNum := 5
	goRoutineNum := len(importPaths)
	semaPhore := make(chan int, maxProcessNum)
	execChannel := make(chan val.Result, goRoutineNum)
	for _, importPath := range importPaths {
		mapMethod.ReplaceConByMap(
			&importPath,
			mainRepValMap,
		)
		go readImportCon(
			ioGetter,
			importPath,
			semaPhore,
			execChannel,
		)
	}
	var importCon string
	for i := 0; i < goRoutineNum; i++ {
		result := <-execChannel
		err = result.Error
		if err != nil {
			return
		}
		importCon = fmt.Sprintf("%s\n%s", importCon, result.GetCon)
	}
	close(semaPhore)
	close(execChannel)
	*shellCon = fmt.Sprintf(
		"%s\n%s",
		importCon,
		*shellCon,
	)
	return
}

type IoGetter interface {
	GetFromUrl(url string) (result val.Result)
	GetFannelPath(mainRepValMap map[string]string) string
}

func NewIoGetter() IoGetter {
	return &ioGetter{}
}

type ioGetter struct{}

func (ioGetter ioGetter) GetFromUrl(url string) (result val.Result) {
	return getFromWeb(url)
}

func (ioGetter ioGetter) GetFannelPath(mainRepValMap map[string]string) string {
	return getFannelPath(mainRepValMap)
}

func readImportCon(
	ioGetter IoGetter,
	importPath string,
	semaPhore chan int,
	execChannel chan val.Result,
) {
	semaPhore <- 0
	//fmt.Printf("## start importPath %s\n", importPath)
	//time.Sleep(
	//	time.Duration(rand.Int63n(3)) * time.Second)
	getResult := readHandler(
		ioGetter,
		importPath,
	)

	<-semaPhore
	execChannel <- getResult
	return
	//fmt.Printf("## end importPath %s\n", importPath)
}

func readHandler(
	ioGetter IoGetter,
	importPath string,
) (getResult val.Result) {
	isWeb := strings.HasPrefix(
		importPath,
		"https://",
	) || strings.HasPrefix(
		importPath,
		"http://",
	)
	switch isWeb {
	case true:
		return ioGetter.GetFromUrl(
			importPath,
		)
	default:
		return getFromLocal(
			importPath,
		)
	}
}

func getFromWeb(
	urlPath string,
) (getResult val.Result) {
	req, _ := http.NewRequest("GET", urlPath, nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return val.Result{
			"",
			err,
		}
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			getResult = val.Result{
				"",
				fmt.Errorf("defer close error: %v", closeErr),
			}
			return
		}
	}()

	byteArray, _ := io.ReadAll(resp.Body)
	return val.Result{string(byteArray), nil}
}

func getFromLocal(
	importPath string,
) (getResult val.Result) {
	shellConByte, err := os.ReadFile(importPath)
	if err != nil {
		getResult = val.Result{
			"", fmt.Errorf(
				"%v\n",
				err.Error(),
			),
		}
		return
	}
	//fmt.Printf("## importPath $s %s", importPath, string(shellConByte))
	getResult = val.Result{
		string(shellConByte),
		nil,
	}
	return
}

func replaceByAppVal(
	ioGetter IoGetter,
	targetCon *string,
	argsStruct val.ArgumentsStruct,
	mainRepValMap map[string]string,
) {
	launchShellPath := argsStruct.LaunchShellPath
	appRootPath := os.Getenv("APP_ROOT_PATH")
	appDirPath := os.Getenv("APP_DIR_PATH")
	replaceVariablesTsvRelativePath := os.Getenv("REPLACE_VARIABLES_TSV_RELATIVE_PATH")
	ubuntuServiceTempDirPath := os.Getenv("UBUNTU_SERVICE_TEMP_DIR_PATH")
	ubuntuEnvTsvName := os.Getenv("UBUNTU_ENV_TSV_NAME")
	*targetCon = strings.Replace(
		*targetCon,
		"\\\n",
		" ",
		-1,
	)
	*targetCon = strings.Replace(
		*targetCon,
		"${0}",
		launchShellPath,
		-1,
	)
	*targetCon = strings.Replace(
		*targetCon,
		"${APP_ROOT_PATH}",
		appRootPath,
		-1,
	)
	*targetCon = strings.Replace(
		*targetCon,
		"${APP_DIR_PATH}",
		appDirPath,
		-1,
	)
	*targetCon = strings.Replace(
		*targetCon,
		"${REPLACE_VARIABLES_TSV_RELATIVE_PATH}",
		replaceVariablesTsvRelativePath,
		-1,
	)
	*targetCon = strings.Replace(
		*targetCon,
		"${UBUNTU_SERVICE_TEMP_DIR_PATH}",
		ubuntuServiceTempDirPath,
		-1,
	)
	*targetCon = strings.Replace(
		*targetCon,
		"${UBUNTU_ENV_TSV_NAME}",
		ubuntuEnvTsvName,
		-1,
	)
	*targetCon = removeLaunchBashLine(
		targetCon,
		argsStruct,
	)
	replaceByCmdVal(
		ioGetter,
		targetCon,
		mainRepValMap,
	)
	return
}

func removeLaunchBashLine(
	targetCon *string,
	argsStruct val.ArgumentsStruct,
) (output string) {
	launchShellPath := argsStruct.LaunchShellPath
	isSaveRepBashLine := argsStruct.IsSaveRepbashLine
	for _, line := range strings.Split(*targetCon, "\n") {
		isExec := strings.Contains(line, "exec ")
		isBash := strings.Contains(line, "repbash ")
		isLaunchShellPath := strings.Contains(line, launchShellPath)
		if isExec &&
			isBash &&
			isLaunchShellPath &&
			!isSaveRepBashLine {
			continue
		}
		output += fmt.Sprintf("%s\n", line)
	}
	return
}

func getFannelPath(
	mainRepValMap map[string]string,
) string {
	fannelPath := mainRepValMap["FANNEL_PATH"]
	if stringMethod.IsEmpty(fannelPath) {
		return ""
	}
	return fannelPath
}

func readFannelCon(
	path string,
) (string, error) {
	fannelByteCon, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("not found fannel path %v", path)
	}
	return string(fannelByteCon), err
}

func replaceByCmdVal(
	ioGetter IoGetter,
	input *string,
	mainRepValMap map[string]string,
) {
	cmdValMap := makeCmdValMap(
		ioGetter,
		mainRepValMap,
	)
	if cmdValMap == nil {
		return
	}
	mapMethod.ReplaceConByMap(
		input,
		cmdValMap,
	)
	return
}

func makeCmdValMap(
	ioGetter IoGetter,
	mainRepValMap map[string]string,
) (cmdValMap map[string]string) {
	cmdValMap = make(map[string]string)
	fannelPath := ioGetter.GetFannelPath(
		mainRepValMap,
	)
	if stringMethod.IsEmpty(fannelPath) {
		return
	}
	fannelCon, err := readFannelCon(
		fannelPath,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cmdValStartHolderName := "CMD_VARIABLE_SECTION_START"
	cmdValEndHolderName := "CMD_VARIABLE_SECTION_END"
	cmdValStartHolderCount := 0
	cmdValEndHolderCount := 0
	for _, lineSrc := range strings.Split(fannelCon, "\n") {
		if strings.Contains(lineSrc, cmdValStartHolderName) {
			cmdValStartHolderCount++
		}
		if strings.Contains(lineSrc, cmdValEndHolderName) {
			cmdValEndHolderCount++
		}
		if cmdValStartHolderCount == 0 {
			continue
		}
		if cmdValEndHolderCount > 0 {
			break
		}
		line := strings.TrimSpace(lineSrc)
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "//") {
			continue
		}
		trimLine := strings.TrimSpace(line)
		valNameAndValueList := strings.Split(trimLine, "=")
		if len(valNameAndValueList) < 2 {
			continue
		}
		if len(valNameAndValueList) < 2 {
			continue
		}
		key := strings.TrimSpace(valNameAndValueList[0])
		value := strings.TrimSpace(
			strings.Join(valNameAndValueList[1:], "="),
		)
		stringMethod.TrimBothQuote(
			&value,
		)
		cmdValMap[key] = value
	}
	return
}
