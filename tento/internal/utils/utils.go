package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"tento/internal/utils/color"
)

type tentoLogger struct {
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	level       string
}

func (logger tentoLogger) Info(v ...interface{}) {
	if strings.ToLower(logger.level) == "info" {
		logger.InfoLogger.Println(v...)
	}
}

func (logger tentoLogger) Warn(v ...interface{}) {
	if strings.ToLower(logger.level) != "error" && strings.ToLower(logger.level) != "none" {
		logger.WarnLogger.Println(v...)
	}
}

func (logger tentoLogger) Error(v ...interface{}) {
	if strings.ToLower(logger.level) != "none" {
		logger.ErrorLogger.Println(v...)
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
		log.Ldate|log.Ltime|log.Lshortfile)

	TentoLogger.WarnLogger = log.New(
		os.Stdout,
		color.Yellow+"[tento-WARN]: "+color.Reset,
		log.Ldate|log.Ltime|log.Lshortfile)

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
		TentoLogger.level = "ERROR"
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
