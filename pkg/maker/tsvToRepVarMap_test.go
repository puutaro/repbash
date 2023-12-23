package maker

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/puutaro/repbash/mocks"
	"github.com/puutaro/repbash/testdata/testMethod"
	"path/filepath"
	"reflect"
	"testing"
)

type testStrForReplaceRepValRecursively struct {
	fact []string
	want map[string]string
}

func TestReplaceRepValRecursively(t *testing.T) {
	testForIsEmptys := []struct {
		caseName                           string
		testStrForReplaceRepValRecursively testStrForReplaceRepValRecursively
	}{
		{
			"normal",
			testStrForReplaceRepValRecursively{
				[]string{
					"valName1\tvalValue1",
					"valName2\t${valValue1}_valValue2",
					"valName3\t${valName2}_valValue3",
				},
				map[string]string{
					"valName1": "valValue1",
					"valName2": "valValue1_valValue2",
					"valName3": "valValue1_valValue2_valValue3",
				},
			},
		},
	}
	for _, testForReplaceRepValRecursively := range testForIsEmptys {
		t.Run(testForReplaceRepValRecursively.caseName, func(t *testing.T) {
			testStrForReplaceRepValRecursively := testForReplaceRepValRecursively.testStrForReplaceRepValRecursively
			testStrForReplaceRepValRecursivelyFact := testStrForReplaceRepValRecursively.fact
			testStrForReplaceRepValRecursivelyWant := testStrForReplaceRepValRecursively.want

			replaceRepValRecursively(testStrForReplaceRepValRecursivelyFact)
			if reflect.DeepEqual(
				testStrForReplaceRepValRecursivelyFact,
				testStrForReplaceRepValRecursivelyWant,
			) {
				t.Errorf(
					"testStrForReplaceRepValRecursivelyFact: %v, want %v",
					testStrForReplaceRepValRecursivelyFact,
					testStrForReplaceRepValRecursivelyWant,
				)
			} else {
				t.Logf("test %s ok", testForReplaceRepValRecursively.caseName)
			}
		})
	}
}

type testTsvToRepVarMapStruct struct {
	testTsvPathList
	testTsvReaderService
	testErrForReadTsv
}

type testTsvReaderService struct {
	tsvReadService tsvReadService
}

type testTsvPathList struct {
	fact string
	want map[string]string
	//want []string
}

type testErrForReadTsv struct {
	want error
}

func createNewTsvReader(
	ctrl *gomock.Controller,
	tsvPath string,
) *mocks.MockTsvReader {
	newTsvReaderMock := mocks.NewMockTsvReader(ctrl)
	newTsvReaderMock.EXPECT().ReadTsv(tsvPath).Return(
		"", fmt.Errorf(
			"read err %s\n",
			tsvPath,
		),
	)
	return newTsvReaderMock
}

func TestTsvToRepVarMap(t *testing.T) {
	testDataDirPath := testMethod.GetTestDataDirPath()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testStructsForMakeRepValSrcMap := []struct {
		caseName                 string
		testTsvToRepVarMapStruct testTsvToRepVarMapStruct
	}{
		{
			"normal",
			testTsvToRepVarMapStruct{
				testTsvPathList{
					fact: filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable.tsv"),
					want: map[string]string{
						"FANNEL_PATH":                        filepath.Join(testDataDirPath, "normal/shell/launch.sh"),
						"IMPORT_GIT1":                        "https://raw.githubusercontent.com/puutaro/CommandClick-Linux/master/echoFromGit.sh",
						"IMPORT_PATH1":                       filepath.Join(testDataDirPath, "normal/case1/facts/importShell1.sh"),
						"IMPORT_PATH2":                       filepath.Join(testDataDirPath, "normal/case1/facts/importShell2.sh"),
						"IMPORT_PATH3":                       filepath.Join(testDataDirPath, "normal/case1/facts/importShell3.sh"),
						"REPLACE_VARIABLE_TABLE_TSV_PATH":    filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable2.tsv"),
						"TXT_LABEL":                          "label",
						"cmdMusicPlayerDirListFilePath":      filepath.Join("/home/dummy/dir/path/list/music_dir_list"),
						"cmdMusicPlayerDirPath":              "/home/dummy/dir/path",
						"cmdMusicPlayerEditDirPath":          "/home/dummy/dir/path/edit",
						"cmdMusicPlayerListDirPath":          "/home/dummy/dir/path/list",
						"cmdMusicPlayerSmallDirListFilePath": "/home/dummy/dir/path/list/smallList/small_list",
						"cmdMusicPlayerSmallListDirPath":     "/home/dummy/dir/path/list/smallList",
					},
				},
				testTsvReaderService{
					tsvReadService{reader: NewTsvReader()},
				},
				testErrForReadTsv{
					want: nil,
				},
			},
		},
		{
			"not found err",
			testTsvToRepVarMapStruct{
				testTsvPathList{
					fact: filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable22.tsv"),
					want: make(map[string]string),
				},
				testTsvReaderService{
					tsvReadService{reader: NewTsvReader()},
				},
				testErrForReadTsv{
					want: fmt.Errorf("not found %s\n", filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable22.tsv")),
				},
			},
		},
		{
			"read err",
			testTsvToRepVarMapStruct{
				testTsvPathList{
					fact: filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable22.tsv"),
					want: make(map[string]string),
				},
				testTsvReaderService{
					tsvReadService{
						createNewTsvReader(
							ctrl,
							filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable22.tsv"),
						),
					},
				},
				testErrForReadTsv{
					want: fmt.Errorf("read err %s\n", filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable22.tsv")),
				},
			},
		},
	}
	for _, testStructForTsvToRepVarMap := range testStructsForMakeRepValSrcMap {
		t.Run(testStructForTsvToRepVarMap.caseName, func(t *testing.T) {
			testTsvToRepVarMapStruct := testStructForTsvToRepVarMap.testTsvToRepVarMapStruct
			testTsvPathList := testTsvToRepVarMapStruct.testTsvPathList
			testTsvPathListFact := testTsvPathList.fact
			repValMapWant := testTsvPathList.want
			tsvReaderForTest := testStructForTsvToRepVarMap.testTsvToRepVarMapStruct.testTsvReaderService.tsvReadService
			testErrForReadTsv := testTsvToRepVarMapStruct.testErrForReadTsv
			testErrForReadTsvWant := testErrForReadTsv.want
			getRepValMap, getErr := TsvToRepVarMap(
				testTsvPathListFact,
				tsvReaderForTest,
			)
			if !reflect.DeepEqual(getRepValMap, repValMapWant) ||
				!testMethod.IsErrEqual(getErr, testErrForReadTsvWant) {
				t.Errorf(
					"testErrForReadTsvFact: %v,\n want %v\n",
					getRepValMap,
					repValMapWant,
				)
				t.Errorf(
					"getErr: %v\n, want %v\n",
					getErr,
					testErrForReadTsvWant,
				)
			} else {
				t.Logf("test %s ok", testStructForTsvToRepVarMap.caseName)
			}
		})
	}
}
