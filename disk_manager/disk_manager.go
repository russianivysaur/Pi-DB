package disk_manager

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"pidb/config"
	"time"
)
import bufferPool "pidb/buffer_pool"
import "log"

const logDirectory string = "disk_logs"

type DiskManager struct {
	databaseFileName string
	databaseFile     *os.File
	logger           *log.Logger
	config           config.Config
}

func NewDiskManager(fileName string, directory string, conf config.Config) *DiskManager {
	logger, err := createLogger(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	dbFile, err := openDatabaseFile(path.Join(directory, fileName), logger)
	if err != nil {
		log.Fatalln(err)
	}
	return &DiskManager{
		databaseFileName: fileName,
		databaseFile:     dbFile,
		logger:           logger,
		config:           conf,
	}
}

func (manager *DiskManager) ReadPageFromDisk(pageNumber int, buf *bufferPool.BufferPoolPage,
	bufDesc *bufferPool.BufferPoolDescriptor) {
	//write lock on the buffer
	bufDesc.Lock.Lock()
	defer bufDesc.Lock.Unlock()
	pageSize := manager.config.PoolConfig.PageSize
	_, err := manager.databaseFile.Seek(int64(pageSize)*int64(pageNumber), 0)
	if err != nil {
		log.Println(err)
		return
	}
	bufReader := bufio.NewReader(manager.databaseFile)
	_, err = io.Copy(buf, bufReader)
	if err != nil {
		log.Println(err)
		return
	}
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
