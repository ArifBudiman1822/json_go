package decode

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

func TestUnmarshal(t *testing.T) {

	data := `{"Name":"Muhammad Arif Budiman","Email":"arif@gmail.com","Age":21,"Married":false}`

	customer := &Customer{}

	err := json.Unmarshal([]byte(data), customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Email)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
