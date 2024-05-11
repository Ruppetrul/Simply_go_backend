package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var db *sql.DB

func main() {
	fmt.Println("Hello World")
	dbConnection()

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

func dbConnection() {
	fmt.Println("Db connection start")
	cfg := mysql.Config{
		User:   "myuser",
		Passwd: "mypassword",
		Net:    "tcp",
		Addr:   "db:3306",
		DBName: "mydatabase",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database")
}
