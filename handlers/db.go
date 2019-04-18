package handlers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dbc *sql.DB				//ПОДКЛЮЧЕНИЕ К БД

func init() {
	var err error
	dbc, err = sql.Open("mysql", "airline:qqq@tcp(localhost:3306)/test")
	if err != nil {					    // подключуние к ДБ(имя)(пароль)tcp(адрес:порт)путь
		panic(err)
	}

	err = dbc.Ping()
	if err != nil {
		panic(err)
	}

	//defer close()

}
