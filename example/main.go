package main

import (
	"encoding/json"
	"fmt"
	jflat "go-json-flatten"
)

type A struct {
	Field1 int64 `json:"field1"`
}

type B struct {
	Field2 int64 `json:"field2"`
}

func main() {
	a := &A{
		Field1: 123,
	}

	b := &B{
		Field2: 321,
	}

	flat := jflat.MustFlatten(a, b)

	data, err := json.Marshal(flat)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	err = json.Unmarshal([]byte("{\"field1\":222,\"field2\":555}"), flat)
	if err != nil {
		panic(err)
	}

	fmt.Println(a, b)
}
