package logger

import (
	"errors"
	"io"
	"os"
	"time"

	// rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	filename "github.com/keepeye/logrus-filename"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// var ErrorLogger, AccessLogger, WorkLogger, MONGOLogger *logrus.Logger

type Logams struct {
	// logcollections map[string]*logrus.Logger
	logcollections *logrus.Logger
}

var (
	LogPools = make(map[string]*Logams)
	Log      *logrus.Logger
)

//GetConnection returns the connection which was instantiated
func GetLogConnection(name string) (*Logams, error) {
	if c, ok := LogPools[name]; ok {
		return c, nil
	}
	err := errors.New("NoPools")
	return nil, err
}

func (l *Logams) SetLogger(outpath string) *logrus.Logger {
	if c, ok := LogPools[outpath]; ok {
		return c.logcollections
	}
	filenameHook := filename.NewHook()
	outpath = "datadog/" + outpath
	logger := logrus.New()
	filenameHook.Field = "line"
	logger.AddHook(filenameHook)
	// logger.SetReportCaller(true)
	formatter := runtime.Formatter{ChildFormatter: &logrus.JSONFormatter{}}
	// formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{}}

	// // Enable line number logging as well
	// formatter.Line = true
	logger.SetFormatter(&formatter)
	// // l.logcollections["info"] = logger

	writer, _ := rotatelogs.New(
		outpath+".%Y%m%d%H%M"+".json",
		// outpath+".%d-%m-%Y.%H"+".log",
		// fmt.Println("adasdfadsfsd", t.Format("20060102150405"))
		// fmt.Println("adasdfadsfsd", t.Format("02-Jan-2006-15h.04m.05s.000"))
		// outpath+"."+time.Now().Format("02-Jan-2006")+".log",
		// outpath+"."+time.Now().Format("02-Jan-2006:15")+".log",
		rotatelogs.WithLinkName(outpath+".json"),
		// rotatelogs.WithMaxAge(time.Duration(2)*time.Hour),
		// rotatelogs.WithRotationTime(time.Duration(1)*time.Hour),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	logger.SetOutput(io.MultiWriter(writer, os.Stdout))
	Log = logger
	return logger
	// logams.logcollections["info"] = logger
}

func init() {
	// WorkLogger = Logger("logs/work.log")
	// // AccessLogger = Logger("logs/access.log")
	// ErrorLogger = Logger("logs/error.log")
	// MONGOLogger = Logger("logs/mongo.log")
	var logams = new(Logams)
	logsetup := logams.SetLogger("info")
	Log = logsetup
}

// //NewConnection returns the new connection object
// func NewConnection(outpath string) *Logams {
// 	if c, ok := LogPools[outpath]; ok {
// 		return c
// 	}
// 	outpath = "logs/" + outpath + ".log"
// 	logger := logrus.New()
// 	_, err := os.Stat(outpath)
// 	if os.IsNotExist(err) {
// 		// 文件不存在,创建
// 		os.Create(outpath)
// 	}
// 	// file, err := os.OpenFile(outpath, os.O_CREATE|os.O_WRONLY, 0666)
// 	file, err := os.OpenFile(outpath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

// 	if err == nil {
// 		logger.Out = file
// 	} else {
// 		logger.Info("打开 " + outpath + " 下的日志文件失败, 使用默认方式显示日志！")
// 	}
// 	// formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{}}
// 	formatter := runtime.Formatter{ChildFormatter: &logrus.JSONFormatter{}}

// 	// Enable line number logging as well
// 	formatter.Line = true
// 	logger.SetFormatter(&formatter)
// 	l.logcollections["info"] = logger
// 	return logger
// }

// func (logams *LOGAMS) Logger(outpath string) *logrus.Logger {
// 	logger := logrus.New()
// 	_, err := os.Stat(outpath)
// 	if os.IsNotExist(err) {
// 		// 文件不存在,创建
// 		os.Create(outpath)
// 	}
// 	file, err := os.OpenFile(outpath, os.O_CREATE|os.O_WRONLY, 0666)
// 	if err == nil {
// 		logger.Out = file
// 	} else {
// 		logger.Info("打开 " + outpath + " 下的日志文件失败, 使用默认方式显示日志！")
// 	}
// 	return logger
// }

// func (l *Logams) SetLogger(outpath string) *logrus.Logger {
// 	filenameHook := filename.NewHook()
// 	outpath = "datadog/" + outpath
// 	logger := logrus.New()
// 	filenameHook.Field = "line"
// 	logger.AddHook(filenameHook)

// 	// _, err := os.Stat(outpath)
// 	// if os.IsNotExist(err) {
// 	// 	// 文件不存在,创建
// 	// 	os.Create(outpath)
// 	// }
// 	// // file, err := os.OpenFile(outpath, os.O_CREATE|os.O_WRONLY, 0666)
// 	// file, err := os.OpenFile(outpath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

// 	// if err == nil {
// 	// 	logger.Out = io.MultiWriter(file, os.Stdout)
// 	// } else {
// 	// 	logger.Info("打开 " + outpath + " 下的日志文件失败, 使用默认方式显示日志！")
// 	// }
// 	// formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{}}
// 	formatter := runtime.Formatter{ChildFormatter: &logrus.JSONFormatter{}}
// 	// Enable line number logging as well
// 	formatter.Line = true
// 	logger.SetFormatter(&formatter)
// 	// l.logcollections["info"] = logger

// 	writer, _ := rotatelogs.New(
// 		// outpath+".%Y%m%d%H%M"+".log",
// fmt.Println("adasdfadsfsd", t.Format("20060102150405"))
// fmt.Println("adasdfadsfsd", t.Format("02-Jan-2006-15h.04m.05s.000"))
// 		// outpath+".%d-%m-%Y.%H"+".log",
// 		// outpath+"."+time.Now().Format("02-Jan-2006")+".log",
// 		outpath+"."+time.Now().Format("02-Jan-2006-15")+".log",
// 		rotatelogs.WithLinkName(outpath),
// 		rotatelogs.WithMaxAge(time.Duration(2)*time.Hour),
// 		rotatelogs.WithRotationTime(time.Duration(1)*time.Hour),
// 		// rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
// 		// rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
// 	)
// 	logger.SetOutput(writer)

// 	return logger
// 	// logams.logcollections["info"] = logger
// }
