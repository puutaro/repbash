package maker

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/puutaro/repbash/mocks"
	"github.com/puutaro/repbash/pkg/val"
	"github.com/puutaro/repbash/testdata/testMethod"
	"path/filepath"
	"testing"
)

type testMakeLaunchShellStruct struct {
	testMakeLaunchShellInStruct
	testMakeLaunchShellWantStruct
}

type testMakeLaunchShellInStruct struct {
	testArgsInStruct            val.ArgumentsStruct
	testMainRepValMapIn         map[string]string
	testNewIoReaderInStruct     testNewIoReaderInStruct
	testRepValEnvForSearchsFact map[string]string
}

type testMakeLaunchShellWantStruct struct {
	wantResultShellCon string
	wantErr            error
}

type testNewIoReaderInStruct struct {
	testIoGetter IoGetter
}

var urlCon = `#!/bin/bash


function echoGit(){
echo git
}`

func createNormalIoGetter(
	ctrl *gomock.Controller,
) *mocks.MockIoGetter {
	newMockIoGetter := mocks.NewMockIoGetter(ctrl)
	newMockIoGetter.EXPECT().GetFromUrl(
		gomock.Any(),
	).Return(
		val.Result{
			GetCon: urlCon,
			Error:  nil,
		},
	)
	newMockIoGetter.EXPECT().GetFannelPath(
		gomock.Any(),
	).Return(
		filepath.Join(
			testMethod.GetTestDataDirPath(),
			"normal/shell/launch.sh",
		),
	)
	return newMockIoGetter
}

func createUrlConnectErrIoGetter(
	ctrl *gomock.Controller,
) *mocks.MockIoGetter {
	newMockIoGetter := mocks.NewMockIoGetter(ctrl)
	newMockIoGetter.EXPECT().GetFromUrl(
		gomock.Any(),
	).Return(
		val.Result{
			GetCon: "",
			Error:  fmt.Errorf("url connect err"),
		},
	)
	return newMockIoGetter
}

