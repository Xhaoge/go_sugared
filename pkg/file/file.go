package file

import (
	"fmt"
	"go_sugared/config"
	"log"
	"os"
	"time"
)

var (
	LogSaveName = "go_sugared"
	LogFileExt  = "log"
	TimeFormat  = "202003151653"
)

func getLogFilePath() string {
	logfile := config.ConfigGetLoggingPath()
	return fmt.Sprintf("%s", logfile)
}

func GetLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func OpenLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		MkDir(config.ConfigGetLoggingPath())
	case os.IsPermission(err):
		log.Fatalf("Perssion: %v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to openfile :%v", err)
	}
	return handle
}

func MkDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
