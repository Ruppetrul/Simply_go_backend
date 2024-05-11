package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/product", productHandler)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}

func productHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Only GET supported.", http.StatusMethodNotAllowed)
		return
	}

	var testProduct = new(product)
	testProduct.Name = "Test product"
	testProduct.Price = 777.77

	jsonData, err := json.Marshal(testProduct)
	if err != nil {
		http.Error(w, "Json convert error", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, string(jsonData))
	if err != nil {
		fmt.Println("Response error:", err)
		return
	}
}
