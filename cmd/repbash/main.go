package main

import (
	"fmt"
	"github.com/puutaro/repbash/internal/apps/repbash/pkg/args"
	"github.com/puutaro/repbash/pkg/maker"
	"github.com/puutaro/repbash/pkg/searcher"
	"github.com/puutaro/repbash/pkg/shell"
	"github.com/puutaro/repbash/pkg/util/mapMethod"
	"log"
	"os"
)

func main() {
	argsStruct, err := args.Args()
	srcTsvPaths := argsStruct.SrcTsvPaths
	argsMap := argsStruct.ArgsMap
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	replaceVariableTsvPath, err := searcher.Search(argsStruct.LaunchShellPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	srcTsvPaths = append([]string{replaceVariableTsvPath}, srcTsvPaths...)
	mainRepValMap, err := maker.MakeMainRepValMap(
		srcTsvPaths,
		maker.NewTsvReader(),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	mapMethod.Concat(
		mainRepValMap,
		argsMap,
	)

	shellCon, err := maker.MakeLaunchShell(
		maker.NewIoGetter(),
		argsStruct,
		mainRepValMap,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	out, errout, err := shell.ExecBashCommand(shellCon)
	if err != nil {
		log.Println(errout)
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(out)
	//fmt.Printf(shellCon)

}
