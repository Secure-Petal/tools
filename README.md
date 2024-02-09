# tools
Tools repository for our services. Publicly available for everyone to improve and use.

## Logger
Our logger instance is based on `slog` package. To use it:
```
	logger.InitLogger(slog.LevelInfo)
	logger.GetLogger().InfoContext(context.Background(), "Hello World", "name", "Secure Petal")
```

## Server
To start the server
```
ginRouter := gin.Default()

server.New(
    ginRouter,
    server.WithPort(8080),
).Run()
```