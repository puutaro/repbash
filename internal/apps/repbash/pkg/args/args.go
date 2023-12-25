package args

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/puutaro/repbash/pkg/util/stringMethod"
	"github.com/puutaro/repbash/pkg/val"
	"os"
	"strings"
)

func Args() (
	argsType val.ArgumentsStruct,
	err error,
) {
	var args struct {
		LaunchShellPath   string   `arg:"positional, required" help:"launch shell path"`
		SrcTsvPaths       []string `arg:"-t,--src-tsv-path" help:"val name to val value tsv paths"`
		ArgsCon           string   `arg:"-a,--args-con" help:"args contetns key to value map list (Replace ${REPBASH_ARGS_CON} with this in script)"`
		IsSaveRepbashLine bool     `arg:"-s,--save-repbash-line" help:"save flag for repbash cmd line (Ordinary, replace 'exec repbash ~' with '### REPBASH_CON')"`
		ImportPaths       []string `arg:"-i,--import" help:"import shell file (include http url9"`
	}

	arg.MustParse(&args)
	argsType = val.ArgumentsStruct{
		args.LaunchShellPath,
		makeListType(args.SrcTsvPaths),
		args.ArgsCon,
		makeArgsMap(args.ArgsCon),
		makeListType(args.ImportPaths),
		args.IsSaveRepbashLine,
	}
	err = checkLaunchShellPath(argsType.LaunchShellPath)
	return
}

func makeListType(
	srcTsvPaths []string,
) []string {
	if len(srcTsvPaths) == 0 {
		return []string{}
	}
	return srcTsvPaths
}

func makeArgsMap(
	argsMapStr string,
) (returnArgsMap map[string]string) {
	returnArgsMap = make(map[string]string)

	if stringMethod.IsEmpty(argsMapStr) {
		return returnArgsMap
	}
	argsMapStrList := strings.Split(argsMapStr, ",")
	if len(argsMapStrList) == 0 {
		return returnArgsMap
	}
	for _, line := range argsMapStrList {
		keyAndValList := strings.Split(line, "=")
		if len(keyAndValList) < 2 {
			continue
		}
		valName := keyAndValList[0]
		stringMethod.TrimBothQuote(
			&valName,
		)
		valValue := strings.Join(keyAndValList[1:], "\t")
		stringMethod.TrimBothQuote(
			&valValue,
		)
		returnArgsMap[valName] = valValue
	}
	return
}

func checkLaunchShellPath(
	launchShellPath string,
) error {
	if len(launchShellPath) == 0 {
		return fmt.Errorf(
			"first LaunchShellPath not found",
		)
	}
	_, firstSrcTsvPathNotExistErr := os.Stat(launchShellPath)
	if firstSrcTsvPathNotExistErr != nil {
		return fmt.Errorf(
			"first LaunchShellPath not exist: %s",
			launchShellPath,
		)
	}
	return nil
}
