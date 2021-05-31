package main

func Example_help() {
	goMain([]string{"/some/path/of/scv", "-h"})
	// Output:
	// scv [OPTIONS] <VECTORS...>
	// OPTIONS
	//     -a, --algorithm <ALGORITHM>    specifies the calculating algorithm.  This option is mandatory.
	//                                    The value of this option accepts several values separated with comma.
	//                                    Available values are: simpson, jaccard, dice, and cosine.
	//     -f, --format <FORMAT>          specifies the resultant format. Default is default.
	//                                    Available values are: default, json, and xml.
	//     -t, --input-type <TYPE>        specifies the type of VECTORS. Default is file.
	//                                    If TYPE is separated with comma, each type shows
	//                                    the corresponding VECTORS.
	//                                    Available values are: file, string, and json.
	//     -h, --help                     prints this message.
	// VECTORS
	//     the source of vectors for calculation.
}
