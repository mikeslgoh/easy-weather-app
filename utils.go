package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func writeJsonToFile(filePath string, jsonString string) {
	os.WriteFile(filePath, []byte(jsonString), os.ModePerm)
}

func processRequest(requestString string) string {
	resp, err := http.Get(requestString)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	data := location{}
	err = data.UnmarshalLocation([]byte(body))
	if err != nil {
		log.Println(err)
		return ""
	}
	j_data, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(j_data)
}

func (loc *location) UnmarshalLocation(data []byte) error {
	respData := map[string]any{}
	if err := json.Unmarshal(data, &respData); err != nil {
		return err
	}
	locationData := respData["location"].(map[string]interface{})
	loc.Name = locationData["name"].(string)
	loc.Region = locationData["region"].(string)
	loc.Country = locationData["country"].(string)

	metricData := forecast{}
	err := metricData.UnmarshalCurrentData(data, metricMeasurement)
	if err != nil {
		return err
	}
	loc.CurrentMetricData = metricData

	imperialData := forecast{}
	err = imperialData.UnmarshalCurrentData(data, imperialMeasurement)
	if err != nil {
		return err
	}
	loc.CurrentImperialData = imperialData

	var forecastData []forecastDay
	forecastData, err = UnmarshalForecastDayData(data)
	if err != nil {
		return err
	}
	loc.FutureForecastData = forecastData

	return nil
}

func (fc *forecast) UnmarshalCurrentData(data []byte, measurement measurement) error {
	respData := map[string]any{}
	if err := json.Unmarshal(data, &respData); err != nil {
		return err
	}
	currentWeatherData, ok := respData["current"].(map[string]interface{})
	if !ok {
		currentWeatherData = respData
	}
	lastUpdated, ok := currentWeatherData["last_updated"].(string)
	if ok {
		fc.LastUpdated = lastUpdated
	}
	time, ok := currentWeatherData["time"].(string)
	if ok {
		fc.Time = time
	}
	chanceOfRain, ok := currentWeatherData["chance_of_rain"].(float64)
	if ok {
		fc.ChanceOfRain = chanceOfRain
	}

	fc.Temperature = currentWeatherData[fmt.Sprintf("temp_%v", measurement.temperature)].(float64)
	fc.WeatherCondition = currentWeatherData["condition"].(map[string]interface{})
	fc.Humidity = currentWeatherData["humidity"].(float64)
	fc.Uv = currentWeatherData["uv"].(float64)
	fc.WindSpeed = currentWeatherData[fmt.Sprintf("wind_%v", measurement.windSpeed)].(float64)
	fc.AirPressure = currentWeatherData[fmt.Sprintf("pressure_%v", measurement.pressure)].(float64)
	fc.PrecipitationLevel = currentWeatherData[fmt.Sprintf("precip_%v", measurement.precipLevel)].(float64)
	fc.FeelsLike = currentWeatherData[fmt.Sprintf("feelslike_%v", measurement.feelsLike)].(float64)
	fc.HeatIndex = currentWeatherData[fmt.Sprintf("heatindex_%v", measurement.heatIndex)].(float64)
	fc.DewPoint = currentWeatherData[fmt.Sprintf("dewpoint_%v", measurement.dewPoint)].(float64)
	return nil
}

func UnmarshalForecastDayData(data []byte) ([]forecastDay, error) {
	respData := map[string]any{}
	if err := json.Unmarshal(data, &respData); err != nil {
		return nil, err
	}

	forecastDayList := []forecastDay{}
	rawForecast, ok := respData["forecast"].(map[string]interface{})
	if !(ok) {
		log.Println("No forecast data found!")
		return forecastDayList, nil
	}
	rawForecastDayListData := rawForecast["forecastday"].([]interface{})

	for _, forecastDayData := range rawForecastDayListData {
		fcd := &forecastDay{}
		hourlyData := forecastDayData.(map[string]interface{})["hour"].([]interface{})
		encodedForecastDay, _ := json.Marshal(forecastDayData)
		json.Unmarshal(encodedForecastDay, &fcd)

		for _, forecastHour := range hourlyData {
			fc := forecast{}
			encodedForecastHour, _ := json.Marshal(forecastHour)
			fc.UnmarshalCurrentData(encodedForecastHour, metricMeasurement)
			fcd.HourMetric = append(fcd.HourMetric, fc)

			fc.UnmarshalCurrentData(encodedForecastHour, imperialMeasurement)
			fcd.HourImperial = append(fcd.HourImperial, fc)
		}

		forecastDayList = append(forecastDayList, *fcd)
	}
	return forecastDayList, nil
}
