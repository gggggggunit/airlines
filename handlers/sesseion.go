package handlers

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func SessionStart(rw http.ResponseWriter, r *http.Request) *SessionData { //проверка на авторизацию

	propusk, err := r.Cookie("propusk") //если куки?
	if err != nil {                     //если нет такого даем куки
		return createCookie(rw, r)
	}

	data := cartStorage.Get(propusk.Value) //проверяем хеш на существование
	if data == nil {                       //если данные =нил то ошибка(кука протухла)
		return createCookie(rw, r)
	}

	sd, ok := data.(*SessionData) //
	if !ok {
		panic("a-a-a-a!")
	}
	return sd
}

func createCookie(rw http.ResponseWriter, r *http.Request) *SessionData {

	propusk := &http.Cookie{}
	propusk.Name = "propusk"     //добавляем имя "пропуск"
	propusk.Value = CreateHash() //добавляем функцию генирирования хеш
	propusk.Path = "/"           //куки для всех путей
	http.SetCookie(rw, propusk)  //посылаем браузеру куки с хеш
	sd := &SessionData{}
	//// r.RemoteAddr == "127.0.0.1:44595" //ipv4
	//// r.RemoteAddr == "[::1]:44595" //ipv6
	index := strings.LastIndex(r.RemoteAddr, ":")

	if index == -1 {
		panic("yebat, sho takoe???")
	}
	ip := r.RemoteAddr[:index]
	sd.Ip = ip //берем IP

	http.SetCookie(rw, propusk)
	cartStorage.Set(propusk.Value, sd) //первый параметр пропусквалюе,вторым сесию

	return sd

	//propusk := &http.Cookie{}
	//propusk.Name = "propusk"
	//propusk.Value = CreateHash()
	//
	//sd := &SessionData{}
	//// TODO: remove port from ip:port
	//// r.RemoteAddr == "127.0.0.1:44595" //ipv4
	//// r.RemoteAddr == "[::1]:44595" //ipv6
	////index := strings.LastIndex(r.RemoteAddr, ":")
	////if index == -1 {
	////  panic("yebat, sho takoe???")
	////}
	////ip := r.RemoteAddr[:index]
	////sd.IP = ip
	//sd.IP =r.RemoteAddr
	//
	//http.SetCookie(rw, propusk)
	//cartStorage.Set(propusk.Value, &sd)
	//return sd
}

func CreateHash() string {

	// FIXME
	now := time.Now().UnixNano()

	rnd := rand.Int63()

	randomString := fmt.Sprintf("%d_%d", now, rnd) //.Sprintf-возвращает в переменую(переводит int в строку)

	fmt.Printf("rnd to hash:%s\n", randomString) //смотрим рандомхеш

	randomHashBytes := sha1.Sum([]byte(randomString)) //получаем hash(хешировали)возвращает 20 байт

	randomHash := fmt.Sprintf("%x", randomHashBytes) //X- 16ричное(масив байт переводим в 16ричное)40 байт

	fmt.Printf("LEN:%d \nhash:%s\n", len(randomHash), randomHash) //смотрим хеш

	return randomHash
}
