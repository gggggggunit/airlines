package handlers

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
)

func Registration(rw http.ResponseWriter, r *http.Request) {

	SessionStart(rw, r)

	if r.Method == "GET" {

		err := tpl.ExecuteTemplate(rw, "registration.html", nil)
		if err != nil {
			fmt.Printf("Registreation get error: %s\n", err)
		}

	} else if r.Method == "POST" {

		email := strings.TrimSpace(r.PostFormValue("email")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if email == "" {
			err := tpl.ExecuteTemplate(rw, "registration.html", nil)
			if err != nil {
				fmt.Printf("Registreation email get error: %s\n", err)
				return
			}
			return
		}
		fmt.Printf("email:%s\n", email)

		password1 := strings.TrimSpace(r.PostFormValue("password1")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if password1 == "" {
			err := tpl.ExecuteTemplate(rw, "registration.html", nil)
			if err != nil {
				fmt.Printf("Registreation password1 get error: %s\n", err)
				return
			}
			return
		}
		fmt.Printf("password1:%s\n", password1)

		password2 := strings.TrimSpace(r.PostFormValue("password2")) //(ДЛЯ регистрации емейла)если метот post пишем r.PostFormValue,если get то r.FormValue
		if password1 != password2 {
			err := tpl.ExecuteTemplate(rw, "registration.html", nil)
			if err != nil {
				fmt.Printf("Registreation password2 get error: %s\n", err)
				return
			}
			return
		}
		fmt.Printf("password2:%s\n", password2)


		password1=fmt.Sprintf("%x",sha1.Sum([]byte(password1)))	//хешируем пароль

		query := "insert into `users`(`ID`,`Email`,`Password`,`Ballance`) values (NULL,?,?,0)"
		_, err := dbc.Exec(query, email, password1,)
			if err!=nil{
				fmt.Printf("INSERT user error: %s\n",err )
				http.Redirect(rw,r,"/",302)
				return
			}

		http.Redirect(rw,r,"/login",302)
		fmt.Printf("registration OK: %s\n",email)

	}



}
