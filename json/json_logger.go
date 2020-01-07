package json

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"

	"github.com/gen0cide/genolog"
)

var (
	// Logger is a global singleton logger
	Logger genolog.Logger
)

var (
	defaultProg  = `APP`
	startName    = `cli`
	defaultLevel = logrus.InfoLevel

	global *logrus.Logger
)

type jsonLogger struct {
	internal *logrus.Logger
	prog     string
	context  string
}

// NewJSONLogger creates an genolog compatible console logger that prints to line delimited JSON.
func NewJSONLogger(prog, name string, existing genolog.Logger) genolog.Logger {
	if prog == "" {
		prog = defaultProg
	}

	if name == "" {
		name = startName
	}

	if existing != nil {
		existing.SetName(name)
		existing.SetProg(prog)
		return existing
	}

	if global == nil {
		global = logrus.New()
		global.SetLevel(defaultLevel)
		global.Out = color.Output
	}

	logger := &jsonLogger{
		internal: global,
		prog:     prog,
		context:  name,
	}

	global.Formatter = &logrus.JSONFormatter{}

	return logger
}

// ZapLogger implements the genolog.Logger interface.
func (j *jsonLogger) ZapLogger() *zap.SugaredLogger {
	return nil
}

// GetWriter implements the genolog.Logger interface.
func (j *jsonLogger) GetWriter() io.Writer {
	return j.internal.Out
}

// LogrusLogger implements the genolog.Logger interface.
func (j *jsonLogger) LogrusLogger() *logrus.Logger {
	return j.internal
}

// SetName implements the genolog.Logger interface.
func (j *jsonLogger) SetName(s string) {
	j.context = s
}

// SetProg implements the genolog.Logger interface.
func (j *jsonLogger) SetProg(s string) {
	j.prog = s
}

// GetOutput implements the genolog.Logger interface.
func (j *jsonLogger) GetOutput() io.Writer {
	return j.internal.Out
}

// SetOutput implements the genolog.Logger interface.
func (j *jsonLogger) SetOutput(w io.Writer) {
	j.internal.Out = w
}

// Print implements the genolog.Logger interface.
func (j *jsonLogger) Print(args ...interface{}) {
	j.internal.Print(args...)
	return
}

// Printf implements the genolog.Logger interface.
func (j *jsonLogger) Printf(format string, args ...interface{}) {
	j.internal.Printf(format, args...)
	return
}

// Println implements the genolog.Logger interface.
func (j *jsonLogger) Println(args ...interface{}) {
	j.internal.Println(args...)
	return
}

// Printw implements the genolog.Logger interface.
func (j *jsonLogger) Printw(body string, args ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(args); i++ {
		if i%2 == 0 || i == 0 {
			fields[fmt.Sprintf("%v", args[i])] = nil
			continue
		}
		fields[fmt.Sprintf("%v", args[i-1])] = args[i]
	}

	j.internal.WithFields(fields).Println(body)
	return
}

// Debug implements the genolog.Logger interface.
func (j *jsonLogger) Debug(args ...interface{}) {
	j.internal.Debug(args...)
	return
}

// Debugf implements the genolog.Logger interface.
func (j *jsonLogger) Debugf(format string, args ...interface{}) {
	j.internal.Debugf(format, args...)
	return
}

// Debugln implements the genolog.Logger interface.
func (j *jsonLogger) Debugln(args ...interface{}) {
	j.internal.Debugln(args...)
	return
}

// Debugw implements the genolog.Logger interface.
func (j *jsonLogger) Debugw(body string, args ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(args); i++ {
		if i%2 == 0 || i == 0 {
			fields[fmt.Sprintf("%v", args[i])] = nil
			continue
		}
		fields[fmt.Sprintf("%v", args[i-1])] = args[i]
	}

	j.internal.WithFields(fields).Debug(body)
	return
}

// Info implements the genolog.Logger interface.
func (j *jsonLogger) Info(args ...interface{}) {
	j.internal.Info(args...)
	return
}

// Infof implements the genolog.Logger interface.
func (j *jsonLogger) Infof(format string, args ...interface{}) {
	j.internal.Infof(format, args...)
	return
}

