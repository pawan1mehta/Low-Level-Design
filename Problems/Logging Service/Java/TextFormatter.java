public class TextFormatter implements Formatter {

    @Override
    public String format(LogRecord logRecord) {
        return String.format("%s [%s] [%s] %s",
                logRecord.getTimestamp(), logRecord.getLevel(), logRecord.getMessage(), logRecord.getThreadName());
    }
}
