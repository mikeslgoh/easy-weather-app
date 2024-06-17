package main

type measurement struct {
	temperature string
	windSpeed   string
	pressure    string
	precipLevel string
	feelsLike   string
	windChill   string
	heatIndex   string
	dewPoint    string
}

var metricMeasurement = measurement{
	temperature: "c",
	windSpeed:   "kph",
	pressure:    "mb",
	precipLevel: "mm",
	feelsLike:   "c",
	windChill:   "c",
	heatIndex:   "c",
	dewPoint:    "c",
}

var imperialMeasurement = measurement{
	temperature: "f",
	windSpeed:   "mph",
	pressure:    "in",
	precipLevel: "in",
	feelsLike:   "f",
	windChill:   "f",
	heatIndex:   "f",
	dewPoint:    "f",
}

type forecast struct {
	LastUpdated        string                 `json:"lastUpdated,omitempty"`
	Time               string                 `json:"time,omitempty"`
	Temperature        float64                `json:"temperature"`
	WeatherCondition   map[string]interface{} `json:"weatherCondition"`
	WindSpeed          float64                `json:"windSpeed"`
	AirPressure        float64                `json:"airPressure"`
	PrecipitationLevel float64                `json:"precipitationLevel"`
	Humidity           float64                `json:"humidity"`
	FeelsLike          float64                `json:"feelsLike"`
	WindChill          float64                `json:"windchill"`
	HeatIndex          float64                `json:"heatIndex"`
	DewPoint           float64                `json:"dewPoint"`
	ChanceOfRain       float64                `json:"chanceOfRain,omitempty"`
	Uv                 float64                `json:"uv"`
}

type forecastDay struct {
	Date string `json:"date"`
	Day  struct {
		Maxtemp_c            float64 `json:"maxtemp_c"`
		Maxtemp_f            float64 `json:"maxtemp_f"`
		Mintemp_c            float64 `json:"mintemp_c"`
		Mintemp_f            float64 `json:"mintemp_f"`
		Daily_chance_of_rain float64 `json:"daily_chance_of_rain"`
		Uv                   float64 `json:"uv"`
	} `json:"day"`
	Astro struct {
		Sunrise string `json:"sunrise"`
		Sunset  string `json:"sunset"`
	} `json:"astro"`
	HourMetric   []forecast `json:"hourMetric"`
	HourImperial []forecast `json:"hourImperial"`
}

type location struct {
	Name                string        `json:"name"`
	Region              string        `json:"region"`
	Country             string        `json:"country"`
	CurrentMetricData   forecast      `json:"currentMetricData"`
	CurrentImperialData forecast      `json:"currentImperialData"`
	FutureForecastData  []forecastDay `json:"futureForecastData"`
}
