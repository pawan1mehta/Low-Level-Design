public class FileLogger extends Logger {
    private final Formatter formatter;
    private final LogFile logFile;
    private final int flushThresholdBytes;
    private int bufferedBytes = 0;

    public FileLogger(Formatter formatter, LogFile logFile, int flushThresholdBytes) {
        this.formatter = formatter;
        this.logFile = logFile;
        this.flushThresholdBytes = flushThresholdBytes;
    }

    @Override
    protected void write(LogRecord logRecord) {
        byte[] line = formatter.format(logRecord).getBytes();
        logFile.write(line);
        bufferedBytes += line.length;
        if(bufferedBytes >= flushThresholdBytes) {
            flush();
        }
    }

    private void flush() {
        logFile.flush();
        bufferedBytes = 0;
    }
}
