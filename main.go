package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func helpMessage(originalProgramName string) string {
	name := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS] <VECTORS...>
OPTIONS
    -a, --algorithm <ALGORITHM>    specifies the calculating algorithm.  This option is mandatory.
                                   The value of this option accepts several values separated with comma.
                                   Available values are: simpson, jaccard, dice, and cosine.
    -f, --format <FORMAT>          specifies the resultant format. Default is default.
                                   Available values are: default, json, and xml.
    -t, --input-type <TYPE>        specifies the type of VECTORS. Default is file.
                                   If TYPE is separated with comma, each type shows
                                   the corresponding VECTORS.
                                   Available values are: file, string, and json.
    -h, --help                     prints this message.
VECTORS
    the source of vectors for calculation.`, name)
}

func goMain(args []string) int {
	fmt.Println(helpMessage(args[0]))
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
