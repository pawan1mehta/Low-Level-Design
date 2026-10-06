public enum Level {
    DEBUG, INFO, WARN, ERROR, FATAL;

    public boolean isEnabledFor(Level threshold) {
        return this.ordinal() >= threshold.ordinal();
    }
}
