package buffer_pool

import "context"
import "pidb/config"

type BufferPoolManager struct {
	pool       *BufferPool
	appContext context.Context
}

func NewBufferPoolManager(appContext context.Context) *BufferPoolManager {
	appConfig := appContext.Value("config").(config.Config)
	pageSize := appConfig.PoolConfig.PageSize
	pageCount := appConfig.PoolConfig.PageCount
	return &BufferPoolManager{
		pool: NewBufferPool(pageSize, pageCount),
	}
}
