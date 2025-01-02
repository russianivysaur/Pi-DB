package buffer_pool

import (
	"pidb/config"
	"sync"
)

type BufferPool struct {
	poolMap     map[BufferTag]*BufferPoolDescriptor
	pages       []*BufferPoolPage
	descriptors []*BufferPoolDescriptor
	lock        sync.Mutex
}

func NewBufferPool(conf config.Config) *BufferPool {
	pageCount := conf.PoolConfig.PageCount
	pageSize := conf.PoolConfig.PageSize

	//pool pages and descriptors
	pages := make([]*BufferPoolPage, pageCount)
	descriptors := make([]*BufferPoolDescriptor, pageCount)
	for index, _ := range pages {
		pages[index] = NewBufferPoolPage(pageSize)
		descriptors[index] = NewBufferPoolDescriptor(index, index+1)
	}
	return &BufferPool{
		pages:       pages,
		descriptors: descriptors,
		poolMap:     make(map[BufferTag]*BufferPoolDescriptor),
	}
}

func (pool *BufferPool) GetPage(objectId int) (*BufferPoolPage, *BufferPoolDescriptor) {
	tag := BufferTag(objectId)
	if desc, found := pool.poolMap[tag]; found {
		return pool.pages[desc.bufId], desc
	}
	desc := pool.allocateNewPage(tag)
	return pool.pages[desc.bufId], desc
}

func (pool *BufferPool) allocateNewPage(tag BufferTag) *BufferPoolDescriptor {
	//read start
	//locking the bufferpool as a whole
	pool.lock.Lock()
	defer pool.lock.Lock()
	//finding free descriptor
	desc := pool.findFreeDescriptor()
	desc.tag = tag
	pool.poolMap[tag] = desc
	return desc
}

func (pool *BufferPool) findFreeDescriptor() *BufferPoolDescriptor {
	for _, desc := range pool.descriptors {
		if desc.tag == -1 {
			return desc
		}
	}
	return nil
}
