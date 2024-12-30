package buffer_pool

import (
	"context"
	"log"
)
import "pidb/config"

type BufferPoolManager struct {
	pools      []*BufferPool
	appContext context.Context
}

func NewBufferPoolManager(appContext context.Context) *BufferPoolManager {
	appConfig := appContext.Value("config").(config.Config)
	pageSize := appConfig.PoolConfig.PageSize
	pageCount := appConfig.PoolConfig.PageCount
	poolCount := appConfig.PoolConfig.PoolCount
	pools := make([]*BufferPool, poolCount)
	for i, _ := range pools {
		pools[i] = NewBufferPool(pageSize, pageCount)
	}
	return &BufferPoolManager{
		pools: pools,
	}
}

func (manager *BufferPoolManager) FindFreePage() *BufferPoolPage {
	var selectedPool *BufferPool
	for _, pool := range manager.pools {
		if pool.free > 1 {
			selectedPool = pool
			break
		}
	}
	if selectedPool == nil {
		log.Println("No free page in any buffer pool")
		return nil
	}
	selectedPage := selectedPool.getFreePage()
	if selectedPage == nil {
		log.Println("Could not allocate page")
		return nil
	}
	return selectedPage
}
