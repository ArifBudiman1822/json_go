package encode

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Name    string
	Email   string
	Age     int
	Married bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		Name:    "Muhammad Arif Budiman",
		Email:   "arif@gmail.com",
		Age:     21,
		Married: false,
	}

	bytes, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}
