package handlers

import (
	"fmt"
	"net/http"
)

func Congrat(rw http.ResponseWriter, r *http.Request) {

	sd:=SessionStart(rw,r)
	sd.Cart=make([]*Avia,0)

	err:=tpl.ExecuteTemplate(rw, "congrat.html", nil) //отправляем файл штмл на браузер
	if err != nil {
		fmt.Printf("Congrat execute template:%s \n", err)
	}
}