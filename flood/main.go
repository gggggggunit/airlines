package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

var cities = []string{"Marrakesh", "Amsterdam", "Baghdad", "Bangkok", "Barcelona",
	"Beijing", "Belgrade", "Berlin", "Bogota", "Bratislava", "Brussels", "Bucharest", "Budapest",
	"Buenos Aires", "Cairo", "Cape Town", "Caracas", "Chicago", "Copenhagen", "Dhaka",
	"Dubai", "Dublin", "Frankfurt", "Geneva", "Marrakesh", "Manila", "Mexico City",
	"Montreal", "Moscow", "Mumbai", "Nairobi", "New Delhi", "New York", "Nicosia",
	"Oslo", "Ottawa", "Paris", "Prague", "Reykjavik", "Riga", "Rio de Janeiro", "Rome", "Saint Petersburg",
	"San Francisco", "Santiago ", "São Paulo", "Seoul", "Shanghai", "Singapore", "Sofia", "Stockholm",
	"Sydney", "Tallinn", "Tallinn", "Tehran"}

const n = 500

func main() {
	rand.Seed(time.Now().UnixNano())


	dbc, err := sql.Open("mysql", "airline:qqq@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	err = dbc.Ping()
	if err != nil {
		panic(err)
	}
	defer dbc.Close()

	for i := 0; i < n; i++ { // получаем два разных города
		source := cities[rand.Intn(len(cities))]
		dest := ""
		for {
			dest = cities[rand.Intn(len(cities))]
			if source != dest {
				break
			}
		}

		when := time.Now()
		rnd := rand.Intn(31)
		when = when.Add(time.Hour * 24 * time.Duration(rnd))

		prc:=rand.Intn(200)+60

		_, err = dbc.Exec("insert into `airlines`(`Source`,`Destination`,`When`,`Price`) values (?,?,?,?)", source, dest, when.Format("2006-01-02"),prc)

		if err != nil {
			panic(err)
		}

	}

}
