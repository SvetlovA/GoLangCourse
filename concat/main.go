// concat project main.go
package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

type Record struct {
	A string
	B string
	C string
	D string
	E string
	F string
}

func initializeRecord() Record {
	var record Record
	record.A = "_a"
	record.B = "_b"
	record.C = "_c"
	record.D = "_d"
	record.E = "_e"
	record.F = "_f"

	return record
}

func ConcatFmt(record Record) string {
	return fmt.Sprintf("%v%v%v%v%v%v",
		record.A,
		record.B,
		record.C,
		record.D,
		record.E,
		record.F)
}

func ConcatBuffer(record Record) string {
	var buffer bytes.Buffer
	_, err := buffer.WriteString(record.A)
	_, err = buffer.WriteString(record.B)
	_, err = buffer.WriteString(record.C)
	_, err = buffer.WriteString(record.D)
	_, err = buffer.WriteString(record.E)
	_, err = buffer.WriteString(record.F)

	if err != nil {
		log.Fatal(err)
	}

	return buffer.String()
}

func ConcatJoin(record Record) string {
	return strings.Join([]string{record.A, record.B, record.C, record.D, record.E, record.F}, "")
}

func main() {
	record := initializeRecord()

	fmt.Println(ConcatFmt(record))
	fmt.Println(ConcatBuffer(record))
	fmt.Println(ConcatJoin(record))
}
