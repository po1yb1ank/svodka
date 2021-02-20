package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	APIKEY = "5b1f8a100f176f261e7e6a4b7294af08"
	CITY = "Moscow"
	)
type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func WeatherStats() {
	c := http.DefaultClient
	req, err := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/weather?q="+CITY+"&appid="+APIKEY+"&units=metric&lang=ru", nil)
	if err != nil {
		log.Fatal("Can't create request", err)
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatal("Can't get response", err)
	}

	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)

	w := Weather{}
	err = json.Unmarshal(buf, &w)
	if err != nil {
		log.Fatal("Can't parse response: ", err)
	}
	weatherPrinter(w)
}

func weatherPrinter(w Weather)  {
	fmt.Println("=== Погода в", w.Name, "===")
	for _, v := range w.Weather{
		fmt.Println(v.Main,":", v.Description)
	}
	fmt.Println("Температура:", w.Main.Temp,", ощущается:", w.Main.FeelsLike)
	fmt.Println("Влажность:", strconv.Itoa(w.Main.Humidity) + "%",", давление:", w.Main.Pressure)
	fmt.Println("Скорость ветра:", w.Wind.Speed,"м/c", ", облачность:", strconv.Itoa(w.Clouds.All)+"%")

}
