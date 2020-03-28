package logging

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
	logfile := config.ConfigGetLoggingFilePath()
	return fmt.Sprintf("%s", logfile)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mlDir()
	case os.IsPermission(err):
		log.Fatalf("Perssion: %v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to openfile :%v", err)
	}
	return handle
}

func mlDir() {
	//dir, _ := os.Getwd()
	dd := config.ConfigGetLoggingFilePath()
	err := os.MkdirAll(dd, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
