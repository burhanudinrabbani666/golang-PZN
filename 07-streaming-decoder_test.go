package golangpzn

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJsonStreamingDecoder(t *testing.T) {

	reader, _ := os.Open("./data/customer.json")
	decoder := json.NewDecoder(reader)

	customer := []Customer{}
	decoder.Decode(&customer)

	fmt.Println(customer)
}
