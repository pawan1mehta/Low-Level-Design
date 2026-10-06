import java.util.ArrayList;
import java.util.List;

public class InMemoryLogger extends Logger {
    private final Formatter formatter;
    private final List<LogRecord> logs = new ArrayList<>();

    public InMemoryLogger() {
        this.formatter = new TextFormatter();
    }

    public InMemoryLogger(Formatter formatter) {
        this.formatter = formatter;
    }

    @Override
    protected synchronized void write(LogRecord logRecord) {
        logs.add(logRecord);
    }
}
