package matches

import (
	"bufio"
	"bytes"
	"encoding/json"
	"gopkg.in/vmihailenco/msgpack.v2"
	"math"
	"testing"
)

type CommonStruct struct {
	Int              int
	Uint             uint
	String           string
	SliceString      []string
	SubStruct        SubStruct
	SubStructSlice   []SubStruct
	SubStructPointer *SubStruct
}

type SubStruct struct {
	SliceInt    []int
	SliceUint   []uint
	SliceString []string
}

var TestStruct = CommonStruct{
	Int:         math.MaxInt32,
	Uint:        math.MaxUint32,
	String:      "StringStringStringStringStringStringStringStringStringString",
	SliceString: []string{"String", "String", "String", "String", "String"},
	SubStruct: SubStruct{
		SliceInt:    []int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32},
		SliceUint:   []uint{math.MaxUint32, math.MaxUint32, math.MaxUint32, math.MaxUint32},
		SliceString: []string{"String", "String", "String", "String", "String"},
	},
	SubStructSlice: []SubStruct{
		{
			SliceInt:    []int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32},
			SliceUint:   []uint{math.MaxUint32, math.MaxUint32, math.MaxUint32, math.MaxUint32},
			SliceString: []string{"String", "String", "String", "String", "String"},
		},
		{
			SliceInt:    []int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32},
			SliceUint:   []uint{math.MaxUint32, math.MaxUint32, math.MaxUint32, math.MaxUint32},
			SliceString: []string{"String", "String", "String", "String", "String"},
		},
		{
			SliceInt:    []int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32},
			SliceUint:   []uint{math.MaxUint32, math.MaxUint32, math.MaxUint32, math.MaxUint32},
			SliceString: []string{"String", "String", "String", "String", "String"},
		},
	},
	SubStructPointer: &SubStruct{
		SliceInt:    []int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32},
		SliceUint:   []uint{math.MaxUint32, math.MaxUint32, math.MaxUint32, math.MaxUint32},
		SliceString: []string{"String", "String", "String", "String", "String"},
	},
}

func BenchmarkJsonMarshal(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_, err := json.Marshal(&TestStruct)

		if err != nil {

			b.Error(err)
		}
	}
}

func BenchmarkMsgpackMarshal(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_, err := msgpack.Marshal(&TestStruct)

		if err != nil {

			b.Error(err)
		}
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {

	b.ReportAllocs()

	data, err := json.Marshal(&TestStruct)

	if err != nil {

		b.Error(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		var result CommonStruct

		err := json.Unmarshal(data, &result)

		if err != nil {

			b.Error(err)
		}
	}
}

func BenchmarkMsgpackUnmarshal(b *testing.B) {

	b.ReportAllocs()

	data, err := msgpack.Marshal(&TestStruct)

	if err != nil {

		b.Error(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		var result CommonStruct

		err := msgpack.Unmarshal(data, &result)

		if err != nil {

			b.Error(err)
		}
	}
}

func BenchmarkJsonEncoder(b *testing.B) {

	b.ReportAllocs()

	buff := &bytes.Buffer{}

	for i := 0; i < b.N; i++ {

		err := json.NewEncoder(buff).Encode(&TestStruct)

		if err != nil {

			b.Error(err)
		}

		buff.Reset()
	}
}

func BenchmarkMsgpackEncoder(b *testing.B) {

	b.ReportAllocs()

	buff := &bytes.Buffer{}

	for i := 0; i < b.N; i++ {

		err := msgpack.NewEncoder(buff).Encode(&TestStruct)

		if err != nil {

			b.Error(err)
		}

		buff.Reset()
	}
}

func BenchmarkJsonDecoder(b *testing.B) {

	b.ReportAllocs()

	data, err := json.Marshal(&TestStruct)

	if err != nil {

		b.Error(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		var (
			result = CommonStruct{}
			buff   = bytes.NewBuffer(data)
		)

		err := json.NewDecoder(bufio.NewReader(buff)).Decode(&result)

		if err != nil {

			b.Error(err)
		}
	}
}

func BenchmarkMsgpackDecoder(b *testing.B) {

	b.ReportAllocs()

	data, err := msgpack.Marshal(&TestStruct)

	if err != nil {

		b.Error(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		var (
			result = CommonStruct{}
			buff   = bytes.NewBuffer(data)
		)

		err := msgpack.NewDecoder(bufio.NewReader(buff)).Decode(&result)

		if err != nil {

			b.Error(err)
		}
	}
}
