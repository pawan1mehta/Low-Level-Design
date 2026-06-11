package ObserverCode;

public class Main {
    public static void main(String[] args) {
        // Create Subject
        WeatherSubject weatherSubject = new WeatherSubject();

        // Create Observers
        CurrentConditionsDisplayObserver observer1 = new CurrentConditionsDisplayObserver(weatherSubject);
        StatisticcDisplayObserver observer2 = new StatisticcDisplayObserver(weatherSubject);

        // Changes happens in subject
        WeatherData weatherData1 = new WeatherData(80, 65, 34);
        WeatherData weatherData2 = new WeatherData(180, 135, 103);

        wetherSubject.setWeatherData(weatherData1);
        wetherSubject.setWeatherData(weatherData2);
    }
}