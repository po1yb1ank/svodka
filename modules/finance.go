package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type curRes struct {
	Date 		string 	`json:"Date"`
	Valute 		vals 	`json:"Valute"`
}
type vals struct {
	USD			currencies `json:"USD"`
	EUR 		currencies `json:"EUR"`
	NOK 		currencies `json:"NOK"`
}
type currencies struct {
	CharCode 	string `json:"CharCode"`
	Name 		string `json:"Name"`
	Value 		float64 `json:"Value"`
	Previous	float64 `json:"Previous"`
}
func CurrencyStats(){
	c := curRes{}
	res, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil{
		log.Fatal("error resolving currency")
	}
	defer res.Body.Close()
	currJSON, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(currJSON, &c)
	if err != nil{
		log.Fatal("error while unmarshalling json", err)
	}
	printer(c)
}
func printer(c curRes){
	fmt.Println("Date: ", c.Date)
	fmt.Println(c.Valute.EUR.Name,":", c.Valute.EUR.Value, ". Previous:", c.Valute.EUR.Previous)
	fmt.Println(c.Valute.USD.Name,":", c.Valute.USD.Value, ". Previous:", c.Valute.USD.Previous)
	fmt.Println(c.Valute.NOK.Name,":", c.Valute.NOK.Value, ". Previous:", c.Valute.NOK.Previous)
}
