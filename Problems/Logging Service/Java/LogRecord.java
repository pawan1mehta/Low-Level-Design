import java.time.Instant;

public class LogRecord {
    private final Instant timestamp;
    private final Level level;
    private final String message;
    private final String threadName;

    public LogRecord(Level level, String message) {
        this.timestamp = Instant.now();
        this.level = level;
        this.message = message;
        this.threadName = Thread.currentThread().getName();
    }

    public Instant getTimestamp() {
        return timestamp;
    }

    public Level getLevel() {
        return level;
    }

    public String getMessage() {
        return message;
    }

    public String getThreadName() {
        return threadName;
    }
}
