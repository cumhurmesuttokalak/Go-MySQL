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
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	//DATABASE CONNECTİON
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/golang")
	checkErr(err)
	defer db.Close()
	//ORDER BY Sıralama Yapar || ASC küçükten büyüğe , DESC büyükten küçüğe sıralama yapar
	result, err := db.Query("SELECT* FROM deneme ORDER BY yas ASC")
	checkErr(err)
	var elemanlar []Deneme
	var eleman Deneme
	for result.Next() {
		result.Scan(&eleman.Id, &eleman.Isim, &eleman.Yas)
		elemanlar = append(elemanlar, eleman)
	}
	for i := 0; i < len(elemanlar); i++ {
		fmt.Println(elemanlar[i])
	}

}
