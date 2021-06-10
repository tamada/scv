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
	fmt.Fprintf(xp.writer, `<algorithm>%s</algorithm><vectors><vector1>%s</vector1><vector2>%s</vector2></vectors>`, r.algorithm, r.pair.Vector1.Value(), r.pair.Vector2.Value())
	if r.err != nil {
		fmt.Fprintf(xp.writer, `<error>%s</error>`, r.err.Error())
	} else {
		fmt.Fprintf(xp.writer, `<result>%f</result>`, r.similarity)
	}
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
	fmt.Fprintf(jp.writer, `{"algorithm:":"%s","vector1":"%s","vector2":"%s"`, r.algorithm, r.pair.Vector1.Value(), r.pair.Vector2.Value())
	if r.err != nil {
		fmt.Fprintf(jp.writer, `,"error":"%s"}`, r.err.Error())
	} else {
		fmt.Fprintf(jp.writer, `,"result":%f}`, r.similarity)
	}
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
	fmt.Fprintf(dp.writer, "%s(%s, %s): ", r.algorithm, r.pair.Vector1.Value(), r.pair.Vector2.Value())
	if r.err != nil {
		fmt.Fprintln(dp.writer, r.err.Error())
	} else {
		fmt.Fprintf(dp.writer, "%f\n", r.similarity)
	}
}

func (dp *defaultPrinter) PrintFooter() {
}
