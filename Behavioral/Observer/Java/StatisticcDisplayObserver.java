package ObserverCode;

public class StatisticcDisplayObserver implements Observer {
    private WeatherSubject weatherSubject;

    public StatisticcDisplayObserver(WeatherSubject weatherSubject) {
        this.weatherSubject = weatherSubject;
        weatherSubject.registerObserver(this);
    }

    public void update(WeatherData weatherData) {
        System.out.println("Received data in StatisticcDisplayObserver");
        businessLogic(weatherData);
    }

    public void businessLogic(WeatherData weatherData) {
        System.out.println("StatisticcDisplayObserver: ");
        System.out.println(weatherData);
    }
}