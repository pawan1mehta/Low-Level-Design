public class JsonFormatter implements Formatter {

    @Override
    public String format(LogRecord logRecord) {
        return "{\"timestamp\":\"" + logRecord.getTimestamp()
                + "\",\"level\":\"" + logRecord.getLevel()
                + "\",\"thread\":\"" + logRecord.getThreadName()
                + "\",\"message\":\"" + logRecord.getMessage() + "\"}";
    }
}
