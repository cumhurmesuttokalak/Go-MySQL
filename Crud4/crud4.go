package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Deneme struct {
	Id   int
	Isim string
	Yas  int
	Time string
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
	//LIKE : searcing WHERE isim LIKE 'M%','%t',''%a%,'____%a'
	result, err := db.Query("SELECT* FROM deneme WHERE isim LIKE '___%a%' ")
	checkErr(err)
	var elemanlar []Deneme
	var eleman Deneme
	for result.Next() {
		result.Scan(&eleman.Id, &eleman.Isim, &eleman.Yas, &eleman.Time)
		elemanlar = append(elemanlar, eleman)
	}
	for i := 0; i < len(elemanlar); i++ {
		fmt.Println(elemanlar[i])
	}
}
