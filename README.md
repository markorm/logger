# Logger

A logfile management library for go:

 * Manages and Writes to log files
 * Configurabile lognames and directories
 * Custom logtypes
 * Rotates logs


##### Using Logger


```javascript
import "github.com/markorm/logger"

// some options
options := logger.Options {
	MaxSize: 1024,
	Directory: "/path/to/logs/",
	ErrorLog	"error.log",
	AccessLog	"access.log",
}

// get a logger
logger := logger.NewLogger(options)

// call some methods
err := logger.LogError("my error message")
err := logger.LogAccess("my access message")
err := logger.LogCustom("my custom message", "custom.log")
```
