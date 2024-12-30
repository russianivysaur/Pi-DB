package disk_manager

import (
	"bufio"
	"context"
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
	databaseFileName  string
	databaseFile      *os.File
	bufferPoolManager *bufferPool.BufferPoolManager
	logger            *log.Logger
	config            config.Config
}

func NewDiskManager(fileName string, directory string, appContext context.Context) *DiskManager {
	logger, err := createLogger(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	dbFile, err := openDatabaseFile(path.Join(directory, fileName), logger)
	if err != nil {
		log.Fatalln(err)
	}
	bufferPoolManager := createBufferPoolManager(appContext)
	return &DiskManager{
		databaseFileName:  fileName,
		databaseFile:      dbFile,
		bufferPoolManager: bufferPoolManager,
		logger:            logger,
		config:            appContext.Value("config").(config.Config),
	}
}

func (manager *DiskManager) ReadPageFromDisk(pageNumber int) *bufferPool.BufferPoolPage {
	pageSize := manager.config.PoolConfig.PageSize
	_, err := manager.databaseFile.Seek(int64(pageSize)*int64(pageNumber), 0)
	if err != nil {
		log.Println(err)
		return nil
	}
	bufReader := bufio.NewReader(manager.databaseFile)
	freePage := manager.bufferPoolManager.FindFreePage()
	_, err = io.Copy(freePage, bufReader)
	if err != nil {
		log.Println(err)
		return nil
	}
	return freePage
}

func createBufferPoolManager(appContext context.Context) *bufferPool.BufferPoolManager {
	return bufferPool.NewBufferPoolManager(appContext)
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
