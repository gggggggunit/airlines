package main

import (
	"airLines/handlers"
	"math/rand"
	"time"

	"net/http"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	//handlers.CreateHash()
	//dbc, err := sql.Open("mysql", "airline:qqq@tcp(localhost:3306)/test")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = dbc.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//defer dbc.Close()
	//
	//fmt.Printf("%+v\n", dbc.Stats())
	//query := "insert into `airlines`(`Source`,`Destination`,`When`)" +
	//	"values('Kiev','Minsk','2019-03-27 17:30:01')"
	//result, err := dbc.Exec(query) //для проверки
	//
	//if err != nil {
	//	panic(err)
	//}
	//num, err := result.RowsAffected()
	//id, err := result.LastInsertId()
	// changes
	//fmt.Printf("num:%d; id:%d\n", num, id)

	http.HandleFunc("/", handlers.List) // handlers пакет с файлами GO

	http.HandleFunc("/search", handlers.Search)

	http.HandleFunc("/buy", handlers.Buy)

	http.HandleFunc("/congrat", handlers.Congrat)

	http.HandleFunc("/addtocart", handlers.Addtocart)

	http.HandleFunc("/cart", handlers.CartList)

	http.HandleFunc("/registration", handlers.Registration)

	http.HandleFunc("/login", handlers.Login)

	http.HandleFunc("/admin/list", handlers.AdminList)

	http.HandleFunc("/admin/add", handlers.AdminAdd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
