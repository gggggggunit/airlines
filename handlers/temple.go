package handlers

import "html/template"

var tpl *template.Template

func init() {

	var err error //где находится штмл
	tpl, err = template.ParseFiles(
		"templates/list.html",
		"templates/search.html",
		"templates/buy.html",
		"templates/congrat.html",
		"templates/addtocart.html",
		"templates/cart.html",
		"templates/registration.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/nav.html",
		"templates/login.html",
		"templates/adminlist.html",
		"templates/adminadd.html",
	)

	if err != nil {
		panic(err)
	}
}
