package main

type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type WeatherSubject struct {
	weatherData *WeatherData
	observers   []Observer
}

func NewWeatherSubject() *WeatherSubject {
	return &WeatherSubject{
		weatherData: nil,
		observers:   []Observer{},
	}
}

func (ws *WeatherSubject) AddObserver(observer Observer) {
	ws.observers = append(ws.observers, observer)
}

func (ws *WeatherSubject) RemoveObserver(observer Observer) {
	var newObservers []Observer
	for _, obs := range ws.observers {
		if obs != observer {
			newObservers = append(newObservers, obs)
		}
	}
	ws.observers = newObservers
}

func (ws *WeatherSubject) NotifyObservers() {
	for _, observer := range ws.observers {
		observer.Update(ws.weatherData)
	}
}

func (ws *WeatherSubject) SetWeatherData(weatherData *WeatherData) {
	ws.weatherData = weatherData
	ws.NotifyObservers()
}
