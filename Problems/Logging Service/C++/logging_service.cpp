#include <chrono>
#include <fstream>
#include <iomanip>
#include <iostream>
#include <memory>
#include <mutex>
#include <sstream>
#include <string>
#include <thread>
#include <vector>
#include <filesystem>

using namespace std;

enum class Level : uint8_t{DEBUG, INFO, WARNING, ERROR, FATAL};

constexpr bool isEnabledFor(const Level level, const Level threshold) {
    return level >= threshold;
}

string toString(const Level level) {
    switch (level) {
        case Level::DEBUG: return "DEBUG";
        case Level::INFO: return "INFO";
        case Level::WARNING: return "WARNING";
        case Level::ERROR: return "ERROR";
        case Level::FATAL: return "FATAL";
        default: return "";
    }
}

struct LogRecord {
    string timestamp;
    Level level;
    string message;
    string threadName;

    LogRecord(const Level level, const string& message) {
        this->timestamp = now();
        this->level = level;
        this->message = message;
        this->threadName = threadId();
    }

    static string now() {
        ostringstream out;
        auto t = chrono::system_clock::to_time_t(chrono::system_clock::now());
        out << put_time(localtime(&t), "%Y-%m-%d %H:%M:%S");
        return out.str();
    }

    static string threadId() {
        ostringstream out;
        out << this_thread::get_id();
        return out.str();
    }
};

class LogFile {
private:
    filesystem::path path;
    ofstream out;
    mutex mtx;

public:
    LogFile(const string& pathStr) : path(pathStr){
        if (path.has_parent_path()) {
            filesystem::create_directories(path.parent_path());
        }
        out.open(path, ios::out | ios::app | ios::binary);
        if (!out) {
            throw runtime_error("Cannot open file " + pathStr);
        }
    }

    void write(const string& data) {
        lock_guard<mutex> lock(mtx);
        out.write(data.data(), data.size());
        if (!out) {
            throw runtime_error("write failed");
        }
    }

    void flush() {
        lock_guard<mutex> lock(mtx);
        out.flush();
        if (!out) {
            throw runtime_error("flush failed");
        }
    }
};

class Formatter {
public:
    virtual ~Formatter() = default;
    virtual string format(const LogRecord& logRecord) const = 0;
};

class TextFormatter : public Formatter {
public:
    string format(const LogRecord& r) const override {
        ostringstream out;
        out << r.timestamp << " [" << toString(r.level) << "] ["
            << r.message << "] " << r.threadName;
        return out.str();
    }
};

class JsonFormatter : public Formatter {
public:
    string format(const LogRecord& r) const override {
        ostringstream out;
        out << "{"
            << "\"timestamp\":\"" << r.timestamp << "\","
            << "\"level\":\""     << toString(r.level) << "\","
            << "\"message\":\""   << r.message << "\","
            << "\"thread\":\""    << r.threadName << "\""
            << "}";
        return out.str();
    }
};

class Logger {
protected:
    Level level;

public:
    virtual ~Logger() = default;

    virtual void debug(string message) {
        log(Level::DEBUG, message);
    }

    virtual void info(string message) {
        log(Level::INFO, message);
    }

    virtual void warn(string message) {
        log(Level::WARNING, message);
    }

    virtual void error(string message) {
        log(Level::ERROR, message);
    }

    virtual void fatal(string message) {
        log(Level::FATAL, message);
    }

    void setLevel(Level level) {
        this->level = level;
    }

    Level getLevel() const {
        return level;
    }

    virtual void write(const Level level, const string& message) = 0;

private:
    void log(const Level level, const string& message) {
        if (!isEnabledFor(level, getLevel())) {
            return;
        }
        write(level, message);
    }
};

class ConsoleLogger : public Logger {
    unique_ptr<Formatter> formatter;
public:

    ConsoleLogger() {
        formatter = make_unique<TextFormatter>();
    }

    explicit ConsoleLogger(unique_ptr<Formatter> formatter) : formatter(std::move(formatter)) {}
    
    void write(const Level level, const string& message) override {
        string line = formatter->format(LogRecord(level, message));
        cout << line << endl;
    }
};

class InMemoryLogger : public Logger {
    unique_ptr<Formatter> formatter;
    vector<LogRecord> records;
public:

    InMemoryLogger() {
        formatter = make_unique<TextFormatter>();
    }

    explicit InMemoryLogger(unique_ptr<Formatter> formatter) : formatter(std::move(formatter)) {}

    void write(const Level level, const string& message) override {
        records.push_back(LogRecord(level, message));
    }
};

class FileLogger : public Logger {
    unique_ptr<Formatter> formatter;
    unique_ptr<LogFile> file;
    int flushThresholdBytes = 1024;
    int bufferedBytes = 0;
public:

    FileLogger() {
        formatter = make_unique<TextFormatter>();
        file = make_unique<LogFile>("/tem");
    }

    FileLogger(unique_ptr<Formatter> formatter, string path, int flushThresholdBytes)
        : formatter(std::move(formatter)),
          file(make_unique<LogFile>(path)),
          flushThresholdBytes(flushThresholdBytes) {}

    void write(const Level level, const string& message) override{
        ostringstream out;
        out << formatter->format(LogRecord(level, message));
        file->write(out.str());
        bufferedBytes += out.str().length();
        if (bufferedBytes >= flushThresholdBytes) {
            file->flush();
            bufferedBytes = 0;
        }
    }
};

struct LoggerConfig {
    Level level;
    unique_ptr<Formatter> formatter;
    string path;
    int flushThresholdBytes;
};

class LoggerFactory {
public:
    static unique_ptr<Logger> createFileLogger(LoggerConfig config) {
        unique_ptr<Logger> logger = make_unique<FileLogger>(
            std::move(config.formatter),
            config.path,
            config.flushThresholdBytes
        );
        logger->setLevel(config.level);
        return logger;
    }

};

int main() {
    LoggerConfig config;
    config.level = Level::INFO;
    config.formatter = make_unique<JsonFormatter>();
    config.path = "/Users/pm57149/Downloads/mylogger";
    config.flushThresholdBytes = 1;

    unique_ptr<Logger> logger = LoggerFactory::createFileLogger(std::move(config));

    logger->debug("debug message");
    logger->info("debug message");
    logger->warn("debug message");
    logger->error("debug message");
    logger->fatal("debug message");

    return 0;
}