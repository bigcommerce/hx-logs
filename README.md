# Logs

Versatile logging for Golang

[![GoDoc](https://godoc.org/github.com/hx/logs?status.svg)](https://godoc.org/github.com/hx/logs)

## Usage

```bash
go get -u github.com/hx/logs
```

Simple loggers only require a `level`:

```go
logger := logs.NewStdoutLogger(logs.Info)
logger.Info("I'm logging!")
```

Prints:

```
16 Nov 2020 02:33:19.402 UTC INFO  I'm logging!
```

More complex loggers can be composed:

```go
// Use a time.Location to localise log timestamps
location, err := time.LoadLocation("Australia/Sydney")
if err != nil {
	panic(err)
}

// Set a Level using user input
level, err := logs.LevelByName("warn")
if err != nil {
	panic(err)
}

// A TimeFormatter can be used by other Formatters
timeFormatter := logs.NewTimeFormatter(
	"2006-01-02 15:04",
	location,
)

// Make your own Formatter using a simple function
formatter := logs.FormatterFunc(func(event *logs.Event) []byte {
	return append(
		timeFormatter.Format(event),
		[]byte(fmt.Sprintf(" (%s) %s\n",
			event.Level.Abbreviation(),
			event.Message,
		))...,
	)
})

// A Publisher can send logs to several Subscribers
publisher := logs.NewPublisher(
	logs.NewWriterWithFormat(os.Stdout, formatter),
)

// Conditionally add more Subscribers to your publisher
if jsonPath := os.Getenv("JSON_LOG_PATH"); jsonPath != "" {
	jsonFile, err := os.Create(jsonPath)
	if err != nil {
		panic(err)
	}
	var messageID uint64
	
	// JsonFormatter produces Logstash-compatible output
	jsonFormatter := logs.NewJsonFormatter(logs.TagsFunc(func() logs.Tags {
		messageID += 1
		return logs.Tags{{"messageID", messageID}}
	}))
	publisher.Add(logs.NewWriterWithFormat(jsonFile, jsonFormatter))
}

// Use a Buffer to ensure log writing doesn't block on IO
logger := logs.NewLogger(level, logs.NewBuffer(1000, publisher))
```

See [the documentation](https://godoc.org/github.com/hx/logs) for the complete API.
