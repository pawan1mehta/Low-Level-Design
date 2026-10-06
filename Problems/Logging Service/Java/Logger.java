abstract class Logger {
    protected volatile Level level;

    public void setLevel(Level level) {
        this.level = level;
    }

    public Level getLevel() {
        return level;
    }

    public void debug(String message) {
        log(Level.DEBUG, message);
    }

    public void info(String message) {
        log(Level.INFO, message);
    }

    public void warn(String message) {
        log(Level.WARN, message);
    }

    public void error(String message) {
        log(Level.ERROR, message);
    }

    public void fatal(String message) {
        log(Level.FATAL, message);
    }

    private void log(Level level, String message) {
        if(!level.isEnabledFor(this.level)) {
            return;
        }
        write(new LogRecord(level, message));
    }

    protected abstract void write(LogRecord logRecord);
}
