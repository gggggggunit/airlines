package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

func Search(rw http.ResponseWriter, r *http.Request) { //ДЛЯ ПОИСКА РЕЙСА

	SessionStart(rw, r)

	list := &AviaList{} //создаем указатель на авиялист(пустая структура) чтоб потом передать в темплейс

	if r.Method == "POST" { //если  метод запроса post
		fmt.Println("post:/search")

		source := strings.TrimSpace(r.PostFormValue("source")) //(ДЛЯ SOURCE)если метот post пишем r.PostFormValue,если get то r.FormValue
		fmt.Printf("source:%s\n", source)
		list.Source = source //сохраняем из формы в структуру,что бы потом отобразить в форме

		destination := strings.TrimSpace(r.PostFormValue("destination")) //ДЛЯ DESTINATION (strings.TrimSpace!!!!!!-убирает пробелы)
		fmt.Printf("destination:%s\n", destination)
		list.Destination = destination //сохраняем из формы в структуру,что бы потом отобразить в форме

		date := strings.TrimSpace(r.PostFormValue("date111")) //ДЛЯ DESTINATION (strings.TrimSpace!!!!!!-убирает пробелы)
		fmt.Printf("date:%s\n", date)
		list.When = date //сохраняем из формы в структуру,что бы потом отобразить в форме

		pricee := strings.TrimSpace(r.PostFormValue("price22")) //ДЛЯ DESTINATION (strings.TrimSpace!!!!!!-убирает пробелы)
		fmt.Printf("Price:%s\n", pricee)
		list.Price = pricee //сохраняем из формы в структуру,что бы потом отобразить в форме

		if source == "" && destination == "" && date == "" && pricee == "" {
			http.Redirect(rw, r, "/search", 302) //redirect перенаправляет с одной страницы на другую
			return
		}

		var err error
		var rows *sql.Rows

		where := ""
		WhereData := make([]interface{}, 0, 3)

		if source != "" {
			where += "`source`=?"
			WhereData = append(WhereData, source)
		}
		if destination != "" {
			if where != "" {
				where += " AND"
			}
			where += " `destination`=?"
			WhereData = append(WhereData, destination)
		}
		if date != "" {
			if where != "" {
				where += " AND"
			}
			where += " `When`=?"
			WhereData = append(WhereData, date)
		}
		if pricee != "" {
			if where != "" {
				where += " AND"
			}
			where += " `Price`=?"
			WhereData = append(WhereData, pricee)
		}
		query := "select * from `airlines` where " + where
		fmt.Println(WhereData)
		rows, err = dbc.Query(query, WhereData...)

		//
		//if source !="" && destination!=""{ //если два не пустые написан запрос
		//
		//	rows, err = dbc.Query("select * from `airlines` where `source` =? and `destination`=?", source, destination)
		//}else if source!=""{		//если только в сонрс написан запрос
		//
		//	rows, err = dbc.Query("select * from `airlines` where `source` =?", source)
		//} else if destination !=""{		//если только в дестенейшей написан запрос
		//
		//	rows, err = dbc.Query("select * from `airlines` where `destination` =?", destination)
		//}

		//rows, err := dbc.Query("select * from `airlines` where `source` =? and `destination`=?", source, destination)
		if err != nil {
			fmt.Printf("select airlines error:%s\n", err)
			fmt.Fprintf(rw, "error")
			return
		}
		for rows.Next() { //до тех пор пока есть строки в базе
			a := &Avia{}                                                         //указатель на структуру которая описует ее
			err = rows.Scan(&a.ID, &a.Source, &a.Destination, &a.When, &a.Price) //берем из rows(которую получили)
			if err != nil {
				panic(err)
			}
			list.List = append(list.List, a) //добавляем полученую строку в лист
		}
	}

	err := tpl.ExecuteTemplate(rw, "search.html", list) //отправляем файл штмл на браузер
	if err != nil {
		fmt.Printf("search execute template:%s \n", err)
	}
}
