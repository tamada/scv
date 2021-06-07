package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tamada/scv/vector"
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
	fmt.Printf("convert(%v)\n", opts.args)
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
	similarity float32
	err        error
}

func calculate(pairs []*vector.VectorPair, algorithm vector.Algorithm) []*result {
	results := []*result{}
	for _, pair := range pairs {
		similarity := pair.Compare(algorithm)
		results = append(results, &result{similarity: similarity, err: nil})
	}
	return results
}

func printResult(algorithmName string, pair *vector.VectorPair, r *result) {
	fmt.Printf("%s(%s, %s): %f\n", algorithmName, pair.Vector1.Source.Value(), pair.Vector2.Source.Value(), r.similarity)
}

func perform(opts *options) int {
	vectors := convert(opts)
	pairs := pairing(vectors)
	algorithm, err := vector.NewAlgorithm(opts.algorithm)
	if err != nil {
		fmt.Println(err.Error())
		return 3
	}
	results := calculate(pairs, algorithm)
	for i, _ := range results {
		printResult(opts.algorithm, pairs[i], results[i])
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
