package encode

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Logjson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}

func TestEncodeJson(t *testing.T) {
	Logjson("Arif")
	Logjson(100)
	Logjson(true)
	Logjson([]string{"Muhammad", "Arif", "Budiman"})
}


