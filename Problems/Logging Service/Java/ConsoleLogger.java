public class ConsoleLogger extends Logger {
    private final Formatter formatter;

    public ConsoleLogger() {
        this.formatter = new TextFormatter();
    }

    public ConsoleLogger(Formatter formatter) {
        this.formatter = formatter;
    }

    protected synchronized  void write(LogRecord logRecord) {
        System.out.println(formatter.format(logRecord));
    }
}
