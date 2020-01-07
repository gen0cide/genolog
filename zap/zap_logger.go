package zap

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/lunixbochs/vtclean"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gen0cide/genolog"
)

type zapLogger struct {
	internal     *zap.SugaredLogger
	outputWriter *proxyWriter
}

// NewZapLogger creates a genolog compliant logger that uses Zap logging library.
func NewZapLogger(logcfg zap.Config) genolog.Logger {

	// Build a custom core with a proxyWriter as output writer, allowing us to swap
	// out the output without recreating everything else.
	outputWriter := &proxyWriter{}
	var encoder zapcore.Encoder
	switch logcfg.Encoding {
	case "json":
		encoder = zapcore.NewJSONEncoder(logcfg.EncoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(logcfg.EncoderConfig)
	default:
		encoder = zapcore.NewJSONEncoder(logcfg.EncoderConfig)
	}
	writeSyncer := zapcore.Lock(zapcore.AddSync(outputWriter))
	core := zapcore.NewCore(encoder, writeSyncer, logcfg.Level.Level())
	fields := zap.Fields(makeZapFields(logcfg.InitialFields)...)

	ilogger, err := logcfg.Build()
	if err != nil {
		panic(err)
	}

	ilogger = ilogger.
		WithOptions(
			zap.AddCallerSkip(1),
			zap.WrapCore(func(zapcore.Core) zapcore.Core {
				return core
			}),
			fields,
		)

	// Redirect golang logging output to this logger
	zap.RedirectStdLog(ilogger)

	// Replaces global Logger/SugaredLogger
	zap.ReplaceGlobals(ilogger)

	return &zapLogger{
		internal:     ilogger.Sugar(),
		outputWriter: outputWriter,
	}
}

func makeZapFields(fieldsMap map[string]interface{}) []zap.Field {
	keys := make([]string, 0, len(fieldsMap))
	for k := range fieldsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fs := make([]zap.Field, len(fieldsMap))
	for idx, k := range keys {
		fs[idx] = zap.Any(k, fieldsMap[k])
	}
	return fs
}

// ZapLogger implements the genolog.Logger interface.
func (z *zapLogger) ZapLogger() *zap.SugaredLogger {
	return z.internal
}

// LogrusLogger implements the genolog.Logger interface.
func (z *zapLogger) LogrusLogger() *logrus.Logger {
	return nil
}

// SetName implements the genolog.Logger interface.
func (z *zapLogger) SetName(_ string) {
	return
}

// SetProg implements the genolog.Logger interface.
func (z *zapLogger) SetProg(_ string) {
	return
}

// GetWriter implements the genolog.Logger interface.
func (z *zapLogger) GetWriter() io.Writer {
	return z.outputWriter.GetWriter()
}

// GetOutput implements the genolog.Logger interface.
func (z *zapLogger) GetOutput() io.Writer {
	return z.GetWriter()
}

// SetOutput implements the genolog.Logger interface.
func (z *zapLogger) SetOutput(w io.Writer) {
	z.outputWriter.SetWriter(w)
}

// Print implements the genolog.Logger interface.
func (z *zapLogger) Print(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Info(newArgs...)
	return
}

// Printf implements the genolog.Logger interface.
func (z *zapLogger) Printf(format string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Infof(format, newArgs...)
	return
}

// Println implements the genolog.Logger interface.
func (z *zapLogger) Println(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Info(newArgs...)
	return
}

// Printw implements the genolog.Logger interface.
func (z *zapLogger) Printw(body string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Infow(body, newArgs...)
	return
}

// Debug implements the genolog.Logger interface.
func (z *zapLogger) Debug(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Debug(newArgs...)
	return
}

// Debugf implements the genolog.Logger interface.
func (z *zapLogger) Debugf(format string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Debugf(format, newArgs...)
	return
}

// Debugln implements the genolog.Logger interface.
func (z *zapLogger) Debugln(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Debug(newArgs...)
	return
}

// Debugw implements the genolog.Logger interface.
func (z *zapLogger) Debugw(body string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
			continue
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Debugw(body, newArgs...)
	return
}

// Info wraps implements the genolog.Logger interface.
func (z *zapLogger) Info(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Info(newArgs...)
	return
}

// Infof implements the genolog.Logger interface.
func (z *zapLogger) Infof(format string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Infof(format, newArgs...)
	return
}

// Infoln implements the genolog.Logger interface.
func (z *zapLogger) Infoln(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Info(newArgs...)
	return
}

// Infow implements the genolog.Logger interface.
func (z *zapLogger) Infow(body string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Infow(body, newArgs...)
	return
}

// Warn implements the genolog.Logger interface.
func (z *zapLogger) Warn(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Warn(newArgs...)
	return
}

// Warnf implements the genolog.Logger interface.
func (z *zapLogger) Warnf(format string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Warnf(format, newArgs...)
	return
}

// Warnln implements the genolog.Logger interface.
func (z *zapLogger) Warnln(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Warn(newArgs...)
	return
}

// Warnw implements the genolog.Logger interface.
func (z *zapLogger) Warnw(body string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Warnw(body, newArgs...)
	return
}

// Error implements the genolog.Logger interface.
func (z *zapLogger) Error(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Error(newArgs...)
	return
}

// Errorf implements the genolog.Logger interface.
func (z *zapLogger) Errorf(format string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Errorf(format, newArgs...)
	return
}

// Errorln implements the genolog.Logger interface.
func (z *zapLogger) Errorln(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Error(newArgs...)
	return
}

// Errorw implements the genolog.Logger interface.
func (z *zapLogger) Errorw(body string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Errorw(body, newArgs...)
	return
}

// Fatal implements the genolog.Logger interface.
func (z *zapLogger) Fatal(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Fatal(newArgs...)
	return
}

// Fatalf implements the genolog.Logger interface.
func (z *zapLogger) Fatalf(format string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Fatalf(format, newArgs...)
	return
}

// Fatalln implements the genolog.Logger interface.
func (z *zapLogger) Fatalln(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Fatal(newArgs...)
	return
}

// Fatalw implements the genolog.Logger interface.
func (z *zapLogger) Fatalw(body string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	z.internal.Fatalw(body, newArgs...)
}

// Raw implements the genolog.Logger interface.
func (z *zapLogger) Raw(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	_, _ = fmt.Fprint(os.Stdout, newArgs...)
	return
}

// Rawf implements the genolog.Logger interface.
func (z *zapLogger) Rawf(format string, args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	_, _ = fmt.Fprintf(os.Stdout, format, newArgs...)
	return
}

// Rawln implements the genolog.Logger interface.
func (z *zapLogger) Rawln(args ...interface{}) {
	var newArgs []interface{}
	for _, x := range args {
		str, ok := x.(string)
		if !ok {
			newArgs = append(newArgs, x)
		}
		newArgs = append(newArgs, vtclean.Clean(str, false))
	}
	_, _ = fmt.Fprintln(os.Stdout, newArgs...)
	return
}

// SetLogLevel implements the genolog.Logger interface.
func (z *zapLogger) SetLogLevel(_ string) {
	return
}

type proxyWriter struct {
	writer io.Writer
}

func (w *proxyWriter) Write(p []byte) (int, error) {
	if w.writer == nil {
		return 0, nil
	}
	return w.writer.Write(p)
}

// SetWriter implements the genolog.Logger interface.
func (w *proxyWriter) SetWriter(writer io.Writer) {
	w.writer = writer
}

// GetWriter implements the genolog.Logger interface.
func (w *proxyWriter) GetWriter() io.Writer {
	return w.writer
}
