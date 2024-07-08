package searcher

import (
	"errors"
	"fmt"
	"github.com/puutaro/repbash/pkg/util/stringMethod"
	"github.com/puutaro/repbash/testdata/testMethod"
	"path/filepath"
	"testing"
)

type testSearch struct {
	testSearchPair
	testRepValEnvPair
	testErrorPair
}

type testSearchPair struct {
	fact string
	want string
}

type testRepValEnvPair struct {
	fact string
	want string
}

type testErrorPair struct {
	want error
}

func TestSearch(t *testing.T) {
	testDataDirPath := testMethod.GetTestDataDirPath()
	testSearchs := []struct {
		caseName   string
		testSearch testSearch
	}{
		{
			"normal",
			testSearch{
				testSearchPair{
					filepath.Join(testDataDirPath, "normal/shell"),
					filepath.Join(testDataDirPath, "normal/case1/facts/replaceVariablesTable.tsv"),
				},
				testRepValEnvPair{
					"settingVariables/replaceVariablesTable.tsv",
					"settingVariables/replaceVariablesTable.tsv",
				},
				testErrorPair{
					nil,
				},
			},
		},
		{
			"no rep val tsv",
			testSearch{
				testSearchPair{
					filepath.Join(testDataDirPath, "normal/case1/facts/fact2/fact3/fact4/fact5/fact6/fact7"),
					"",
				},
				testRepValEnvPair{
					"settingVariables/replaceVariablesTable.tsv",
					"settingVariables/replaceVariablesTable.tsv",
				},
				testErrorPair{
					nil,
				},
			},
		},
		{
			"rep val env no exist",
			testSearch{
				testSearchPair{
					testDataDirPath,
					"",
				},
				testRepValEnvPair{
					"",
					"",
				},
				testErrorPair{
					errors.New("not found env: REPLACE_VARIABLES_TSV_RELATIVE_PATH"),
				},
			},
		},
		{
			"launch dir path empty case",
			testSearch{
				testSearchPair{
					"",
					"",
				},
				testRepValEnvPair{
					"settingVariables/replaceVariablesTable.tsv",
					"settingVariables/replaceVariablesTable.tsv",
				},
				testErrorPair{
					errors.New("stat : no such file or directory"),
				},
			},
		},
		{
			"launch dir path no exist case",
			testSearch{
				testSearchPair{
					"/aa/bb",
					"",
				},
				testRepValEnvPair{
					"settingVariables/replaceVariablesTable.tsv",
					"settingVariables/replaceVariablesTable.tsv",
				},
				testErrorPair{
					fmt.Errorf(
						"stat %s: no such file or directory",
						"/aa/bb",
					),
				},
			},
		},
	}
	for _, testSearch := range testSearchs {
		t.Run(testSearch.caseName, func(t *testing.T) {
			testStrForSearchs := testSearch.testSearch
			testSearchPair := testStrForSearchs.testSearchPair
			testStrForSearchsFact := testSearchPair.fact
			testStrForSearchsWant := testSearchPair.want
			testErrorPair := testStrForSearchs.testErrorPair
			testErrorPairWant := testErrorPair.want
			testRepValEnvPair := testStrForSearchs.testRepValEnvPair
			testRepValEnvForSearchsFact := testRepValEnvPair.fact
			testRepValEnvForSearchsWant := testRepValEnvPair.want
			envMap := make(map[string]string)
			envMap["REPLACE_VARIABLES_TSV_RELATIVE_PATH"] = testRepValEnvForSearchsFact
			if !stringMethod.IsEmpty(testRepValEnvForSearchsFact) {
				testMethod.SetupEnv(t, envMap)
			}
			testStrForSearchGet, getErr := Search(testStrForSearchsFact)
			if testStrForSearchGet != testStrForSearchsWant ||
				testRepValEnvForSearchsFact != testRepValEnvForSearchsWant ||
				!testMethod.IsErrEqual(getErr, testErrorPairWant) {
				t.Errorf(
					"testStrForSearchGet: %v, want %v",
					testStrForSearchGet,
					testStrForSearchsWant,
				)
				t.Errorf(
					"testRepValEnvForSearchsFact %v, testRepValEnvForSearchsWant %v",
					testRepValEnvForSearchsFact,
					testRepValEnvForSearchsWant,
				)
				t.Errorf(
					"getErr:%v, want %v",
					getErr,
					testErrorPairWant,
				)
			} else {
				t.Logf("test %s ok", testSearch.caseName)
			}
		})
	}
}