// Infoln implements the genolog.Logger interface.
func (j *jsonLogger) Infoln(args ...interface{}) {
	j.internal.Infoln(args...)
	return
}

// Infow implements the genolog.Logger interface.
func (j *jsonLogger) Infow(body string, args ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(args); i++ {
		if i%2 == 0 || i == 0 {
			fields[fmt.Sprintf("%v", args[i])] = nil
			continue
		}
		fields[fmt.Sprintf("%v", args[i-1])] = args[i]
	}

	j.internal.WithFields(fields).Info(body)
	return
}

// Warn implements the genolog.Logger interface.
func (j *jsonLogger) Warn(args ...interface{}) {
	j.internal.Warn(args...)
	return
}

// Warnf implements the genolog.Logger interface.
func (j *jsonLogger) Warnf(format string, args ...interface{}) {
	j.internal.Warnf(format, args...)
	return
}

// Warnln implements the genolog.Logger interface.
func (j *jsonLogger) Warnln(args ...interface{}) {
	j.internal.Warnln(args...)
	return
}

// Warnw implements the genolog.Logger interface.
func (j *jsonLogger) Warnw(body string, args ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(args); i++ {
		if i%2 == 0 || i == 0 {
			fields[fmt.Sprintf("%v", args[i])] = nil
			continue
		}
		fields[fmt.Sprintf("%v", args[i-1])] = args[i]
	}

	j.internal.WithFields(fields).Warn(body)
	return
}

// Error implements the genolog.Logger interface.
func (j *jsonLogger) Error(args ...interface{}) {
	j.internal.Error(args...)
	return
}

// Errorf implements the genolog.Logger interface.
func (j *jsonLogger) Errorf(format string, args ...interface{}) {
	j.internal.Errorf(format, args...)
	return
}

// Errorln implements the genolog.Logger interface.
func (j *jsonLogger) Errorln(args ...interface{}) {
	j.internal.Errorln(args...)
	return
}

// Errorw implements the genolog.Logger interface.
func (j *jsonLogger) Errorw(body string, args ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(args); i++ {
		if i%2 == 0 || i == 0 {
			fields[fmt.Sprintf("%v", args[i])] = nil
			continue
		}
		fields[fmt.Sprintf("%v", args[i-1])] = args[i]
	}

	j.internal.WithFields(fields).Error(body)
	return
}

// Fatal implements the genolog.Logger interface.
func (j *jsonLogger) Fatal(args ...interface{}) {
	j.internal.Fatal(args...)
	return
}

// Fatalf implements the genolog.Logger interface.
func (j *jsonLogger) Fatalf(format string, args ...interface{}) {
	j.internal.Fatalf(format, args...)
	return
}

// Fatalln implements the genolog.Logger interface.
func (j *jsonLogger) Fatalln(args ...interface{}) {
	j.internal.Fatalln(args...)
	return
}

// Fatalw implements the genolog.Logger interface.
func (j *jsonLogger) Fatalw(body string, args ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(args); i++ {
		if i%2 == 0 || i == 0 {
			fields[fmt.Sprintf("%v", args[i])] = nil
			continue
		}
		fields[fmt.Sprintf("%v", args[i-1])] = args[i]
	}

	j.internal.WithFields(fields).Fatal(body)
	return
}

// Raw implements the genolog.Logger interface.
func (j *jsonLogger) Raw(args ...interface{}) {
	j.internal.Print(args...)
	return
}

// Rawf implements the genolog.Logger interface.
func (j *jsonLogger) Rawf(format string, args ...interface{}) {
	j.internal.Printf(format, args...)
	return
}

// Rawln implements the genolog.Logger interface.
func (j *jsonLogger) Rawln(args ...interface{}) {
	j.internal.Println(args...)
	return
}

// SetLogLevel implements the genolog.Logger interface.
func (j *jsonLogger) SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		j.internal.SetLevel(logrus.DebugLevel)
	case "info":
		j.internal.SetLevel(logrus.InfoLevel)
	case "warn":
		j.internal.SetLevel(logrus.WarnLevel)
	case "error":
		j.internal.SetLevel(logrus.ErrorLevel)
	case "fatal":
		j.internal.SetLevel(logrus.FatalLevel)
	}
}