func TestMakeLaunchShell(t *testing.T) {
	testDataDirPath := testMethod.GetTestDataDirPath()
	ctrl := gomock.NewController(t)
	testMakeLaunchShellStructs := []struct {
		caseName                  string
		testMakeLaunchShellStruct testMakeLaunchShellStruct
	}{
		{"normal",
			testMakeLaunchShellStruct{
				testMakeLaunchShellInStruct{
					testArgsInStruct: val.ArgumentsStruct{
						LaunchShellPath: filepath.Join(testDataDirPath, "normal/shell/launch.sh"),
						SrcTsvPaths:     []string{filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable2.tsv")},
						ArgsCon:         "valName1=vv1,valName2=vv2",
						ArgsMap:         map[string]string{"valName1": "vv1", "valName2": "vv2"},
						ImportPaths: []string{
							"${IMPORT_PATH1}",
							"${IMPORT_PATH2}",
							"${IMPORT_PATH3}",
							"${IMPORT_GIT1}",
						},
						IsSaveRepbashLine: false,
					},
					testMainRepValMapIn: map[string]string{
						"FANNEL_PATH":                     filepath.Join(testDataDirPath, "normal/shell/launch.sh"),
						"IMPORT_GIT1":                     "https://raw.githubusercontent.com/puutaro/CommandClick-Linux/master/echoFromGit.sh",
						"IMPORT_PATH1":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell1.sh"),
						"IMPORT_PATH2":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell2.sh"),
						"IMPORT_PATH3":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell3.sh"),
						"REPLACE_VARIABLE_TABLE_TSV_PATH": filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable.tsv"),
						"TXT_LABEL":                       "label",
						"cmdYoutubePlayerBlankVal":        "",
						"cmdMusicPlayerDirListFilePath":   filepath.Join("/home/dummy/dir/path/list/music_dir_list"),
						"cmdMusicPlayerDirPath":           "/home/dummy/dir/path",
						"cmdMusicPlayerEditDirPath":       "/home/dummy/dir/path/edit",
						"cmdMusicPlayerListDirPath":       "/home/dummy/dir/path/list",
						"valName1":                        "vv1",
						"valName2":                        "vv2",
					},
					testNewIoReaderInStruct: testNewIoReaderInStruct{
						createNormalIoGetter(
							ctrl,
						),
					},
					testRepValEnvForSearchsFact: testMethod.UbuntuEnvMap,
				},
				testMakeLaunchShellWantStruct{
					wantResultShellCon: testMethod.ReadFileForTest(
						filepath.Join(testDataDirPath, "normal/case1/wants/want.sh"),
					),
					wantErr: nil,
				},
			},
		},
		{"no exist launchShellPath",
			testMakeLaunchShellStruct{
				testMakeLaunchShellInStruct{
					testArgsInStruct: val.ArgumentsStruct{
						LaunchShellPath:   filepath.Join(testDataDirPath, "normal/shell/launch99.sh"),
						SrcTsvPaths:       []string{},
						ArgsMap:           map[string]string{},
						ImportPaths:       []string{},
						IsSaveRepbashLine: false,
					},
					testMainRepValMapIn: map[string]string{},
					testNewIoReaderInStruct: testNewIoReaderInStruct{
						NewIoGetter(),
					},
					testRepValEnvForSearchsFact: testMethod.UbuntuEnvMap,
				},
				testMakeLaunchShellWantStruct{
					wantResultShellCon: "",
					wantErr: fmt.Errorf(
						"open %s: no such file or directory",
						filepath.Join(testDataDirPath, "normal/shell/launch99.sh"),
					),
				},
			},
		},
		{"no src tsv path",
			testMakeLaunchShellStruct{
				testMakeLaunchShellInStruct{
					testArgsInStruct: val.ArgumentsStruct{
						LaunchShellPath: filepath.Join(testDataDirPath, "normal/shell/launch.sh"),
						SrcTsvPaths:     []string{filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable2.tsv")},
						ArgsMap:         map[string]string{"valName1": "vv1", "valName2": "vv2"},
						ImportPaths: []string{
							"${IMPORT_PATH1}",
							"${IMPORT_PATH2}",
							"${IMPORT_PATH3}",
							"${IMPORT_GIT1}",
						},
						IsSaveRepbashLine: false,
					},
					testMainRepValMapIn: map[string]string{
						"FANNEL_PATH":                     filepath.Join(testDataDirPath, "normal/shell/launch.sh"),
						"IMPORT_GIT1":                     "https://raw.githubusercontent.com/puutaro/CommandClick-Linux/master/echoFromGit.sh",
						"IMPORT_PATH1":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell99.sh"),
						"IMPORT_PATH2":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell2.sh"),
						"IMPORT_PATH3":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell3.sh"),
						"REPLACE_VARIABLE_TABLE_TSV_PATH": filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable2.tsv"),
						"TXT_LABEL":                       "label",
						"cmdMusicPlayerDirListFilePath":   filepath.Join("/home/dummy/dir/path/list/music_dir_list"),
						"cmdMusicPlayerDirPath":           "/home/dummy/dir/path",
						"cmdMusicPlayerEditDirPath":       "/home/dummy/dir/path/edit",
						"cmdMusicPlayerListDirPath":       "/home/dummy/dir/path/list",
						"valName1":                        "vv1",
						"valName2":                        "vv2",
					},
					testNewIoReaderInStruct: testNewIoReaderInStruct{
						NewIoGetter(),
					},
					testRepValEnvForSearchsFact: testMethod.UbuntuEnvMap,
				},
				testMakeLaunchShellWantStruct{
					wantResultShellCon: "",
					wantErr: fmt.Errorf(
						"open %s: no such file or directory",
						filepath.Join(testDataDirPath, "normal/case1/facts/importShell99.sh"),
					),
				},
			},
		},
		{"url connect err",
			testMakeLaunchShellStruct{
				testMakeLaunchShellInStruct{
					testArgsInStruct: val.ArgumentsStruct{
						LaunchShellPath: filepath.Join(testDataDirPath, "normal/shell/launch.sh"),
						SrcTsvPaths:     []string{filepath.Join(testDataDirPath, "/normal/case1/facts/replaceVariablesTable2.tsv")},
						ArgsMap:         map[string]string{"valName1": "vv1", "valName2": "vv2"},
						ImportPaths: []string{
							"${IMPORT_PATH1}",
							"${IMPORT_PATH2}",
							"${IMPORT_PATH3}",
							"${IMPORT_GIT1}",
						},
						IsSaveRepbashLine: false,
					},
					testMainRepValMapIn: map[string]string{
						"FANNEL_PATH":                     filepath.Join(testDataDirPath, "normal/shell/launch.sh"),
						"IMPORT_GIT1":                     "https://raw.githubusercontent.com/puutaro/CommandClick-Linux/master/echoFromGit.sh",
						"IMPORT_PATH1":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell1.sh"),
						"IMPORT_PATH2":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell2.sh"),
						"IMPORT_PATH3":                    filepath.Join(testDataDirPath, "normal/case1/facts/importShell3.sh"),
						"REPLACE_VARIABLE_TABLE_TSV_PATH": filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable2.tsv"),
						"TXT_LABEL":                       "label",
						"cmdMusicPlayerDirListFilePath":   filepath.Join("/home/dummy/dir/path/list/music_dir_list"),
						"cmdMusicPlayerDirPath":           "/home/dummy/dir/path",
						"cmdMusicPlayerEditDirPath":       "/home/dummy/dir/path/edit",
						"cmdMusicPlayerListDirPath":       "/home/dummy/dir/path/list",
						"valName1":                        "vv1",
						"valName2":                        "vv2",
					},
					testNewIoReaderInStruct: testNewIoReaderInStruct{
						createUrlConnectErrIoGetter(ctrl),
					},
					testRepValEnvForSearchsFact: testMethod.UbuntuEnvMap,
				},
				testMakeLaunchShellWantStruct{
					wantResultShellCon: "",
					wantErr:            fmt.Errorf("url connect err"),
				},
			},
		},
	}
	for _, testMakeLaunchShellStructsEl := range testMakeLaunchShellStructs {
		t.Run(testMakeLaunchShellStructsEl.caseName, func(t *testing.T) {
			testMakeLaunchShellStruct := testMakeLaunchShellStructsEl.testMakeLaunchShellStruct
			testMakeLaunchShellInStruct := testMakeLaunchShellStruct.testMakeLaunchShellInStruct
			testArgsInStruct := testMakeLaunchShellInStruct.testArgsInStruct
			testRepValEnvForSearchsFact := testMakeLaunchShellInStruct.testRepValEnvForSearchsFact
			testMainRepValMapIn := testMakeLaunchShellInStruct.testMainRepValMapIn
			testIoGetterIn := testMakeLaunchShellStructsEl.testMakeLaunchShellStruct.testNewIoReaderInStruct.testIoGetter
			testMakeLaunchShellWantStruct := testMakeLaunchShellStruct.testMakeLaunchShellWantStruct
			wantResultShellCon := testMakeLaunchShellWantStruct.wantResultShellCon
			wantErr := testMakeLaunchShellWantStruct.wantErr
			testMethod.SetupEnv(t, testRepValEnvForSearchsFact)
			getResultShellCon, getErr := MakeLaunchShell(
				testIoGetterIn,
				testArgsInStruct,
				testMainRepValMapIn,
			)
			if !testMethod.IsConEqual(getResultShellCon, wantResultShellCon) ||
				!testMethod.IsErrEqual(getErr, wantErr) {
				t.Errorf(
					"getResultShellCon: %v,\n want %v\n",
					getResultShellCon,
					wantResultShellCon,
				)
				t.Errorf(
					"getErr: %v\n, want %v\n",
					getErr,
					wantErr,
				)
			} else {
				t.Logf("test %s ok", testMakeLaunchShellStructsEl.caseName)
			}
		})
	}
}
