package service

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
)

var errLog, infoLog *os.File

var infoLogPath = os.Getenv("GOPATH") + "/src/agenda/data/info.log"
var errLogPath = os.Getenv("GOPATH") + "/src/agenda/data/error.log"

var infoWriter, errWriter []io.Writer

var fileAndStdoutWriter1 io.Writer
var fileAndStdoutWriter2 io.Writer

func init() {
	infoLog = getLogFile(infoLogPath)
	errLog = getLogFile(errLogPath)
	errWriter = []io.Writer{
		errLog,
		os.Stdout,
	}
	infoWriter = []io.Writer{
		os.Stdout,
		infoLog,
	}
	fileAndStdoutWriter1 = io.MultiWriter(infoWriter...)
	fileAndStdoutWriter2 = io.MultiWriter(errWriter...)
	Info =  log.New(fileAndStdoutWriter1, "Info: ", log.Ldate | log.Ltime)
	Error = log.New(fileAndStdoutWriter2, "Error: ", log.Ldate | log.Ltime)
}

func getLogFile(logPath string) *os.File  {
	//判断文件是否存在，不存在则创建目录
	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		os.Mkdir(os.Getenv("GOPATH") + "/src/agenda/data", 0777)
	}
	
	//打开文件（文件不存在则会创建）
	LogFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("File open error : %v\n", err)
	}
	return LogFile;
}