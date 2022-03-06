package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Araba struct {
	Marka string
	// Model int
	Fiyat int
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/golang")
	checkErr(err)
	defer db.Close()
	//GROUP BY fonksiyonları...   AVG(fiyat)=ort,
	result, err := db.Query("SELECT marka,MIN(fiyat) FROM araba WHERE denemeid=1 and fiyat>150000 GROUP BY marka")
	checkErr(err)
	var arabalar []Araba
	var araba Araba
	for result.Next() {
		err = result.Scan(&araba.Marka, &araba.Fiyat)
		checkErr(err)
		arabalar = append(arabalar, araba)
	}
	for i := 0; i < len(arabalar); i++ {
		fmt.Println(arabalar[i])
	}
}
