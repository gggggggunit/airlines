package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func Buy(rw http.ResponseWriter, r *http.Request) { //ДЛЯ покупки РЕЙСОВ

	SessionStart(rw, r)

	avia := &Avia{}
	var err error
	var rows *sql.Rows

	idn := r.FormValue("id") //получаем id
	fmt.Printf("id:%s\n", idn)

	id, err := strconv.Atoi(idn)
	if err != nil {
		fmt.Printf("id parse error:%s; str:%s\n", err, idn)
		http.Redirect(rw, r, "/search", 302) //redirect - перенаправление
		return
	}

	rows, err = dbc.Query("select * from `airlines` where `id`=? ", idn) //запрос в базу на получение данных
	if err != nil {
		panic(err)
	}

	if rows.Next() == false { //проверка на наличие строк
		fmt.Printf("no airline whith ID:%d\n", id)
		http.Redirect(rw, r, "/search", 302)
		return
	}
	err = rows.Scan(&avia.ID, &avia.Source, &avia.Destination, &avia.When, &avia.Price) //сканируем полученый с базы запрос
	if err != nil {
		fmt.Printf("airline scan error:%s\n", err)
		http.Redirect(rw, r, "/search", 302)
		return
	}
	fmt.Printf("airline:%+v\n", avia)

	//for rows.Next() { // до тех пор пока в базе есть строки
	//	err := rows.Scan(&avia.ID, &avia.Source, &avia.Destination, &avia.When ,&avia.Price) //загружаем с базы в переменые
	//	if err != nil {                                                         // проверяем на ошибку
	//		panic(err)
	//	}
	//}
	rows.Close()

	err = tpl.ExecuteTemplate(rw, "buy.html", avia) //отправляем файл штмл на браузер

	if err != nil {
		fmt.Printf("------ Buy template error:%s \n", err)
		http.Redirect(rw, r, "/search", 302)
	}
}
