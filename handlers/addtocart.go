package handlers

import (
	"fmt"
	"net/http"
)

var cartStorage = &CartStorage{}

func init() { //создает в CartStorage место
	cartStorage.Data = make(map[string]interface{}, 10)
}

func Addtocart(rw http.ResponseWriter, r *http.Request) {

	sd := SessionStart(rw, r) //СЕСИЯ

	id := r.FormValue("id")
	// var price int
	var a Avia
	//cartStorage.Set("user",id)
	//http.Redirect(rw,r,"/search",302)

	//buyDate := time.Now().Format("2006-01-02 15:04:05")

	rows, err := dbc.Query("select * from `airlines` where `id`=? ", id) //запрос в базу на получение данных
	if err != nil {
		panic(err)
	}

	if rows.Next() == false { //проверка на наличие строк
		fmt.Printf("no airline whith ID:%d\n", id)
		http.Redirect(rw, r, "/search", 302)
		return
	}

	errr := rows.Scan(&a.ID, &a.Source, &a.Destination, &a.When, &a.Price) //сканируем полученый с базы запрос

	if errr != nil {
		fmt.Printf("airline scan error:%s\n", errr)
		http.Redirect(rw, r, "/search", 302)
		return
	}

	sd.Cart = append(sd.Cart, &a)
	sd.Sum+=a.Price					//симируем цену билетов
	sd.Print()
	//TODO

	//idn, _ := strconv.Atoi(id)
	//query := "insert into `cart`(`ID`,`AirlineID`,`Price`,`BuyDate`) values (NULL,?,?,?)"
	//_, erro := dbc.Exec(query, idn, price, buyDate)
	//if erro != nil {
	//	fmt.Printf("insert cart error:%s\n", erro)
	//}
	//fmt.Printf("insert cart!!!")

	http.Redirect(rw, r, "/search", 302)
}
