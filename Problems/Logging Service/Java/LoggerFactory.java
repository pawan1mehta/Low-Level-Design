public class LoggerFactory {

    static public Logger createFileLogger(LoggerConfig config) {
        FileLogger logger = new FileLogger(
                config.getFormatter(),
                new LogFile(config.getPath()),
                config.getFlushThresholdBytes());
        logger.setLevel(config.getLevel());
        return logger;
    }

    static public Logger createConsoleLogger() {
        return new ConsoleLogger();
    }

    static public Logger createInMemoryLogger() {
        return new InMemoryLogger();
    }
}
