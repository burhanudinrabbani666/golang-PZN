package golangpzn

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapEncode(t *testing.T) {
	jsonString := `{"id": "AP-001","name":"Apple Macbook Pro","PRICE": 1234,"image_url":"https://example.com"}`
	jsonByte := []byte(jsonString)

	var result map[string]any
	json.Unmarshal(jsonByte, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])

}

func TestJSONMapDecode(t *testing.T) {
	product := map[string]any{
		"id":    "P-001",
		"name":  "Apple Mac Book",
		"price": 1234,
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
