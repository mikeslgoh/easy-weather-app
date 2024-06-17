package main

import (
	"fmt"
)

const API_KEY = "API KEY HERE"
const REQUEST_URL = "https://api.weatherapi.com/v1/"
const TODAY_WEATHER = "current.json"
const FORECAST_WEATHER = "forecast.json"

func getTodayWeather(city string) string {
	query := fmt.Sprintf("?key=%v&q=%v", API_KEY, city)
	return processRequest(REQUEST_URL + TODAY_WEATHER + query)
}

func getFutureForecast(city string, days int) string {
	query := fmt.Sprintf("?key=%v&q=%v&days=%v", API_KEY, city, days)
	return processRequest(REQUEST_URL + FORECAST_WEATHER + query)
}

func main() {
	days := 1
	city := "Vancouver"
	// jsonString := getTodayWeather(city)
	jsonString := getFutureForecast(city, days)
	writeJsonToFile("weatherData.json", jsonString)
}

/*
TODO:
- set up local webserver
- learn how to display UI -> dropdowns, tables, refresh button, lazyloading?
- User should be able to set number of days, city, unit of measurement
- Package this as a docker container
*/
