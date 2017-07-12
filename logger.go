package logger

import (
	"time"
	"os"
	"log"
)

type LoggerOptions struct {
	MaxSize		int
	MaxAge		time.Time
	Directory	string
	ErrorLog	string
	AccessLog	string
}

type Logger struct {
	Options LoggerOptions
}

// === Get Logger ===
// Returns a new logger
// @param opts: the logger options
func GetLogger(opts LoggerOptions) *Logger {
	lgr := Logger{}
	lgr.Options = opts
	return &lgr
}

// Rotate logs
func (lgr *Logger) rotateLogs(logname string) {
	//dir, err := os.Open()
}

// Write to a log file
func (lgr *Logger) writeLog(logname, message string) error {
	lgr.rotateLogs(logname)
	f, err := os.OpenFile(lgr.Options.Directory + "/" + logname, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err == nil {
		defer f.Close()
		log.SetOutput(f)
		log.Println(message)
	}
	return err
}

// === Log Error ===
func (lgr *Logger) LogError(message string) error {
	return lgr.writeLog(lgr.Options.ErrorLog, message)
}

// === Log Access ===
func (lgr *Logger) LogAccess(message string) error {
	return lgr.writeLog(lgr.Options.AccessLog, message)
}

// === Log Custom ===
func (lgr *Logger) LogCustom(message, logname string) error {
	return lgr.writeLog(logname, message)
}