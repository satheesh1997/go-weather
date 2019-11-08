package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/martinlindhe/notify"
)

type Coordinate struct {
	Lon int
	Lat int
}

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type Main struct {
	Temp     float64
	Pressure int
	Humidity int
	Temp_min float64
	Temp_max float64
}

type Wind struct {
	Speed float64
	Deg   int
}

type Clouds struct {
	All int
}

type Sys struct {
	Type    int
	Id      int
	Country string
	Sunrise int
	Sunset  int
}

type WeatherData struct {
	Coord      Coordinate
	Weather    []Weather
	Base       string
	Main       Main
	Visibility int
	Wind       Wind
	Clouds     Clouds
	Dt         int
	Sys        Sys
	Timezone   int
	Id         int
	Name       string
	Cod        int
}

const (
	weatherAPI = "http://api.openweathermap.org/data/2.5/weather"
	appID      = "7846249a6836f94a21cbacc960aa53a6"
)

func getWeatherByCord(lat string, long string) WeatherData {
	var weatherData WeatherData
	uri := weatherAPI + "?lat=" + lat + "&lon=" + long
	uri = uri + "&APPID=" + appID
	resp, err := http.Get(uri)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", uri, err)
		os.Exit(1)
	}
	json.Unmarshal([]byte(data), &weatherData)
	return weatherData
}

func convertTemperature(k float64) float64 {
	return float64(k - 275.15)
}

func main() {
	weather := getWeatherByCord("12.971599", "77.594566")

	temperature := strconv.FormatFloat(convertTemperature(weather.Main.Temp), 'f', -1, 32)
	msg := "\nTemperature: " + temperature + "°C "

	wData := weather.Weather[0]
	msg = msg + "(" + wData.Description + ")"

	wind_data := weather.Wind
	speed := strconv.FormatFloat(wind_data.Speed, 'f', -1, 32)
	msg = msg + "\nWind: " + speed + " mph"

	cloudData := weather.Clouds
	cloudiness := strconv.Itoa(cloudData.All)
	msg = msg + "\nCloudiness: " + cloudiness + "%"

	notify.Notify("Weather Forecast", "Hourly Weather Update", msg, "")
}
