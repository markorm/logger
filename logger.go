package logger

import (
	"os"
	"log"
	"time"
)

type LoggerOptions struct {
	MaxSize		int64
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
func NewLogger(opts LoggerOptions) *Logger {
	lgr := Logger{}
	lgr.Options = opts
	return &lgr
}

// Write Log
// Takes a file pointer and returns a file pointer
// Returns the same pointer if the file iowriter did not need to be rotated
// Returns a pointer to the new file iowriter if a new logfile was created
// @arg logname: 	path to log file from application root
// @arg message:
func (lgr *Logger) WriteLog(logname, message string) error {

	// Open the log file
	fpath := lgr.Options.Directory + logname
	f, err := os.OpenFile(fpath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil { return err }

	// Get File Stats
	fstat, err := f.Stat()
	if err != nil { return err }

	// Check if size excedes the size limit
	// If it does update f value to be a pointer to the new file
	if fstat.Size() < lgr.Options.MaxSize  {
		newpath := lgr.Options.Directory + time.Now().String() + "__" + logname + ".bak"
		os.Rename(fpath, newpath)
		f, err = os.OpenFile(fpath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil { return err }
	}

	// If no errors continue to write f
	defer f.Close()
	log.SetOutput(f)
	log.Println(message)

	return err
}

// === Log Error ===
func (lgr *Logger) LogError(message string) error {
	return lgr.WriteLog(lgr.Options.ErrorLog, message)
}

// === Log Access ===
func (lgr *Logger) LogAccess(message string) error {
	return lgr.WriteLog(lgr.Options.AccessLog, message)
}

// === Log Custom ===
func (lgr *Logger) LogCustom(message, logname string) error {
	return lgr.WriteLog(logname, message)
}