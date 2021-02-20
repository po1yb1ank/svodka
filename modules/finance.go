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
	currencyPrinter(c)
}
func currencyPrinter(c curRes){
	fmt.Println("=== Курс валюты===")
	fmt.Println("Дата и время: ", c.Date)
	fmt.Println(c.Valute.EUR.Name,":", c.Valute.EUR.Value, ". Предыдущее:", c.Valute.EUR.Previous)
	fmt.Println(c.Valute.USD.Name,":", c.Valute.USD.Value, ". Предыдущее:", c.Valute.USD.Previous)
	fmt.Println(c.Valute.NOK.Name,":", c.Valute.NOK.Value, ". Предыдущее:", c.Valute.NOK.Previous)
}
