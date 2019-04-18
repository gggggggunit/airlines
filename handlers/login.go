package handlers

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
)

func Login(rw http.ResponseWriter, r *http.Request) {

	sd := SessionStart(rw, r) //СЕСИЯ

	if r.Method == "GET" {

		err := tpl.ExecuteTemplate(rw, "login.html", nil)
		if err != nil {
			fmt.Printf("Login get error: %s\n", err)
		}
	} else if r.Method == "POST" {

		email := strings.TrimSpace(r.PostFormValue("email")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if email == "" {
			err := tpl.ExecuteTemplate(rw, "login.html", nil)
			if err != nil {
				fmt.Printf("Email login get error: %s\n", err)
				return
			}
			return
		}
		fmt.Printf("email login:%s\n", email)

		password1 := strings.TrimSpace(r.PostFormValue("password1")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if password1 == "" {
			err := tpl.ExecuteTemplate(rw, "login.html", nil)
			if err != nil {
				fmt.Printf("Login password1 get error: %s\n", err)
				return
			}
			return
		}
		fmt.Printf("Login password1:%s\n", password1)

		password1 = fmt.Sprintf("%x", sha1.Sum([]byte(password1))) //хешируем пароль

		//err := tpl.ExecuteTemplate(rw, "login.html", nil) //отправляем файл штмл на браузер
		//if err != nil {
		//	fmt.Printf("Login execute template:%s \n", err)
		//}
		rows, err := dbc.Query("select `ID`,`Ballance` from `users` where `Email`=? and `Password`=? ", email, password1) //запрос в базу на получение данных
		if err != nil {
			panic(err)
		}

		if rows.Next() == false {
			err := tpl.ExecuteTemplate(rw, "login.html", nil)
			if err != nil {
				fmt.Printf("Login not OK get error: %s\n", err)
				return
			}

		}

		var (
			iid       int
			bballance int
		)

		err = rows.Scan(&iid, &bballance)
		if err != nil {
			panic(err)
		}

		sd.ID=iid
		sd.Ballance=bballance
		sd.Email=email
		sd.Print()
		http.Redirect(rw, r, "/search", 302)


	}

}
