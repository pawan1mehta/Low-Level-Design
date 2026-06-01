package ObserverCode;

public class CurrentConditionsDisplayObserver implements Observer {
    private WeatherSubject weatherSubject;

    public CurrentConditionsDisplayObserver(WeatherSubject weatherSubject) {
        this.weatherSubject = weatherSubject;
        weatherSubject.registerObserver(this);
    }

    public void update(WeatherData weatherData) {
        System.out.println("Received data in CurrentConditionsDisplayObserver");
        businessLogic(weatherData);
    }

    public void businessLogic(WeatherData weatherData) {
        System.out.println("CurrentConditionsDisplayObserver: ");
        System.out.println(weatherData);
    }
}