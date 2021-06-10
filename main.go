package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tamada/scv/vector"
)

func helpMessage(originalProgramName string) string {
	name := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS] <VECTORS...>
OPTIONS
    -a, --algorithm <ALGORITHM>    specifies the calculating algorithm.  This option is mandatory.
                                   The value of this option accepts several values separated with comma.
                                   Available values are: simpson, jaccard, dice, cosine, pearson,
                                   euclidean, manhattan, chebyshev, and levenshtein.
    -f, --format <FORMAT>          specifies the resultant format. Default is default.
                                   Available values are: default, json, and xml.
    -t, --input-type <TYPE>        specifies the type of VECTORS. Default is string.
                                   If TYPE is separated with comma, each type shows
                                   the corresponding VECTORS.
                                   Available values are: file, string, and json.
    -h, --help                     prints this message.
VECTORS
    the source of vectors for calculation.`, name)
}

func convert(opts *options) []*vector.Vector {
	results := []*vector.Vector{}
	for _, arg := range opts.args {
		vector := vector.NewVectorFromString(arg)
		results = append(results, vector)
	}
	return results
}

func pairing(vectors []*vector.Vector) []*vector.VectorPair {
	pairs := []*vector.VectorPair{}
	for _, vector1 := range vectors {
		for _, vector2 := range vectors {
			if vector1 == vector2 {
				break
			}
			pairs = append(pairs, &vector.VectorPair{Vector1: vector2, Vector2: vector1})
		}
	}
	return pairs
}

type result struct {
	algorithm  string
	pair       *vector.VectorPair
	similarity float64
	err        error
}

func calculate(pairs []*vector.VectorPair, algorithm vector.Algorithm, name string) []*result {
	results := []*result{}
	for _, pair := range pairs {
		similarity, err := pair.Compare(algorithm)
		results = append(results, &result{algorithm: name, similarity: similarity, pair: pair, err: err})
	}
	return results
}

func constructPairs(opts *options) []*vector.VectorPair {
	vectors := convert(opts)
	return pairing(vectors)
}

func perform(opts *options) int {
	pairs := constructPairs(opts)
	algos := strings.Split(opts.algorithm, ",")
	printer := NewPrinter(opts.format, os.Stdout)
	printer.PrintHeader()
	for i, algorithmName := range algos {
		algorithm, err := vector.NewAlgorithm(algorithmName)
		if err != nil {
			fmt.Println(err.Error())
			return 3
		}
		results := calculate(pairs, algorithm, algorithmName)
		printEach(printer, results, i == 0)
	}
	printer.PrintFooter()
	return 0
}

func printEach(printer Printer, results []*result, first bool) int {
	for i := range results {
		printer.PrintEach(results[i], first && i == 0)
	}
	return 0
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Printf("%s: %s\n", filepath.Base(args[0]), err.Error())
		fmt.Println(helpMessage(args[0]))
		return 1
	}
	if opts.helpFlag {
		fmt.Println(helpMessage(args[0]))
		return 0
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
