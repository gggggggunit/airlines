package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func AdminAdd(rw http.ResponseWriter, r *http.Request) {
	var err error
	sd := SessionStart(rw, r)
	if sd.ID == 0 {
		fmt.Printf("некий с IP: %s хотел зайти\n", sd.Ip)
		http.Redirect(rw, r, "/login", 302)
		return
	}

	if r.Method == "GET" {
		err = tpl.ExecuteTemplate(rw, "adminadd.html", nil)
		if err != nil {
			fmt.Printf("Adminadd ExecuteTemplate error: %s\n", err)
		}
		return
	} else if r.Method == "POST" {

		source := strings.TrimSpace(r.PostFormValue("source")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if source == "" {
			err := tpl.ExecuteTemplate(rw, "adminadd.html", nil)
			if err != nil {
				fmt.Printf("Add source error: %s\n", err)
				return
			}
			return
		}

		destination := strings.TrimSpace(r.PostFormValue("destination")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if destination == "" {
			err := tpl.ExecuteTemplate(rw, "adminadd.html", nil)
			if err != nil {
				fmt.Printf("Add destination error: %s\n", err)
				return
			}
			return
		}

		when := strings.TrimSpace(r.PostFormValue("date111")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if when == "" {
			err := tpl.ExecuteTemplate(rw, "adminadd.html", nil)
			if err != nil {
				fmt.Printf("Add when error: %s\n", err)
				return
			}
			return
		}

		price := strings.TrimSpace(r.PostFormValue("price")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if price == "" {
			err := tpl.ExecuteTemplate(rw, "adminadd.html", nil)
			if err != nil {
				fmt.Printf("Add price error: %s\n", err)
				return
			}
			return
		}
		fmt.Printf("source: %s; destination: %s; when: %s; price: %s\n ", source, destination, when, price)

		query := "insert into `airlines`(`Source`,`Destination`,`When`,`Price`) values (?,?,?,?)"
		_, err := dbc.Exec(query, source, destination, when, price)
		if err != nil {
			fmt.Printf("INSERT in airlines error: %s\n", err)
			http.Redirect(rw, r, "adminadd.html", 302)
			return
		}

		http.Redirect(rw, r, "/admin/add", 302)
		fmt.Printf("INSERT in airlines OK\n")

	}
}
