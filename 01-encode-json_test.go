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

	// Byte harus di conversi ke string supaya bisa dibaca.
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {

	fmt.Println("Bani") // Kalau print string biasa tidak ada ""
	LogJson("Bani")     // "Bani", di Json string add ""
	LogJson(1)          // 1
	LogJson(true)       // true

	LogJson([]string{
		"Burhanudin",
		"D",
		"Rabbani",
	}) // ["Burhanudin", "D", "Rabbani"]
}
