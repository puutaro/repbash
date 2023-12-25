package args

import (
	"fmt"
	"github.com/puutaro/repbash/testdata/testMethod"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type testArgsStruct struct {
	launchShellPathTestPair   launchShellPathTestPair
	srcTsvPathsTestPair       srcTsvPathsTestPair
	argsMapTestPair           argsMapTestPair
	isSaveRepbashLineTestPair isSaveRepbashLineTestPair
	importPathsTestPair       importPathsTestPair
	errTestPair               errTestPair
}

type launchShellPathTestPair struct {
	fact string
	want string
}

type srcTsvPathsTestPair struct {
	facts []string
	wants []string
}

type argsMapTestPair struct {
	facts map[string]string
	wants map[string]string
}

type errTestPair struct {
	want error
}

type isSaveRepbashLineTestPair struct {
	fact bool
	want bool
}

type importPathsTestPair struct {
	fact []string
	want []string
}

func split(s string) []string {
	return strings.Split(s, " ")
}

func makeCmdLine(
	launchShellPathFact string,
	srcTsvPathFacts []string,
	argsMapFacts map[string]string,
	isSaveRepbashLineFacts bool,
	importPathsFacts []string,
) (testCmds []string) {
	testCmds = []string{
		"./main.go",
		launchShellPathFact,
	}
	if len(srcTsvPathFacts) > 0 {
		testCmds = append(testCmds, "-t")
		testCmds = append(testCmds, srcTsvPathFacts...)
	}
	if len(argsMapFacts) > 0 {
		testCmds = append(testCmds, "-a")
		testCmds = append(testCmds, setArgsFromMap(argsMapFacts))
	}
	if isSaveRepbashLineFacts {
		testCmds = append(testCmds, "-s")
	}
	if len(importPathsFacts) > 0 {
		testCmds = append(testCmds, "-i")
		testCmds = append(testCmds, importPathsFacts...)
	}
	return
}

func setArgsFromMap(argsMapFacts map[string]string) string {
	var argsList = []string{}
	for k, v := range argsMapFacts {
		argsList = append(argsList, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(argsList, ",")

}

func setup(t *testing.T, envs map[string]string) {
	for k, v := range envs {
		t.Setenv(k, v)
	}
}

func TestArgs(t *testing.T) {

	testDataDirPath := testMethod.GetTestDataDirPath()
	testsArgs := []struct {
		caseName       string
		testArgsStruct testArgsStruct
	}{
		{
			caseName: "normal case",
			testArgsStruct: testArgsStruct{
				launchShellPathTestPair: launchShellPathTestPair{
					fact: filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable.tsv"),
					want: filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable.tsv"),
				},
				srcTsvPathsTestPair: srcTsvPathsTestPair{
					facts: []string{filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable2.tsv")},
					wants: []string{filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable2.tsv")},
				},
				argsMapTestPair: argsMapTestPair{
					facts: map[string]string{"valName1": "vv1", "valName2": "vv2"},
					wants: map[string]string{"valName1": "vv1", "valName2": "vv2"},
				},
				importPathsTestPair: importPathsTestPair{
					fact: []string{
						"${IMPORT_PATH1}",
						"${IMPORT_PATH2}",
						"${IMPORT_PATH3}",
					},
					want: []string{
						"${IMPORT_PATH1}",
						"${IMPORT_PATH2}",
						"${IMPORT_PATH3}",
					},
				},
				isSaveRepbashLineTestPair: isSaveRepbashLineTestPair{
					fact: false,
					want: false,
				},
				errTestPair: errTestPair{
					want: nil,
				},
			},
		},
		{
			caseName: "optional blank case",
			testArgsStruct: testArgsStruct{
				launchShellPathTestPair: launchShellPathTestPair{
					fact: filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable.tsv"),
					want: filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable.tsv"),
				},
				srcTsvPathsTestPair: srcTsvPathsTestPair{
					facts: []string{
						filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable2.tsv"),
					},
					wants: []string{
						filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable2.tsv"),
					},
				},
				argsMapTestPair: argsMapTestPair{
					facts: map[string]string{},
					wants: map[string]string{},
				},
				importPathsTestPair: importPathsTestPair{
					fact: []string{
						filepath.Join(testDataDirPath, "/normal/case1/facts/importShell1.sh"),
						filepath.Join(testDataDirPath, "/normal/case1/facts/importShell2.sh"),
						filepath.Join(testDataDirPath, "/normal/case1/facts/importShell3.sh"),
					},
					want: []string{
						filepath.Join(testDataDirPath, "/normal/case1/facts/importShell1.sh"),
						filepath.Join(testDataDirPath, "/normal/case1/facts/importShell2.sh"),
						filepath.Join(testDataDirPath, "/normal/case1/facts/importShell3.sh"),
					},
				},
				isSaveRepbashLineTestPair: isSaveRepbashLineTestPair{
					fact: false,
					want: false,
				},
				errTestPair: errTestPair{
					want: nil,
				},
			},
		},
		{
			caseName: "err no exist launchShellPath  case",
			testArgsStruct: testArgsStruct{
				launchShellPathTestPair: launchShellPathTestPair{
					fact: filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable99.tsv"),
					want: filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable99.tsv"),
				},
				srcTsvPathsTestPair: srcTsvPathsTestPair{
					facts: []string{},
					wants: []string{},
				},
				argsMapTestPair: argsMapTestPair{
					facts: map[string]string{},
					wants: map[string]string{},
				},
				importPathsTestPair: importPathsTestPair{
					fact: []string{},
					want: []string{},
				},
				isSaveRepbashLineTestPair: isSaveRepbashLineTestPair{
					fact: false,
					want: false,
				},
				errTestPair: errTestPair{
					want: fmt.Errorf(
						"first LaunchShellPath not exist: %s",
						filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable99.tsv"),
					),
				},
			},
		},
		{
			caseName: "err blank launchShellPath  case",
			testArgsStruct: testArgsStruct{
				launchShellPathTestPair: launchShellPathTestPair{
					fact: "",
					want: "",
				},
				srcTsvPathsTestPair: srcTsvPathsTestPair{
					facts: []string{},
					wants: []string{},
				},
				argsMapTestPair: argsMapTestPair{
					facts: map[string]string{},
					wants: map[string]string{},
				},
				importPathsTestPair: importPathsTestPair{
					fact: []string{},
					want: []string{},
				},
				isSaveRepbashLineTestPair: isSaveRepbashLineTestPair{
					fact: false,
					want: false,
				},
				errTestPair: errTestPair{
					want: fmt.Errorf(
						"first LaunchShellPath not found",
					),
				},
			},
		},
		{
			caseName: "err no exist srcTsvPathsTestPair  case",
			testArgsStruct: testArgsStruct{
				launchShellPathTestPair: launchShellPathTestPair{
					fact: filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable.tsv"),
					want: filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable.tsv"),
				},
				srcTsvPathsTestPair: srcTsvPathsTestPair{
					facts: []string{
						filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable99.tsv"),
					},
					wants: []string{
						filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable99.tsv"),
					},
				},
				argsMapTestPair: argsMapTestPair{
					facts: map[string]string{},
					wants: map[string]string{},
				},
				importPathsTestPair: importPathsTestPair{
					fact: []string{},
					want: []string{},
				},
				isSaveRepbashLineTestPair: isSaveRepbashLineTestPair{
					fact: false,
					want: false,
				},
				errTestPair: errTestPair{
					want: nil,
				},
			},
		},
	}
	for _, tt := range testsArgs {
		t.Run(tt.caseName, func(t *testing.T) {
			testArgsSet := tt.testArgsStruct
			launchShellPathTestPair := testArgsSet.launchShellPathTestPair
			launchShellPathFact := launchShellPathTestPair.fact
			launchShellPathWant := launchShellPathTestPair.want
			srcTsvPathsTestPair := testArgsSet.srcTsvPathsTestPair
			srcTsvPathFacts := srcTsvPathsTestPair.facts
			srcTsvPathWants := srcTsvPathsTestPair.wants
			argsMapTestPair := testArgsSet.argsMapTestPair
			argsMapFacts := argsMapTestPair.facts
			argsMapWants := argsMapTestPair.wants
			isSaveRepbashLineTestPair := testArgsSet.isSaveRepbashLineTestPair
			isSaveRepbashLineFacts := isSaveRepbashLineTestPair.fact
			isSaveRepbashLineWants := isSaveRepbashLineTestPair.want
			importPathsTestPair := testArgsSet.importPathsTestPair
			importPathsFacts := importPathsTestPair.fact
			importPathsWants := importPathsTestPair.want
			errTestPair := testArgsSet.errTestPair
			errTestWants := errTestPair.want
			//envMap := make(map[string]string)
			//envMap["REPLACE_VARIABLES_CONFIG"] = setReplaceVariablesJsFact
			//setup(t, envMap)
			os.Args = makeCmdLine(
				launchShellPathFact,
				srcTsvPathFacts,
				argsMapFacts,
				isSaveRepbashLineFacts,
				importPathsFacts,
			)
			getArgsStruct, getErr := Args()
			getLaunchShellPath := getArgsStruct.LaunchShellPath
			getSrcPaths := getArgsStruct.SrcTsvPaths
			getImportPaths := getArgsStruct.ImportPaths
			getArgsMap := getArgsStruct.ArgsMap
			getIsSaveRepbashLine := getArgsStruct.IsSaveRepbashLine
			if getLaunchShellPath != launchShellPathWant ||
				!reflect.DeepEqual(getSrcPaths, srcTsvPathWants) ||
				!reflect.DeepEqual(getImportPaths, importPathsWants) ||
				!reflect.DeepEqual(getArgsMap, argsMapWants) ||
				getIsSaveRepbashLine != isSaveRepbashLineWants ||
				!testMethod.IsErrEqual(getErr, errTestWants) {
				t.Errorf(
					"getLaunchShellPath: %v, want %v",
					getLaunchShellPath,
					launchShellPathWant,
				)
				t.Errorf(
					"getIsSaveRepbashLine: %v, want %v",
					getIsSaveRepbashLine,
					isSaveRepbashLineWants,
				)
				t.Errorf(
					"getSrcTsvPaths: %v, len %d, wants %v, len %d",
					strings.Join(getSrcPaths, ","),
					len(getSrcPaths),
					strings.Join(srcTsvPathWants, ","),
					len(srcTsvPathWants),
				)
				t.Errorf(
					"getImportPaths: %v, len %d, wants %v, len %d",
					strings.Join(getImportPaths, ","),
					len(getImportPaths),
					strings.Join(importPathsWants, ","),
					len(importPathsWants),
				)
				t.Errorf(
					"getArgsMap: %v, len %d, wants %v, len %d",
					getArgsMap,
					len(getArgsMap),
					argsMapWants,
					len(argsMapWants),
				)
				t.Errorf(
					"getError: %v, wants %v",
					getErr,
					errTestWants,
				)
			} else {
				t.Logf("test %s ok", tt.caseName)
			}
		})
	}
}
