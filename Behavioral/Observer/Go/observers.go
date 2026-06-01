package main

import "fmt"

type Observer interface {
	Update(data *WeatherData)
}

type CurrentConditionsDisplayObserver struct {
	weatherData *WeatherData
}

func NewCurrentConditionsDisplayObserver(weatherData *WeatherData) *CurrentConditionsDisplayObserver {
	return &CurrentConditionsDisplayObserver{
		weatherData: weatherData,
	}
}

func (c *CurrentConditionsDisplayObserver) Update(weatherData *WeatherData) {
	fmt.Printf("Received data in CurrentConditionsDisplayObserver. data: %v \n", weatherData)
	c.BusinessLogic(weatherData)
}

func (c *CurrentConditionsDisplayObserver) BusinessLogic(weatherData *WeatherData) {
	fmt.Println("CurrentConditionsDisplayObserver BusinessLogic!")
}

type StaticDisplayObserver struct {
	weatherData *WeatherData
}

func NewStaticDisplayObserver(weatherData *WeatherData) *StaticDisplayObserver {
	return &StaticDisplayObserver{
		weatherData: weatherData,
	}
}

func (c *StaticDisplayObserver) Update(weatherData *WeatherData) {
	fmt.Printf("Received data in StaticDisplayObserver. data: %v \n", weatherData)
	c.BusinessLogic(weatherData)
}

func (c *StaticDisplayObserver) BusinessLogic(weatherData *WeatherData) {
	fmt.Println("StaticDisplayObserver BusinessLogic!")
}
