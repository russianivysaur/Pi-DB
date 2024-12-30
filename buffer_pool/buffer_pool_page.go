package buffer_pool

import "pidb/types"

type BufferPoolPage struct {
	flags    uint8 //LSB 0 : pinned LSB 1 : can evict
	contents []byte
}

func NewBufferPoolPage(pageSize types.PageSize) *BufferPoolPage {
	return &BufferPoolPage{
		0,
		make([]byte, pageSize),
	}
}
