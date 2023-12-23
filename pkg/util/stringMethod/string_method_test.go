package stringMethod

import (
	"testing"
)

type testStrForIsEmpty struct {
	fact string
	want bool
}

func TestIsEmpty(t *testing.T) {
	testForIsEmptys := []struct {
		caseName          string
		testStrForIsEmpty testStrForIsEmpty
	}{
		{
			"normal",
			testStrForIsEmpty{
				"",
				true,
			},
		},
		{
			"is space",
			testStrForIsEmpty{
				"  ",
				true,
			},
		},
		{
			"is zenkaku space",
			testStrForIsEmpty{
				" 　 　",
				true,
			},
		},
	}
	for _, testForIsEmpty := range testForIsEmptys {
		t.Run(testForIsEmpty.caseName, func(t *testing.T) {
			testStrForIsEmpty := testForIsEmpty.testStrForIsEmpty
			testStrForIsEmptyFact := testStrForIsEmpty.fact
			testStrForIsEmptyWant := testStrForIsEmpty.want
			testStrForIsEmptyGet := IsEmpty(testStrForIsEmptyFact)
			if testStrForIsEmptyWant != testStrForIsEmptyGet {
				t.Errorf(
					"testStrForIsEmptyGet: %v, want %v",
					testStrForIsEmptyGet,
					testStrForIsEmptyWant,
				)
			} else {
				t.Logf("test %s ok", testForIsEmpty.caseName)
			}
		})
	}
}

type testStrForTrimBothQuote struct {
	fact string
	want string
}

func TestTrimBothQuote(t *testing.T) {
	testForIsEmptys := []struct {
		caseName                string
		testStrForTrimBothQuote testStrForTrimBothQuote
	}{
		{
			"trim double quote",
			testStrForTrimBothQuote{
				`"aa bb"`,
				`aa bb`,
			},
		},
		{
			"trim double quote with space",
			testStrForTrimBothQuote{
				`　 "aa bb" `,
				`aa bb`,
			},
		},
		{
			"trim single quote",
			testStrForTrimBothQuote{
				`'aa bb'`,
				`aa bb`,
			},
		},
		{
			"trim single quote with double quote",
			testStrForTrimBothQuote{
				` 　'aa bb' 　`,
				`aa bb`,
			},
		},
		{
			"trim back quote",
			testStrForTrimBothQuote{
				"`aa bb`",
				`aa bb`,
			},
		},
		{
			"trim back quote",
			testStrForTrimBothQuote{
				" 　`aa bb` 　　 ",
				`aa bb`,
			},
		},
	}
	for _, testForTrimBothQuote := range testForIsEmptys {
		t.Run(testForTrimBothQuote.caseName, func(t *testing.T) {
			testStrForTrimBothQuote := testForTrimBothQuote.testStrForTrimBothQuote
			testStrForTrimBothQuoteFact := testStrForTrimBothQuote.fact
			testStrForTrimBothQuoteWant := testStrForTrimBothQuote.want
			testSrtForIsEmptyGet := testStrForTrimBothQuoteFact
			TrimBothQuote(&testSrtForIsEmptyGet)
			if testStrForTrimBothQuoteWant != testSrtForIsEmptyGet {
				t.Errorf(
					"testSrtForIsEmptyGet: %v, want %v",
					testSrtForIsEmptyGet,
					testStrForTrimBothQuoteWant,
				)
			} else {
				t.Logf("test %s ok", testForTrimBothQuote.caseName)
			}
		})
	}
}
