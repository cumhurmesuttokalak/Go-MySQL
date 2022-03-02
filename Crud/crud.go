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
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/golang")
	checkErr(err)
	defer db.Close()
	result, err := db.Query("SELECT* FROM deneme WHERE yas <51 AND yas >22 ")
	checkErr(err)
	for result.Next() {
		var denemeObje Deneme
		result.Scan(&denemeObje.Id, &denemeObje.Isim, &denemeObje.Yas)
		fmt.Println(denemeObje)
	}
	_, err = db.Query("DELETE FROM deneme WHERE isim='Halim' ")
	checkErr(err)

	_, err = db.Query("UPDATE deneme SET isim='Veli',yas=65 WHERE isim='Mehmet' ")
	checkErr(err)

}
