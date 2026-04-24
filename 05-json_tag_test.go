package golangpzn

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       int    `json:"id"` // `json="xdxdxd"` harus nempel dengan =
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:    1,
		Name:  "Rinso",
		Price: 1000,
	}

	byte, _ := json.Marshal(product)

	fmt.Println(string(byte))
}

func TestJSONTagDecode(t *testing.T) {

	// Masih bisa terbaca oleh golang
	// `{"id":1,"name":"Rinso","price":1000,"image_url":""}`
	// `{"ID":1,"NAme":"Rinso","priCE":1000,"IMAge_url":""}`
	// `{"ID":1,"NAME":"Rinso","PRICE":1000,"IMAGE_URL":""}`

	// asal semua nya lengkap,
	// Jika image_url jadi ImageUrl, ini tidak terbaca karena tidak ada _

	jsonString := `{"id":1,"name":"Rinso","PRICE":1000,"image_url":""}`
	jsonBytes := []byte(jsonString)

	product := Product{}
	err := json.Unmarshal(jsonBytes, &product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
