package main

import (
	"fmt"
	"strings"

	flag "github.com/spf13/pflag"
)

type runtimeOpts struct {
	algorithm string
	verbose   bool
}

type outputOpts struct {
	format   string
	helpFlag bool
}

type inputOpts struct {
	inputType string
	args      []string
}

type options struct {
	runtime *runtimeOpts
	output  *outputOpts
	input   *inputOpts
}

func isIn(originalValue string, set []string) error {
	value := strings.ToLower(originalValue)
	for _, element := range set {
		if element == value {
			return nil
		}
	}
	return fmt.Errorf("%s: unavailable value, availables: %v", originalValue, set)
}

func validateMultipleValues(values string, availableSet []string) error {
	splittedValues := strings.Split(values, ",")
	for _, value := range splittedValues {
		err := isIn(value, availableSet)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateAlgorithm(algorithms string) error {
	set := []string{"simpson", "jaccard", "dice", "cosine", "pearson", "euclidean", "manhattan", "chebyshev", "levenshtein"}
	return validateMultipleValues(algorithms, set)
}

func validateInputType(inputTypes string) error {
	return validateMultipleValues(inputTypes, []string{"string", "byte_file", "term_file", "json"})
}

func validateFormat(format string) error {
	return isIn(format, []string{"default", "json", "xml"})
}

func validateEachOpt(opts *options) error {
	data := []struct {
		value     string
		validator func(value string) error
	}{
		{opts.input.inputType, validateInputType},
		{opts.output.format, validateFormat},
		{opts.runtime.algorithm, validateAlgorithm},
	}
	for _, datum := range data {
		if err := datum.validator(datum.value); err != nil {
			return err
		}
	}
	return nil
}

func validate(opts *options) (*options, error) {
	if opts.output.helpFlag {
		return opts, nil
	}
	if len(opts.input.args) <= 1 {
		return nil, fmt.Errorf("two arguments are required at the least")
	}
	if err := validateEachOpt(opts); err != nil {
		return nil, err
	}
	types := strings.Split(opts.input.inputType, ",")
	if len(types) != 1 && (len(types) != len(opts.input.args)) {
		return nil, fmt.Errorf("%s: input type must be one or same length with args (%d != %d)", opts.input.inputType, len(types), len(opts.input.args))
	}
	return opts, nil
}

func parseArgs(args []string) (*options, error) {
	opts := &options{runtime: &runtimeOpts{}, input: &inputOpts{}, output: &outputOpts{}}
	flags := flag.NewFlagSet("scv", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.StringVarP(&opts.runtime.algorithm, "algorithm", "a", "unknown", "specifies the calculating algorithm.")
	flags.BoolVarP(&opts.runtime.verbose, "verbose", "V", false, "runs on verbose mode.")
	flags.StringVarP(&opts.output.format, "format", "f", "default", "specifies the output format.")
	flags.StringVarP(&opts.input.inputType, "input-type", "t", "string", "specifies the type of VECTORS.")
	flags.BoolVarP(&opts.output.helpFlag, "help", "h", false, "prints this message")
	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	opts.input.args = flags.Args()[1:]
	return validate(opts)
}
