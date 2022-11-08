package json_array

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Name      string
	Email     string
	Age       int
	Married   bool
	Hobbies   []string
	Addresses []Address
}

type Address struct {
	Street     string
	City       string
	Postalcode int
}

func TestJsonArray(t *testing.T) {

	customer := Customer{
		Name:    "Muhammad Arif Budiman",
		Hobbies: []string{"Swimming", "Running", "Coding"},
	}

	bytes, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}

func TestJsonArrayDecode(t *testing.T) {

	jsonString := `{"Name":"Muhammad Arif Budiman","Email":"","Age":0,"Married":false,"Hobbies":["Swimming","Running","Coding"]}`

	customer := &Customer{}

	err := json.Unmarshal([]byte(jsonString), customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Hobbies)
}

func TestJsonObject(t *testing.T) {

	customer := Customer{
		Name: "Muhammad Arif",
		Addresses: []Address{
			{
				Street:     "Jalan Singokarso",
				City:       "Sidoarjo",
				Postalcode: 9999,
			},
			{
				Street:     "Jalan Ngampelsari",
				City:       "Surabaya",
				Postalcode: 44444,
			},
		},
	}

	bytes, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonObjectDecode(t *testing.T) {

	jsonString := `[{"Street":"Jalan Singokarso","City":"Sidoarjo","Postalcode":9999},{"Street":"Jalan Ngampelsari","City":"Surabaya","Postalcode":44444}]`

	addresses := &[]Address{}

	err := json.Unmarshal([]byte(jsonString), addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
