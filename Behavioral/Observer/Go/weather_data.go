package main

type WeatherData struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
}

func NewWeatherData(temperature, humidity, pressure float32) *WeatherData {
	return &WeatherData{
		Temperature: temperature,
		Humidity:    humidity,
		Pressure:    pressure,
	}
}
