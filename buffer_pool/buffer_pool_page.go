package buffer_pool

import "pidb/types"

type BufferPoolPage []byte

func NewBufferPoolPage(pageSize types.PageSize) *BufferPoolPage {
	bufferPoolPage := BufferPoolPage(make([]byte, pageSize))
	return &bufferPoolPage
}

// implements writer interface
func (page *BufferPoolPage) Write(data []byte) (int, error) {
	for i, b := range data {
		(*page)[i] = b
	}
	return len(data), nil
}
