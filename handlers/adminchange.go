package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func AdminChange(rw http.ResponseWriter, r *http.Request) {

	sd := SessionStart(rw, r)
	if sd.ID == 0 {
		fmt.Printf("некий с IP: %s хотел зайти\n", sd.Ip)
		http.Redirect(rw, r, "/login", 302)
		return
	}
	//-----------------------------------------------------------------------------------
	var err error
	var id int
	list := &AviaList{} //создаем указатель на авиялист(пустая структура) чтоб потом передать в темплейс
	//-----------------------------------------------------------------------------------

	if r.Method == "GET" {

		idString := r.FormValue("id") //получаем id в строке
		fmt.Printf("id from change: %s\n", idString)

		id, err = strconv.Atoi(idString)
		if err != nil {
			fmt.Printf("id parse error:%s; ID:%s\n", err, idString)
			http.Redirect(rw, r, "adminlist.html", 302) //redirect - перенаправление
			return
		}
		//-----------------------------------------------------------------------------------

		query := "select * from `airlines` where `id`=?"
		rows, err := dbc.Query(query, id)
		if err != nil {
			fmt.Printf("Change in airlines error: %s\n", err)
			http.Redirect(rw, r, "adminlist.html", 302)
			return
		}
		//-----------------------------------------------------------------------------------
		if rows.Next() == false {
			err = tpl.ExecuteTemplate(rw, "adminlist.html", nil)
			if err != nil {
				fmt.Printf("Request on Change error: %s\n", err)
				return
			}
		}
		//------------------------------------------------------------------------------------
		err = rows.Scan(&list.ID, &list.Source, &list.Destination, &list.When, &list.Price)
		if err != nil {
			panic(err)
		}
		//-----------------------------------------------------------------------------------
		err = tpl.ExecuteTemplate(rw, "adminchange.html", list)
		if err != nil {
			fmt.Printf("AdminChange ExecuteTemplate error: %s\n", err)
		}
		return
		//-----------------------------------------------------------------------------------
	} else if r.Method == "POST" {
		//-----------------------------------------------------------------------------------
		source := strings.TrimSpace(r.PostFormValue("source")) //(ДЛЯ изменения source)если метот post пишем r.PostFormValue,если get то r.FormValue
		if source == "" {
			err = tpl.ExecuteTemplate(rw, "adminchange.html", nil)
			if err != nil {
				fmt.Printf("Change source error: %s\n", err)
				return
			}
			return
		}
		//-----------------------------------------------------------------------------------
		destination := strings.TrimSpace(r.PostFormValue("destination")) //(ДЛЯ изменения destination)если метот post пишем r.PostFormValue,если get то r.FormValue
		if destination == "" {
			err = tpl.ExecuteTemplate(rw, "adminchange.html", nil)
			if err != nil {
				fmt.Printf("Change destination error: %s\n", err)
				return
			}
			return
		}
		//-----------------------------------------------------------------------------------
		when := strings.TrimSpace(r.PostFormValue("date111")) //(ДЛЯ изменения When)если метот post пишем r.PostFormValue,если get то r.FormValue
		if when == "" {
			err = tpl.ExecuteTemplate(rw, "adminchange.html", nil)
			if err != nil {
				fmt.Printf("Change when error: %s\n", err)
				return
			}
			return
		}
		//-----------------------------------------------------------------------------------
		price := strings.TrimSpace(r.PostFormValue("price")) //(ДЛЯ изменения price)если метот post пишем r.PostFormValue,если get то r.FormValue
		if price == "" {
			err = tpl.ExecuteTemplate(rw, "adminchange.html", nil)
			if err != nil {
				fmt.Printf("Change price error: %s\n", err)
				return
			}
			return
		}
		//fmt.Printf("source: %s; destination: %s; when: %s; price: %s\n ", source, destination, when, price)
		//-----------------------------------------------------------------------------------
		exec := "update `airlines` set `Source` = ? , `Destination` = ?, `When` = ?, `Price` = ? where (`id` = ?)"
		_, err = dbc.Exec(exec, source, destination, when, price, id)
		if err != nil {
			fmt.Printf("UPDATE in airlines error: %s\n", err)
			http.Redirect(rw, r, "adminadd.html", 302)
			return
		}
		//-----------------------------------------------------------------------------------
		http.Redirect(rw, r, "/admin/list", 302)
		fmt.Printf("After source: %s; destination: %s; when: %s; price: %s\n ", source, destination, when, price)
		fmt.Printf("UPDATE in airlines OK\n")

	}
}
