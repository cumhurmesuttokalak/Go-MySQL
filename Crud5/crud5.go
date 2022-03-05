package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Deneme struct {
	//Id    int
	Isim string
	//Yas   int
	//Time  string
	Araba struct {
		Marka string
		Model int
	}
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
	//INNER,LEFT,RIGHT and FULL OUTER JOIN ınner=tabloların kesişimini alır left:hem kesişim hem sol tablo verileri rıght=leftin tersi full outer:bütün verileri getirir.
	//OUTER JOIN yaparken 2 değişken belirleyim hem lefi hem rightı alıyoruz uniondan sonra boşluk bırakmayı unutmuyoruz
	q1 := "SELECT COALESCE(deneme.isim,'isimsiz'),COALESCE(araba.marka,'tanımsız'),COALESCE(araba.model,0) FROM deneme LEFT JOIN araba ON deneme.id=araba.denemeid"
	q2 := "SELECT COALESCE(deneme.isim,'isimsiz'),COALESCE(araba.marka,'tanımsız'),COALESCE(araba.model,0) FROM deneme RIGHT JOIN araba ON deneme.id=araba.denemeid"
	result, err := db.Query(q1 + " union " + q2)
	//result, err := db.Query("SELECT COALESCE(deneme.isim,'isimsiz'),COALESCE(araba.marka,'tanımsız'),COALESCE(araba.model,0) FROM deneme RIGHT JOIN araba ON deneme.id=araba.denemeid")
	checkErr(err)
	var elemanlar []Deneme
	var eleman Deneme
	for result.Next() {
		result.Scan(&eleman.Isim, &eleman.Araba.Marka, &eleman.Araba.Model)
		elemanlar = append(elemanlar, eleman)
	}
	for i := 0; i < len(elemanlar); i++ {
		fmt.Println(elemanlar[i])
	}
}
