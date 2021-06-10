package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/tamada/scv/vector"
)

type Printer interface {
	PrintHeader()
	PrintEach(algorithmName string, v1, v2 *vector.Vector, result *result, first bool)
	PrintFooter()
}

func NewPrinter(printerName string, writer io.Writer) Printer {
	switch strings.ToLower(printerName) {
	case "json":
		return &jsonPrinter{writer: writer}
	case "xml":
		return &xmlPrinter{writer: writer}
	}
	return &defaultPrinter{writer: writer}
}

type xmlPrinter struct {
	writer io.Writer
}

func (xp *xmlPrinter) PrintHeader() {
	fmt.Fprintf(xp.writer, "<scv-results>")
}

func (xp *xmlPrinter) PrintEach(algorithmName string, v1, v2 *vector.Vector, result *result, first bool) {
	fmt.Fprint(xp.writer, "<scv-result>")
	fmt.Fprintf(xp.writer, `<algorithm>%s</algorithm><vectors><vector1>%s</vector1><vector2>%s</vector2></vectors><result>%s</result>`, algorithmName, v1.Value(), v2.Value(), result)
	fmt.Fprint(xp.writer, "</scv-result>")
}
func (xp *xmlPrinter) PrintFooter() {
	fmt.Fprintf(xp.writer, "</scv-results>")
}

type jsonPrinter struct {
	writer io.Writer
}

func (jp *jsonPrinter) PrintHeader() {
	fmt.Fprintf(jp.writer, "[")
}

func (jp *jsonPrinter) PrintEach(algorithmName string, v1, v2 *vector.Vector, result *result, first bool) {
	if !first {
		fmt.Fprintf(jp.writer, ",")
	}
	fmt.Fprintf(jp.writer, `{"algorithm:":"%s","vector1":"%s","vector2":"%s","result":"%s"}`, algorithmName, v1.Value(), v2.Value(), result)
}
func (jp *jsonPrinter) PrintFooter() {
	fmt.Fprintf(jp.writer, "]")
}

type defaultPrinter struct {
	writer io.Writer
}

func (dp *defaultPrinter) PrintHeader() {
}

func (dp *defaultPrinter) PrintEach(algorithmName string, v1, v2 *vector.Vector, result *result, first bool) {
	fmt.Fprintf(dp.writer, "%s(%s, %s): %s\n", algorithmName, v1.Value(), v2.Value(), result)
}

func (dp *defaultPrinter) PrintFooter() {
}
