package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// OpenWeatherAPI struct
type OpenWeatherAPI struct {
	Coord      typeCoord     `json:"coord"`
	Weather    []typeWeather `json:"weather"`
	Base       string        `json:"base"`
	Main       typeMain      `json:"main"`
	Visibility int           `json:"visibility"`
	Wind       typeWind      `json:"wind"`
	Clouds     typeClouds    `json:"clouds"`
	Dt         int           `json:"dt"`
	Sys        typeSys       `json:"sys"`
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	COD        int           `json:"cod"`
}

type typeCoord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type typeWeather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type typeMain struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type typeWind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type typeClouds struct {
	All int `json:"all"`
}

type typeSys struct {
	Tyoe    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// end OpenWeatherAPI struct

// AccuWeatherAPI struct
type AccuWeatherAPI struct {
	LocalObservationDateTime string          `json:"LocalObservationDateTime"`
	EpochTime                int             `json:"EpochTime"`
	WeatherText              string          `json:"WeatherText"`
	WeatherIcon              int             `json:"WeatherIcon"`
	IsDayTime                bool            `json:"IsDayTime"`
	Temperature              typeTemperature `json:"Temperature"`
}

type typeTemperature struct {
	Metric typeTemperatureUnit `json:"Metric"`
}

type typeTemperatureUnit struct {
	Value    float64 `json:"Value"`
	Unit     string  `json:"Unit"`
	UnitType int     `json:"UnitType"`
}

// end AccuWeatherAPI struct

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func clientHTTP(apiKey, url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func getCurrentOpenWeather(t time.Time) {
	apiKey := "b300bd5251fbf44dc2e8207e30d7f3ef"
	url := "https://api.openweathermap.org/data/2.5/weather?id=1580578&units=metric&appid=" + apiKey

	body := clientHTTP(apiKey, url)

	var openWeatherAPI OpenWeatherAPI
	if err := json.Unmarshal(body, &openWeatherAPI); err != nil {
		panic(err)
	}

	f, err := os.OpenFile("./temp/open.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	text := strconv.Itoa(openWeatherAPI.Dt) + ": "
	text = text + strconv.FormatFloat(openWeatherAPI.Main.Temp, 'f', -1, 64) + "\n"
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func getCurrentAccuWeather(t time.Time) {
	apiKey := "if3CzCIQRVH3yCXKz7Qe1cSspM34ZjbH"
	url := "http://dataservice.accuweather.com/currentconditions/v1/1-353981_1_AL?apikey=" + apiKey

	body := clientHTTP(apiKey, url)

	var accuWeatherAPI []AccuWeatherAPI
	if err := json.Unmarshal(body, &accuWeatherAPI); err != nil {
		panic(err)
	}

	f, err := os.OpenFile("./temp/accu.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, acc := range accuWeatherAPI {
		text := strconv.Itoa(acc.EpochTime) + ": "
		text = text + strconv.FormatFloat(acc.Temperature.Metric.Value, 'f', -1, 64) + "\n"
		if _, err = f.WriteString(text); err != nil {
			panic(err)
		}
	}

}

func getCurrentWeather(t time.Time) {
	// getCurrentAccuWeather(t)
	getCurrentOpenWeather(t)
}

func main() {
	if _, err := os.Stat("./temp/accu.txt"); err == nil {
		os.Remove("./temp/accu.txt")
	}
	if _, err := os.Stat("./temp/open.txt"); err == nil {
		os.Remove("./temp/open.txt")
	}
	doEvery(time.Second, getCurrentWeather)
}
