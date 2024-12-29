package disk_manager

import (
	"fmt"
	"os"
	"path"
	"time"
)
import bufferPool "pidb/buffer_pool"
import "log"

const logDirectory string = "disk_logs"

type DiskManager struct {
	databaseFileName  string
	databaseFile      *os.File
	bufferPoolManager *bufferPool.BufferPoolManager
	logger            *log.Logger
}

func NewDiskManager(fileName string, directory string) *DiskManager {
	logger, err := createLogger(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	dbFile, err := openDatabaseFile(path.Join(directory, fileName), logger)
	if err != nil {
		log.Fatalln(err)
	}
	bufferPoolManager := createBufferPoolManager()
	return &DiskManager{
		databaseFileName:  fileName,
		databaseFile:      dbFile,
		bufferPoolManager: bufferPoolManager,
		logger:            logger,
	}
}

func createBufferPoolManager() *bufferPool.BufferPoolManager {
	return bufferPool.NewBufferPoolManager()
}

func createLogger(fileName string) (*log.Logger, error) {
	currTime, err := time.Now().MarshalText()
	if err != nil {
		return nil, err
	}
	//create log file
	logFileName := fmt.Sprintf("%s %s", fileName, string(currTime))
	logFilePath := path.Join(logDirectory, logFileName)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	logger := log.Default()
	logger.SetOutput(logFile)
	logger.SetPrefix("Disk Log : ")
	return logger, nil
}

func openDatabaseFile(path string, diskLogger *log.Logger) (*os.File, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(path, os.O_SYNC|os.O_RDWR, 0666)
	return file, err
}
