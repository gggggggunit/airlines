package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func AdminDelete(rw http.ResponseWriter, r *http.Request) {

	var err error
	var id int

	sd := SessionStart(rw, r)
	if sd.ID == 0 {
		fmt.Printf("некий с IP: %s хотел зайти\n", sd.Ip)
		http.Redirect(rw, r, "/login", 302)
		return
	}

	idString := r.FormValue("id") //получаем id
	fmt.Printf("id from delete: %s\n", idString)

	id, err = strconv.Atoi(idString)
	if err != nil {
		fmt.Printf("id parse error:%s; ID:%s\n", err, idString)
		http.Redirect(rw, r, "adminlist.html", 302) //redirect - перенаправление
		return
	}

	query := "delete from `airlines` where `id`=?"
	_, err = dbc.Exec(query, id)
	if err != nil {
		fmt.Printf("DELETE in airlines error: %s\n", err)
		http.Redirect(rw, r, "adminlist.html", 302)
		return
	}
	fmt.Printf("DELETE ID:%d in airlines OK\n", id)
	http.Redirect(rw, r, "/admin/list", 302)

}
