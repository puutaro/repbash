package maker

import (
	"fmt"
	"github.com/puutaro/repbash/testdata/testMethod"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

type testMakeMainRepValMapStruct struct {
	testInStructForMakeMainRepValMap
	testMakeMainMapWantStruct
}

type testInStructForMakeMainRepValMap struct {
	testSrcTsvPathsStruct testSrcTsvPathsStruct
	testEnvStruct         testEnvStruct
}

type testEnvStruct struct {
	in map[string]string
}

type testSrcTsvPathsStruct struct {
	in []string
}

type testMakeMainMapWantStruct struct {
	wantMainRepValMap map[string]string
	wantErr           error
}

type testTsvReader struct{}

func (tr testTsvReader) ReadTsv(tsvPath string) (string, error) {
	tsvConSrc, err := readFile(
		tsvPath,
	)
	if err != nil {
		return "", err
	}
	testDataDirPath := testMethod.GetTestDataDirPath()
	testDataDirPathRegex := regexp.MustCompile("\t.*/testdata/")
	return testDataDirPathRegex.ReplaceAllString(
		tsvConSrc, fmt.Sprintf("\t%s/", testDataDirPath),
	), nil
}

func TestMakeMainRepValMap(t *testing.T) {
	testDataDirPath := testMethod.GetTestDataDirPath()
	testMakeMainRepValMapStructs := []struct {
		caseName                    string
		testMakeMainRepValMapStruct testMakeMainRepValMapStruct
	}{
		{
			"normal",
			testMakeMainRepValMapStruct{
				testInStructForMakeMainRepValMap{
					testSrcTsvPathsStruct: testSrcTsvPathsStruct{
						in: []string{
							filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable.tsv"),
						},
					},
					testEnvStruct: testEnvStruct{
						in: testMethod.UbuntuEnvMap,
					},
				},
				testMakeMainMapWantStruct{
					wantMainRepValMap: map[string]string{
						"FANNEL_PATH":                        filepath.Join(testDataDirPath, "normal/shell/launch.sh"),
						"IMPORT_GIT1":                        "https://raw.githubusercontent.com/puutaro/CommandClick-Linux/master/echoFromGit.sh",
						"IMPORT_PATH1":                       filepath.Join(testDataDirPath, "normal/case1/facts/importShell1.sh"),
						"IMPORT_PATH2":                       filepath.Join(testDataDirPath, "normal/case1/facts/importShell2.sh"),
						"IMPORT_PATH3":                       filepath.Join(testDataDirPath, "normal/case1/facts/importShell3.sh"),
						"REPLACE_VARIABLE_TABLE_TSV_PATH":    filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable2.tsv"),
						"TXT_LABEL":                          "label",
						"UBUNTU_ENV_TSV_PATH":                "/home/xbabu/Desktop/share/android/cmds/repbash/testdata/normal/case1/facts/ubuntu_env.tsv",
						"cmdMusicPlayerDirListFilePath":      filepath.Join("/home/dummy/dir/path/list/music_dir_list"),
						"cmdMusicPlayerDirPath":              "/home/dummy/dir/path",
						"cmdMusicPlayerSmallListDirPath":     "/home/dummy/dir/path/list/smallList",
						"cmdMusicPlayerEditDirPath":          "/home/dummy/dir/path/edit",
						"cmdMusicPlayerListDirPath":          "/home/dummy/dir/path/list",
						"cmdMusicPlayerSmallDirListFilePath": "/home/dummy/dir/path/list/smallList/small_list",
					},
					wantErr: nil,
				},
			},
		},
		{
			"not found err",
			testMakeMainRepValMapStruct{
				testInStructForMakeMainRepValMap{
					testSrcTsvPathsStruct: testSrcTsvPathsStruct{
						in: []string{
							filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable99.tsv"),
						},
					},
					testEnvStruct: testEnvStruct{
						in: testMethod.UbuntuEnvMap,
					},
				},
				testMakeMainMapWantStruct{
					wantMainRepValMap: map[string]string{},
					wantErr: fmt.Errorf(
						"not found %s\n",
						filepath.Join(
							testDataDirPath,
							"normal/case1/facts/replaceVariablesTable99.tsv",
						),
					),
				},
			},
		},
	}
	for _, testMakeMainRepValMapStructEl := range testMakeMainRepValMapStructs {
		t.Run(testMakeMainRepValMapStructEl.caseName, func(t *testing.T) {
			testMakeMainRepValMapStruct := testMakeMainRepValMapStructEl.testMakeMainRepValMapStruct
			testEnvStructIn := testMakeMainRepValMapStruct.testEnvStruct.in
			testSrcTsvPathsIn := testMakeMainRepValMapStruct.testSrcTsvPathsStruct.in
			testMakeMainMapWantStruct := testMakeMainRepValMapStruct.testMakeMainMapWantStruct
			wantMainRepValMap := testMakeMainMapWantStruct.wantMainRepValMap
			wantErr := testMakeMainMapWantStruct.wantErr
			testMethod.SetupEnv(t, testEnvStructIn)
			getMainRepValMap, getErr := MakeMainRepValMap(
				testSrcTsvPathsIn,
				testTsvReader{},
			)
			if !reflect.DeepEqual(getMainRepValMap, wantMainRepValMap) ||
				!testMethod.IsErrEqual(getErr, wantErr) {
				t.Errorf(
					"getMainRepValMap: %v,\n want %v\n",
					getMainRepValMap,
					wantMainRepValMap,
				)
				t.Errorf(
					"getErr: %v\n, want %v\n",
					getErr,
					wantErr,
				)
			} else {
				t.Logf("test %s ok", testMakeMainRepValMapStructEl.caseName)
			}
		})
	}
}
