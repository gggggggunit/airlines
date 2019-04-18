package handlers

import (
	"fmt"
	"sync"
)

type Avia struct { //СТРУКТУРА ТАКАЯ ЖЕ КАК И ТАБЛИЦА В БД
	ID          int
	Source      string
	Destination string
	When        string
	Price       int
}

type AviaList struct {
	List        []*Avia
	ID          string
	Source      string //создана для того что б показивать потом в форме запрос который был
	Destination string //создана для того что б показивать потом в форме запрос который был
	When        string
	Price       string
	Pages   []    int
}

type Cart struct { //корзина(структура как БД)
	ID        int
	AirlineID int
	Price     int
	BuyDate   string
}

type SessionData struct {
	Cart []*Avia
	Ip   string
	Sum  int
	ID int
	Email string
	Ballance int
}

func (s *SessionData) Print() {
	fmt.Printf("ip: %v\n", s.Ip)
	fmt.Println("Cart:")

	for _, it := range s.Cart {
		fmt.Printf("\t%+v\n", it)
	}
	fmt.Printf("Login OK: %s,ID: %d\n", s.Email,s.ID)
}

type CartStorage struct {
	sync.RWMutex //Блокировка чтения НАЧАЛО
	Data         map[string]interface{}
}

func (s *CartStorage) Set(key string, data interface{}) { //функция добаления в хранилище
	s.Lock()           //лочим на запись
	s.Data[key] = data //добавляем в хранилище
	s.Unlock()         //разлочили
}

func (s *CartStorage) Get(key string) interface{} { //функция чтения из мапов
	s.RLock() //лочим на чтение
	defer s.RUnlock()
	return s.Data[key]

}
