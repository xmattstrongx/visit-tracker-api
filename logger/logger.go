package logger

import (
	stdlog "log"
	"os"

	kitlog "github.com/go-kit/kit/log"
)

func init() {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
}

func Message(message string) {
	stdlog.Printf(message)
}

