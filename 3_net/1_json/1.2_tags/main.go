package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type mySctruct struct {
	// Будет использовать название и тип поля
	Field1 int
	// Изменено название
	Field2 int `json:"renamedField2"`
	// Вырезается zero-value и изменено название
	Field3 int `json:"renamedField3,omitempty"`
	// Вырезается zero-value
	Field4 int `json:",omitempty"`
	// Поле игнорируется
	Field5 int `json:"-"`
	// Поле будет называться `-`
	Field6 int `json:"-,"`
	// Примет строковое представление (слишком большое число для JS)
	Int64String int64 `json:",string"`
}

func main() {
	fmt.Println()

	var jsonBytes []byte

	{
		in := mySctruct{
			Field1:      10,
			Field2:      20,
			Field3:      30,
			Field4:      40,
			Field5:      50,
			Field6:      60,
			Int64String: math.MaxInt64,
		}
		var err error
		jsonBytes, err = json.MarshalIndent(in, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	}
	fmt.Println()

	{
		out := mySctruct{}
		err := json.Unmarshal(jsonBytes, &out)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", out)
	}
}
