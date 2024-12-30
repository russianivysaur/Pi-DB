package buffer_pool

import "pidb/types"

type BufferPool struct {
	id       int //buffer pool id
	size     int //in bytes
	occupied int
	free     int
	pages    []*BufferPoolPage
}

func NewBufferPool(pageSize types.PageSize, pageCount int) *BufferPool {
	bufferPoolPages := allocatePoolSpace(pageSize, pageCount)
	poolSize := int(pageSize) * pageCount
	return &BufferPool{
		0,
		poolSize,
		0,
		pageCount,
		bufferPoolPages,
	}
}

func allocatePoolSpace(pageSize types.PageSize, pageCount int) []*BufferPoolPage {
	pages := make([]*BufferPoolPage, pageCount)
	for i := 0; i < pageCount; i++ {
		pages[i] = NewBufferPoolPage(pageSize)
	}
	return pages
}

func (pool *BufferPool) getFreePage() *BufferPoolPage {
	for _, page := range pool.pages {
		if page.flags == 0 {
			page.flags = 1
			return page
		}
	}
	return nil
}
