package main

func main() {

	weatherData1 := NewWeatherData(12.0, 234.4, 67.6)

	currentConditionDisplayObserver := NewCurrentConditionsDisplayObserver(weatherData1)
	staticDisplayObserver := NewStaticDisplayObserver(weatherData1)

	weatherSubject := NewWeatherSubject()

	weatherSubject.AddObserver(currentConditionDisplayObserver)
	weatherSubject.AddObserver(staticDisplayObserver)

	weatherData2 := NewWeatherData(62.0, 204.4, 77.6)
	
	weatherSubject.SetWeatherData(weatherData2)
}
