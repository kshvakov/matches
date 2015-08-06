package matches

import (
	"bufio"
	"bytes"
	"encoding/gob"
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

func BenchmarkGobEncoder(b *testing.B) {

	b.ReportAllocs()

	buff := &bytes.Buffer{}

	for i := 0; i < b.N; i++ {

		err := gob.NewEncoder(buff).Encode(&TestStruct)

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

	buff := bytes.NewBuffer(data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		result := CommonStruct{}

		err := json.NewDecoder(bufio.NewReader(buff)).Decode(&result)

		if err != nil {

			b.Error(err)
		}

		buff.Write(data)
	}
}

func BenchmarkMsgpackDecoder(b *testing.B) {

	b.ReportAllocs()

	data, err := msgpack.Marshal(&TestStruct)

	if err != nil {

		b.Error(err)
	}

	buff := bytes.NewBuffer(data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		result := CommonStruct{}

		err := msgpack.NewDecoder(bufio.NewReader(buff)).Decode(&result)

		if err != nil {

			b.Error(err)
		}

		buff.Write(data)
	}
}

func BenchmarkGobDecoder(b *testing.B) {

	b.ReportAllocs()

	dataBuff := &bytes.Buffer{}

	err := gob.NewEncoder(dataBuff).Encode(&TestStruct)

	if err != nil {

		b.Error(err)
	}

	data := dataBuff.Bytes()
	buff := bytes.NewBuffer(data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		result := CommonStruct{}

		err := gob.NewDecoder(bufio.NewReader(buff)).Decode(&result)

		if err != nil {

			b.Error(err)
		}

		buff.Write(data)
	}
}
