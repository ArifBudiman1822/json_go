package stream

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Customer struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
	Email     string `json: "email"`
}

func TestStreamDecoder(t *testing.T) {

	reader, _ := os.Open("Customer.json")

	decoder := json.NewDecoder(reader)

	customer := &Customer{}

	decoder.Decode(customer)

	fmt.Println(customer)

}

func TestStreamEncoder(t *testing.T) {

	writer, _ := os.Create("NewCustomer.json")

	encoder := json.NewEncoder(writer)

	var customer = Customer{
		Firstname: "Muhammad",
		Lastname:  "Arif",
		Email:     "arif02@gmail.com",
	}

	encoder.Encode(customer)

}

func TestAja(t *testing.T) {

	reader, _ := os.Open("NewCustomer.json")

	decoder := json.NewDecoder(reader)

	newcustomer := &Customer{}

	decoder.Decode(newcustomer)

	fmt.Println(newcustomer)
}
