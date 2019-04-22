package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func List(rw http.ResponseWriter, r *http.Request) { //ДЛЯ ОТОБРАЖЕНИЯ ТАБЛИЦЫ РЕЙСОВ

	var page int = 1
	var count int = 20
	var err error

	PageStr := r.FormValue("p")
	if PageStr != "" {
		page, err = strconv.Atoi(PageStr)
		if err != nil {
			page = 1
		}
	}
	fmt.Printf("page: %v; count: %v\n", page, count)

	var offset int = 0
	offset = page*count - count

	rows, err := dbc.Query("select count(*) from `airlines`")
	rows.Next()
	var total int
	rows.Scan(&total)
	fmt.Printf("total: %v\n", total)
	rows.Close()

	var pages = total / count

	query := fmt.Sprintf("select * from `airlines` limit %d,%d", offset, count)

	if r.RequestURI == "/favicon.ico" {
		return
	}

	//sd.Cart=append(sd.Cart,&Avia{})    //показ что происходит
	//fmt.Printf("%+v\n",sd)

	list := &AviaList{}
	list.List = make([]*Avia, 0, 2000)

	rows, err = dbc.Query(query) //"select * from `airlines` ") //запрос в базу на получение данных
	if err != nil {
		panic(err)
	}

	for rows.Next() { // до тех пор пока в базе есть строки
		avia := &Avia{}
		err := rows.Scan(&avia.ID, &avia.Source, &avia.Destination, &avia.When, &avia.Price) //загружаем с базы в переменые
		if err != nil {                                                                      // проверяем на ошибку
			panic(err)
		}
		list.List = append(list.List, avia) //добавляем в масив

	}
	rows.Close()

	for i := 1; i <= pages; i++ {
		list.Pages = append(list.Pages, i)
	}

	err = tpl.ExecuteTemplate(rw, "list.html", list) //отправляем файл штмл на браузер
	if err != nil {
		fmt.Printf("List execute template:%s \n", err)
	}
}
