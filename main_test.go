package main

import "testing"

func Example_similarities() {
	goMain([]string{"scv", "--algorithm", "jaccard,pearson", "I have a pen", "This is a pen"})
	// Output:
	// jaccard(I have a pen, This is a pen): 0.545455
	// pearson(I have a pen, This is a pen): 0.220433
}

func Example_help() {
	goMain([]string{"/some/path/of/scv", "-h"})
	// Output:
	// scv [OPTIONS] <VECTORS...>
	// OPTIONS
	//     -a, --algorithm <ALGORITHM>    specifies the calculating algorithm.  This option is mandatory.
	//                                    The value of this option accepts several values separated with comma.
	//                                    Available values are: simpson, jaccard, dice, cosine, pearson,
	//                                    euclidean, manhattan, chebyshev, and levenshtein.
	//     -f, --format <FORMAT>          specifies the resultant format. Default is default.
	//                                    Available values are: default, json, and xml.
	//     -t, --input-type <TYPE>        specifies the type of VECTORS. Default is string.
	//                                    If TYPE is separated with comma, each type shows
	//                                    the corresponding VECTORS.
	//                                    Available values are: file, string, and json.
	//     -h, --help                     prints this message.
	// VECTORS
	//     the source of vectors for calculation.
}

func TestParseArgs(t *testing.T) {
	testdata := []struct {
		giveArgs   []string
		wontStatus int
	}{
		{[]string{"scv", "--invalid-option", "a1", "a2"}, 1},
		{[]string{"scv"}, 1}, // required parameters missing
		{[]string{"scv", "--algorithm", "unknown_algorithm", "a1", "a2"}, 1},
		{[]string{"scv", "--algorithm", "simpson", "--format", "unknown_format", "a1", "a2"}, 1},
		{[]string{"scv", "--algorithm", "simpson", "--input-type", "unknown_type", "a1", "a2"}, 1},
		{[]string{"scv", "--algorithm", "simpson", "--input-type", "string", "a1", "a2"}, 0},
		{[]string{"scv", "--algorithm", "simpson", "--input-type", "string,string", "a1", "a2"}, 0},
		{[]string{"scv", "--algorithm", "simpson", "--input-type", "string,string", "a1", "a2", "a3"}, 1},
	}
	for _, td := range testdata {
		gotStatus := goMain(td.giveArgs)
		if gotStatus != td.wontStatus {
			t.Errorf("goMain(%v) status code did not match, wont %d, got %d", td.giveArgs, td.wontStatus, gotStatus)
		}
	}
}
