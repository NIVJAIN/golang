package logger

// import (
// 	"os"

// 	runtime "github.com/banzaicloud/logrus-runtime-formatter"
// 	log "github.com/sirupsen/logrus"
// )

// // LoggerStruct ...
// type LoggerStruct struct {
// 	log *log.Logger
// }

// // LoggerFileInitialization ...
// func (ls *LoggerStruct) LoggerFileInitialization() {
// 	// file, err := os.OpenFile("logfile", os.O_CREATE|os.O_APPEND, 0644)
// 	file, err := os.OpenFile("info.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Log as JSON instead of the default ASCII formatter, but wrapped with the runtime Formatter.
// 	// formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
// 	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{}}

// 	// Enable line number logging as well
// 	formatter.Line = true

// 	// Replace the default Logrus Formatter with our runtime Formatter
// 	log.SetFormatter(&formatter)

// 	// Output to stdout instead of the default stderr
// 	// Can be any io.Writer, see below for File example
// 	// log.SetOutput(os.Stdout)
// 	log.SetOutput(file)

// 	// Only log the info severity or above.
// 	log.SetLevel(log.InfoLevel)
// 	ls.log = log
// 	// return &log
// }
