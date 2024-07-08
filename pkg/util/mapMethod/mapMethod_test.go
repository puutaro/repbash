package mapMethod

import (
	"github.com/puutaro/repbash/testdata/testMethod"
	"reflect"
	"testing"
)

type testConcatForMapStruct struct {
	testConcatForMapInStruct  testConcatForMapInStruct
	testConcatForMapOutStruct testConcatForMapOutStruct
}

type testConcatForMapInStruct struct {
	mainMapIn map[string]string
	addMapIn  map[string]string
}

type testConcatForMapOutStruct struct {
	want map[string]string
}

func TestConcat(t *testing.T) {
	testConcatForMapStructs := []struct {
		caseName               string
		testConcatForMapStruct testConcatForMapStruct
	}{
		{
			"normal",
			testConcatForMapStruct{
				testConcatForMapInStruct{
					mainMapIn: map[string]string{
						"valName1": "valValue1",
						"valName2": "valValue2",
					},
					addMapIn: map[string]string{
						"valName3": "valValue3",
					},
				},
				testConcatForMapOutStruct{
					want: map[string]string{
						"valName1": "valValue1",
						"valName2": "valValue2",
						"valName3": "valValue3",
					},
				},
			},
		},
		{
			"mainMapIn: nil",
			testConcatForMapStruct{
				testConcatForMapInStruct{
					mainMapIn: map[string]string{},
					addMapIn: map[string]string{
						"valName1": "valValue1",
						"valName2": "valValue2",
					},
				},
				testConcatForMapOutStruct{
					want: map[string]string{
						"valName1": "valValue1",
						"valName2": "valValue2",
					},
				},
			},
		},
		{
			"addMapIn: nil",
			testConcatForMapStruct{
				testConcatForMapInStruct{
					mainMapIn: map[string]string{
						"valName1": "valValue1",
						"valName2": "valValue2",
					},
					addMapIn: map[string]string{},
				},
				testConcatForMapOutStruct{
					want: map[string]string{
						"valName1": "valValue1",
						"valName2": "valValue2",
					},
				},
			},
		},
		{
			"mainMapIn: nil, mainMapIn: nil",
			testConcatForMapStruct{
				testConcatForMapInStruct{
					mainMapIn: map[string]string{},
					addMapIn:  map[string]string{},
				},
				testConcatForMapOutStruct{
					want: map[string]string{},
				},
			},
		},
	}
	for _, testConcatForMapStructsEl := range testConcatForMapStructs {
		t.Run(testConcatForMapStructsEl.caseName, func(t *testing.T) {
			testConcatForMapStruct := testConcatForMapStructsEl.testConcatForMapStruct
			testConcatForMapInStruct := testConcatForMapStruct.testConcatForMapInStruct

			mainMapIn := testConcatForMapInStruct.mainMapIn
			addMapIn := testConcatForMapInStruct.addMapIn
			want := testConcatForMapStruct.testConcatForMapOutStruct.want
			Concat(mainMapIn, addMapIn)
			if !reflect.DeepEqual(mainMapIn, want) {
				t.Errorf(
					"mainMapIn: %v\nmainMapIn %v\n want %v\n",
					mainMapIn,
					addMapIn,
					want,
				)
			} else {
				t.Logf("test %s ok", testConcatForMapStructsEl.caseName)
			}
		})
	}
}

type testReplaceConByMapStruct struct {
	tesReplaceConByMapForMapInStruct   tesReplaceConByMapForMapInStruct
	testReplaceConByMapForMapOutStruct testReplaceConByMapForMapOutStruct
}

type tesReplaceConByMapForMapInStruct struct {
	conIn     string
	mainMapIn map[string]string
}

type testReplaceConByMapForMapOutStruct struct {
	want string
}

func TestReplaceConByMap(t *testing.T) {
	testReplaceConByMapStructs := []struct {
		caseName                  string
		testReplaceConByMapStruct testReplaceConByMapStruct
	}{
		{
			"normal",
			testReplaceConByMapStruct{
				tesReplaceConByMapForMapInStruct{
					conIn: "replaceName1=${valName1}\nreplaceName2=${valName1}_${valName2}",
					mainMapIn: map[string]string{
						"valName1": "valValue1",
						"valName2": "valValue2",
					},
				},
				testReplaceConByMapForMapOutStruct{
					want: "replaceName1=valValue1\nreplaceName2=valValue1_valValue2",
				},
			},
		},
		{
			"mainMapIn: nil",
			testReplaceConByMapStruct{
				tesReplaceConByMapForMapInStruct{
					conIn:     "replaceName1=${valName1}\nreplaceName2=${valName1}_${valName1}",
					mainMapIn: map[string]string{},
				},
				testReplaceConByMapForMapOutStruct{
					want: "replaceName1=${valName1}\nreplaceName2=${valName1}_${valName1}",
				},
			},
		},
	}
	for _, testReplaceConByMapStructEl := range testReplaceConByMapStructs {
		t.Run(testReplaceConByMapStructEl.caseName, func(t *testing.T) {
			testReplaceConByMapStruct := testReplaceConByMapStructEl.testReplaceConByMapStruct
			tesReplaceConByMapForMapInStruct := testReplaceConByMapStruct.tesReplaceConByMapForMapInStruct
			mainMapIn := tesReplaceConByMapForMapInStruct.mainMapIn
			conIn := tesReplaceConByMapForMapInStruct.conIn
			want := testReplaceConByMapStruct.testReplaceConByMapForMapOutStruct.want
			ReplaceConByMap(
				&conIn,
				mainMapIn,
			)
			if !testMethod.IsConEqual(conIn, want) {
				t.Errorf(
					"conIn: %v\nmainMapIn %v\n want %v\n",
					conIn,
					mainMapIn,
					want,
				)
			} else {
				t.Logf("test %s ok", testReplaceConByMapStructEl.caseName)
			}
		})
	}
}
