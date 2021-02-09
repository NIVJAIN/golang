package logger

// import (
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"runtime"

// 	"github.com/sirupsen/logrus"
// 	"github.com/sirupsen/logrus/hooks/writer"
// 	"gopkg.in/natefinch/lumberjack.v2"
// )

// func LogSetup() {
// 	var (
// 		_, b, _, _ = runtime.Caller(0)
// 		basepath   = filepath.Dir(b)
// 	)
// 	// fmt.Println(basepath)
// 	logOne := logrus.New()
// 	logTwo := logrus.New()
// 	logOne.SetOutput(ioutil.Discard) // Send all logs to nowhere by default
// 	logTwo.SetOutput(ioutil.Discard) // Send all logs to nowhere by default
// 	// home := os.Getenv("HOME")
// 	home := basepath + "/logs"
// 	file1 := home + "/file1.log"
// 	file11 := home + "/file1Info.log"
// 	file2 := home + "/file2.log"
// 	f1, _ := os.Create(file1)
// 	f11, _ := os.Create(file11)
// 	f2, _ := os.Create(file2)
// 	defer f1.Close()
// 	defer f11.Close()
// 	defer f2.Close()

// 	l1 := lumberjack.Logger{
// 		Filename:   file1,
// 		MaxSize:    1, // megabytes
// 		MaxBackups: 3,
// 		MaxAge:     1, // days
// 	}

// 	l11 := lumberjack.Logger{
// 		Filename:   file11,
// 		MaxSize:    1, // megabytes
// 		MaxBackups: 3,
// 		MaxAge:     1, // days
// 	}

// 	l2 := lumberjack.Logger{
// 		Filename:   file2,
// 		MaxSize:    1, // megabytes
// 		MaxBackups: 3,
// 		MaxAge:     1, // days
// 	}

// 	logOne.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
// 		Writer: &l1,
// 		LogLevels: []logrus.Level{
// 			logrus.PanicLevel,
// 			logrus.FatalLevel,
// 			logrus.ErrorLevel,
// 			logrus.WarnLevel,
// 		},
// 	})

// 	logOne.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
// 		Writer: &l11,
// 		LogLevels: []logrus.Level{
// 			logrus.InfoLevel,
// 			logrus.DebugLevel,
// 		},
// 	})

// 	logTwo.AddHook(&writer.Hook{ // Send info and debug logs to stdout
// 		Writer: &l2,
// 		LogLevels: []logrus.Level{
// 			logrus.InfoLevel,
// 			logrus.DebugLevel,
// 		},
// 	})

// 	logEntryOne := logrus.NewEntry(logOne)
// 	logEntryTwo := logrus.NewEntry(logTwo)
// 	logEntryOne = logOne.WithFields(logrus.Fields{
// 		"osplatform": runtime.GOOS,
// 		"osarch":     runtime.GOARCH,
// 		"deviceId":   "kucch bhi nahi h",
// 		"type":       "HookOne",
// 	})

// 	logEntryTwo = logTwo.WithFields(logrus.Fields{
// 		"osplatform": runtime.GOOS,
// 		"osarch":     runtime.GOARCH,
// 		"type":       "HookTwo",
// 	})
// 	for ii := 0; ii < 100000; ii++ {
// 		logEntryOne.Info("This will go to file 1 info ")
// 		logEntryOne.Error("This will go to file 1  ")
// 		logEntryTwo.Info("This will go to file 2 ")
// 	}
// }
