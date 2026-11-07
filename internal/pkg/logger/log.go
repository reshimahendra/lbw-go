/*
Logger module
- It will create logfile for the server
*/
package logger

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
	goOS   = runtime.GOOS

	// get the root directory of our project
	_, base, _, _ = runtime.Caller(0)
	basePath      = filepath.Join(filepath.Dir(base), "../../..")
)

// init will Initialize the logger setting
func init() {
	// 1. Get the directory where the compiled binary is actually running
	execPath, err := os.Executable()
	if err != nil {
		// Fallback to current working directory if path detection fails
		execPath, _ = os.Getwd()
	}

	// If binary is in project_root/dist/server, EvalSymlinks resolves shortcuts
	realPath, err := filepath.EvalSymlinks(execPath)
	if err == nil {
		execPath = realPath
	}

	// Set basePath to the directory containing the binary
	basePath = filepath.Dir(execPath)

	// 2. Ensure the log folder exists before logrus tries to touch it
	logDir := filepath.Join(basePath, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// 3. Define the file path safely
	logFile := filepath.Join(logDir, ".server.log")

	// 4. Continue with your configuration
	logger.Out = getWriter(logFile)
	logger.Level = logrus.InfoLevel
	logger.Formatter = &formatter{}
	logger.SetReportCaller(true)
	// logDir := filepath.Join(basePath, "log")
	// if err := os.MkdirAll(logDir, 0755); err != nil {
	// 	log.Fatalf("Failed to create log directory: %v", err)
	// }

	// logFile := filepath.Join(basePath, "log", ".server.log")

	// logger.Out = getWriter(logFile)
	// logger.Level = logrus.InfoLevel
	// logger.Formatter = &formatter{}

	// logger.SetReportCaller(true)
}

// SetLogLevel will set the log level mode
func el(level logrus.Level) {
	logger.Level = level
}

type Fields logrus.Fields

// Debugf will logs a message at 'Debug' level
func Debugf(format string, args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Debugf(format, args...)
	}
}

// Infof logs a message at 'Info' level
func Infof(format string, args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Infof(format, args...)
	}
}

// Warnf logs a message at 'Warn' level
func Warnf(format string, args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Warnf(format, args...)
	}
}

// Errorf logs a message at 'Error' level
func Errorf(format string, args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Errorf(format, args...)
	}
}

// Fatalf logs a message at 'Fatal' level
// func Fatalf(format string, args ...interface{}) {
//     if logger.Level >= logrus.FatalLevel {
//         entry := logger.WithFields(logrus.Fields{})
//         entry.Fatalf(format, args...)
//     }
// }

// getWriter will get the logfile as the output of our logger
func getWriter(filepath string) io.Writer {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Errorf("failed to create server log file: %v", err)
		return os.Stdout
	} else {
		return file
	}
}

// Formatter implements logrus.Formatter interface.
type formatter struct {
	prefix string
}

// Format building log message.
func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var sb bytes.Buffer

	var newLine = "\n"
	if goOS == "windows" {
		newLine = "\r\n"
	}

	sb.WriteString(strings.ToUpper(entry.Level.String()))
	sb.WriteString(" ")
	sb.WriteString(entry.Time.Format(time.RFC3339))
	sb.WriteString(" ")
	sb.WriteString(f.prefix)
	sb.WriteString(entry.Message)
	sb.WriteString(newLine)

	return sb.Bytes(), nil
}
