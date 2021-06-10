package main

import (
	"fmt"
	"io"
	"strings"
)

type Printer interface {
	PrintHeader()
	PrintEach(r *result, first bool)
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

func (xp *xmlPrinter) PrintEach(r *result, first bool) {
	fmt.Fprint(xp.writer, "<scv-result>")
	fmt.Fprintf(xp.writer, `<algorithm>%s</algorithm><vectors><vector1>%s</vector1><vector2>%s</vector2></vectors><result>%f</result>`, r.algorithm, r.pair.Vector1.Value(), r.pair.Vector2.Value(), r.similarity)
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

func (jp *jsonPrinter) PrintEach(r *result, first bool) {
	if !first {
		fmt.Fprintf(jp.writer, ",")
	}
	fmt.Fprintf(jp.writer, `{"algorithm:":"%s","vector1":"%s","vector2":"%s","result":%f}`, r.algorithm, r.pair.Vector1.Value(), r.pair.Vector2.Value(), r.similarity)
}
func (jp *jsonPrinter) PrintFooter() {
	fmt.Fprintf(jp.writer, "]")
}

type defaultPrinter struct {
	writer io.Writer
}

func (dp *defaultPrinter) PrintHeader() {
}

func (dp *defaultPrinter) PrintEach(r *result, first bool) {
	fmt.Fprintf(dp.writer, "%s(%s, %s): %f\n", r.algorithm, r.pair.Vector1.Value(), r.pair.Vector2.Value(), r.similarity)
}

func (dp *defaultPrinter) PrintFooter() {
}
