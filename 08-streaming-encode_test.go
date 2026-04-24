package golangpzn

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJSONStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Burhanudin",
		MiddleName: "D",
		LastName:   "",
	}

	encoder.Encode(customer)
	fmt.Println(encoder)
}
