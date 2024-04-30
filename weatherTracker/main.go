package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string "json:OpenWeatherMapApiKey"
}

type weatherData struct {
	Name string `json: "name"`
	Main struct {
		Kelvin float64 `json: "temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err

	}
	var c apiConfigData

	err = json.Unmarshal(bytes, &c)

	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, There are no More Guns in the Valley!\n"))
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")

	if err != nil {
		return weatherData{}, err
	}

	//resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + ",uk&APPID=" + apiConfig.OpenWeatherMapApiKey)

	//resp, err := http.Get("https://api.openweathermap.org/data/3.0/onecall/timemachine?lat=39.099724&lon=-94.578331&dt=1643803200&appid={8ec76989ab8fe5a3b839fa58242f0fff}")
	//resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "APPID=" + apiConfig.OpenWeatherMapApiKey)

	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(city)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder((w)).Encode(data)
		})

	http.ListenAndServe(":8080", nil)

}
