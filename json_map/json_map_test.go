package json_map

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMap(t *testing.T) {
	var product = map[string]interface{}{
		"id":    "P0001",
		"name":  "Beras Raja Lele",
		"price": 250000,
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestJsonMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Beras Raja Lele","price":250000}`

	var result map[string]interface{}

	json.Unmarshal([]byte(jsonString), &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}
