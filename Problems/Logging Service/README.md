# Logging Service

## Clarifying Questions

- Will it be in-memory logging service or ships logs over network? [Ans: in-memory]
- What severity level should we support, and is there any ordering between them? [Ans: DEBUG, INFO, WARN, ERROR, FATAL]
- Where to persiste the logs? [Ans: file]
- Logger will write to multiple destination console & file? [Ans: Yes]
- Can Each destination decide it's own filter? [Ans: Yes for console it would be from DEBUG, and for file it would be from WARN]
- Is the logging format the records are written it. is that fixed or does it vary? [Ans: Yes sometime texts or sometime jsob the format is idependent of the destination type]
- Handle the concurrency.
- Is the configuration static? [Ans: static configured at startup]

## Requirements

- Caller should be able to use different severity methods[DEBUG, INFO, WARN, ERROR, FATAL] to print logs
- Order of the severity is DEBUG < INFO < WARN < ERROR < FATAL
- Each log record should carries: timestamp, level, message, emitting thread name.
- System should be able to persiste the logs in file if it is configured on the startup.
  - after some thresold value the system should flush the file into disk.
- System should be able to previde the the history of logs
- Handle the concurrency

## Entities & Relationships

- Logger
- ConsoleLogger
- FileLogger
- InMemoryLogger

- LogRecord

- LogFile

- Formatter
- TextFormatter
- JsonFormatter

- LoggerFactory

Relationships:

```text
    ConsoleLogger ------inherits----> Logger
    FileLogger ------inherits----> Logger
    InMemoryLogger ------inherits----> Logger

    TextFormatter ------implements-----> Formatter
    JsonFormatter ------implements-----> Formatter

    ConsoleLogger <-----composed of------ Formatter

    InMemoryLogger <----- composed of ------ Formatter
    InMemoryLogger <----- composed of ------ LogRecord List

    FileLogger <----- composed of ------ Formatter
    FileLogger <----- composed of ------ LogFile
```

## Class Design

```code
Class LoggerFactory:

    + createFileLogger(config): Logger
```

```code
Abstract Class Logger:

    - level: (DEBUG, INFO, WARN, ERROR, FATAL) default DEBUG

    + info(message: string)
    + debug(message: string)
    + warn(message: string)
    + error(message: string)
    + fatal(message: string)
```

```code
Class ConsoleLogger inherits Logger:

    - formatter: JSON (Text)

```

```code
Class FileLogger inherits Logger:

    - formatter: JSON (JSONFormatter)
    - logFile: LogFile
    - flushThresholdBytes: int
    
    - write(record)
    + flush()
```

```code
Class InMemoryLogger inherits Logger:

    - formatter: JSON (Text)
    - logs: LogRecord[]
```

```code
Class LogRecord:

    - timestamp
    - level
    - message
    - threadName
```

```code
Class LogFile:

    - path: (/var/log/mylogger)

    + write(bytes)
    + flush()
    + close()
```

```code
Inteface Formatter:

    + format(record): string
```
