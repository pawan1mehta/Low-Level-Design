public class Main {
    public static void main(String[] args) {
        LoggerConfig loggerConfig = new LoggerConfig();
        loggerConfig.setFormatter(new JsonFormatter());
        loggerConfig.setLevel(Level.INFO);
        loggerConfig.setPath("/Users/pm57149/Downloads/mylogger");
        loggerConfig.setFlushThresholdBytes(1);

        Logger logger = LoggerFactory.createFileLogger(loggerConfig);

        logger.debug("debug message");
        logger.info("debug message");
        logger.warn("debug message");
        logger.error("debug message");
        logger.fatal("debug message");
    }
}
