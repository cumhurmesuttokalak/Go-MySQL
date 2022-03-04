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
	// LIMIT: Kısıtlama
	result, err := db.Query("SELECT* FROM deneme ORDER BY yas ASC LIMIT 1")
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
	// MIN MAX :biggest or smallest   SUM:plus all ages    AVG:ortalama  (float32)
	resultM := db.QueryRow("SELECT MIN(yas) FROM deneme")
	var min int
	resultM.Scan(&min)
	fmt.Println(min)

}
