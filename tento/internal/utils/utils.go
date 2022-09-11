package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"tento/internal/utils/color"
)

type tentoLogger struct {
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	level       string
}

func shortenFnPath(fn string) string {
	split := strings.Split(fn, "/")
	return strings.Join(split[len(split)-2:], "/")
}

func getCallLocation(v ...interface{}) string {
	pc, fn, line, _ := runtime.Caller(2)
	funcName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	return fmt.Sprintf("%s() [%s:%d]", funcName, shortenFnPath(fn), line)
}

func (logger tentoLogger) Info(v ...interface{}) {
	if strings.ToLower(logger.level) == "info" {
		prefix := getCallLocation()
		a := append([]interface{}{prefix}, v...)
		logger.InfoLogger.Println(a...)
	}
}

func (logger tentoLogger) Warn(v ...interface{}) {
	if strings.ToLower(logger.level) != "error" && strings.ToLower(logger.level) != "none" {
		prefix := getCallLocation()
		a := append([]interface{}{prefix}, v...)
		logger.WarnLogger.Println(a...)
	}
}

func (logger tentoLogger) Error(v ...interface{}) {
	if strings.ToLower(logger.level) != "none" {
		prefix := getCallLocation()
		a := append([]interface{}{prefix}, v...)
		logger.ErrorLogger.Println(a...)
	}
}

func (logger tentoLogger) SetLogLevel(level string) {
	set := false
	switch strings.ToLower(level) {
	case "info":
		logger.level = "info"
		set = true
		break
	case "warn":
		logger.level = "warn"
		set = true
		break
	case "error":
		logger.level = "error"
		set = true
		break
	case "none":
		logger.level = "none"
		set = true
		break
	}
	if set {
		logger.Info("Setting log level to ", level)
	} else {
		logger.Error(
			fmt.Sprintf("Invalid option: '%v'. Expected one of ('INFO', 'WARN', 'ERROR', 'NONE')\n",
				level))
	}
}

var TentoLogger tentoLogger

func SetupLogger() {
	TentoLogger.InfoLogger = log.New(
		os.Stdout,
		color.Cyan+"[tento-INFO]: "+color.Reset,
		log.Ldate|log.Ltime)

	TentoLogger.WarnLogger = log.New(
		os.Stdout,
		color.Yellow+"[tento-WARN]: "+color.Reset,
		log.Ldate|log.Ltime)

	TentoLogger.ErrorLogger = log.New(
		os.Stdout,
		color.Red+"[tento-ERROR]: "+color.Reset,
		log.Ldate|log.Ltime|log.Lshortfile)

	if os.Getenv("GO_ENV") == "development" {
		TentoLogger.level = Getenv("TENT_LOG_LEVEL", "INFO")
		TentoLogger.Info("Ready")
		TentoLogger.Warn("Ready")
		TentoLogger.Error("Ready")
	} else {
		TentoLogger.level = "INFO"
	}
}

func PrettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		fmt.Println("PrettyPrint Error!!! ", err.Error())
		return ""
	}
	return string(s)
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func ParseIntDefault(number string, defaulValue int) int {
	val, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		TentoLogger.Warn(
			fmt.Sprintf("ParseIntDefault value : '%v' raised error\n%v",
				number,
				err.Error()))
		return defaulValue
	}
	return int(val)
}
