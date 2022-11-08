package json_tag

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	ImageURL string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {

	product := Product{
		Name:     "Kartu Perdana XL",
		Quantity: 10,
		ImageURL: "http://example.co.id/image.png",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {

	jsonString := `{"name":"Kartu Perdana XL","quantity":10,"image_url":"http://example.co.id/image.png"}`

	product := &Product{}

	json.Unmarshal([]byte(jsonString), product)

	fmt.Println(product)
}
