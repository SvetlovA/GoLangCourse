package main

import (
	"testing"
)

var record = initializeRecord()

func BenchmarkConcatFmt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ConcatFmt(record)
	}
}

func BenchmarkConcatBuffer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ConcatBuffer(record)
	}
}

func BenchmarkConcatJoin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ConcatJoin(record)
	}
}

//BenchmarkConcatFmt-8           2000000	       685 ns/op	     112 B/op	       7 allocs/op
//BenchmarkConcatBuffer-8   	 5000000	       242 ns/op	     128 B/op	       2 allocs/op
//BenchmarkConcatJoin-8     	10000000	       136 ns/op	      32 B/op	       2 allocs/op
