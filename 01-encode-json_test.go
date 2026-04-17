package golangpzn

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data any) {
	byte, errorJson := json.Marshal(data)

	if errorJson != nil {
		panic(errorJson)
	}

	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	LogJson("Bani")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Burhanudin", "D", "Rabbani"})
}
