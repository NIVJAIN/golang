package alarm

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sort"
	"time"

	// "ginDemo/common/function"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/config"
)

type errorString struct {
	s string
}

type errorInfo struct {
	Time     string `json:"time"`
	Alarm    string `json:"alarm"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Funcname string `json:"funcname"`
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	alarm("INFO", text, 2)
	return &errorString{text}
}

// 发邮件
func Email(text string) error {
	alarm("EMAIL", text, 2)
	return &errorString{text}
}

// 发短信
func Sms(text string) error {
	alarm("SMS", text, 2)
	return &errorString{text}
}

// 发微信
func WeChat(text string) error {
	alarm("WX", text, 2)
	return &errorString{text}
}

// Panic 异常
func Panic(text string) error {
	alarm("PANIC", text, 5)
	return &errorString{text}
}

// 告警方法
func alarm(level string, str string, skip int) {
	// 当前时间
	currentTime := GetTimeStr()

	// 定义 文件名、行号、方法名
	fileName, line, functionName := "?", 0, "?"

	pc, fileName, line, ok := runtime.Caller(skip)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}

	var msg = errorInfo{
		Time:     currentTime,
		Alarm:    level,
		Message:  str,
		Filename: fileName,
		Line:     line,
		Funcname: functionName,
	}

	// jsons, errs := json.Marshal(msg)
	// jsons, errs := json.Marshal(msg)
	// 	// jsons, errs := json.MarshalIndent(msg, "", "\t🐱")
	// 	// jsons, errs := json.MarshalIndent(msg, "", "🐱")
	// 	// jsons, errs := json.MarshalIndent(msg, "", "\t🚀")
	// jsons, errs := json.MarshalIndent(msg, "", "\t🔥") //🔥
	// 	// jsons, errs := json.MarshalIndent(msg, "", "\t👽")
	jsons, errs := json.MarshalIndent(msg, "", "\t🔥")

	if errs != nil {
		log.Println("json marshal error:", errs)
	}

	errorJSONInfo := string(jsons)

	log.Println(errorJSONInfo)

	if level == "EMAIL" {
		// 执行发邮件

	} else if level == "SMS" {
		// 执行发短信

	} else if level == "WX" {
		// 执行发微信

	} else if level == "INFO" {
		// 执行记日志

	} else if level == "PANIC" {
		// 执行PANIC方式
	}
}

// type errorString struct {
// 	s string
// }

// type errorInfo struct {
// 	Time     string `json:"time"`
// 	Alarm    string `json:"alarm"`
// 	Message  string `json:"message"`
// 	Filename string `json:"filename"`
// 	Line     int    `json:"line"`
// 	Funcname string `json:"funcname"`
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

// func New(text string) error {
// 	alarm("INFO", text)
// 	return &errorString{text}
// }

// // mail
// func Email(text string) error {
// 	alarm("EMAIL", text)
// 	return &errorString{text}
// }

// // send text messages
// func Sms(text string) error {
// 	alarm("SMS", text)
// 	return &errorString{text}
// }

// // micro letter
// func WeChat(text string) error {
// 	alarm("WX", text)
// 	return &errorString{text}
// }

// //Alarm method
// func alarm(level string, str string) {
// 	//Current time
// 	currentTime := GetTimeStr()

// 	//Define file name, line number, method name
// 	fileName, line, functionName := "?", 0, "?"

// 	pc, fileName, line, ok := runtime.Caller(2)
// 	if ok {
// 		functionName = runtime.FuncForPC(pc).Name()
// 		functionName = filepath.Ext(functionName)
// 		functionName = strings.TrimPrefix(functionName, ".")
// 	}
// 	var msg = errorInfo{
// 		Time:     currentTime,
// 		Alarm:    level,
// 		Message:  str,
// 		Filename: fileName,
// 		Line:     line,
// 		Funcname: functionName,
// 	}

// 	// jsons, errs := json.Marshal(msg)
// 	// jsons, errs := json.MarshalIndent(msg, "", "\t🐱")
// 	// jsons, errs := json.MarshalIndent(msg, "", "🐱")
// 	// jsons, errs := json.MarshalIndent(msg, "", "\t🚀")
// 	jsons, errs := json.MarshalIndent(msg, "", "\t🔥") //🔥
// 	// jsons, errs := json.MarshalIndent(msg, "", "\t👽")

// 	if errs != nil {
// 		fmt.Println("json marshal error:", errs)
// 	}

// 	errorJSONInfo := string(jsons)

// 	log.Println(errorJSONInfo)

// 	if level == "EMAIL" {
// 		//Email execution

// 	} else if level == "SMS" {
// 		//Execute SMS

// 	} else if level == "WX" {
// 		//Execute wechat

// 	} else if level == "INFO" {
// 		//Execution logging
// 	}
// }

func GetTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 生成签名
func CreateSign(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" && k != "ts" && k != "debug" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}

	// 自定义签名算法
	sign := MD5(MD5(str) + MD5(config.APP_NAME+config.APP_SECRET))
	return sign
}
