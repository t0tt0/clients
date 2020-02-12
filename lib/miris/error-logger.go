package miris

type ErrorLogger interface {
	Error(msg string, keyvals ...interface{})
}

var errLogger ErrorLogger = defaultErrorLogger{}

func SetErrorLogger(logger ErrorLogger) {
	errLogger = logger
}

type defaultErrorLogger struct{}

func (d defaultErrorLogger) Error(_ string, _ ...interface{}) {
	// ignore this
}

func procSerializeError(_ int, err error) {
	if err != nil {
		errLogger.Error("serialize error", "error", err)
	}
}
