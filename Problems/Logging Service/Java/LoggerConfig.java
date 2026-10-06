public class LoggerConfig {
    private String path;
    private int flushThresholdBytes;
    private Formatter formatter;
    private Level level;

    public String getPath() {
        return path;
    }

    public void setPath(String path) {
        this.path = path;
    }

    public int getFlushThresholdBytes() {
        return flushThresholdBytes;
    }

    public void setFlushThresholdBytes(int flushThresholdBytes) {
        this.flushThresholdBytes = flushThresholdBytes;
    }

    public Formatter getFormatter() {
        return formatter;
    }

    public void setFormatter(Formatter formatter) {
        this.formatter = formatter;
    }

    public Level getLevel() {
        return level;
    }

    public void setLevel(Level level) {
        this.level = level;
    }
}
