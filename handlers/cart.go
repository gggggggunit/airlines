package handlers

import (
	"fmt"
	"net/http"
)

func CartList(rw http.ResponseWriter, r *http.Request)  {

	sd:=SessionStart(rw,r)

	if len(sd.Cart)==0{
		http.Redirect(rw,r,"/search",302)
		return
	}


	err:=tpl.ExecuteTemplate(rw, "cart.html", sd) //отправляем файл штмл на браузер
	if err != nil {
		fmt.Printf("Cart execute template:%s \n", err)
	}
}
