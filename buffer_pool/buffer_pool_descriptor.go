package buffer_pool

import (
	"sync"
	"sync/atomic"
)

type BufferTag int

type BufferPoolDescriptor struct {
	state    atomic.Uint32 //tag state, containing flags,refcount,usagecount
	bufId    int
	freeNext int
	Lock     sync.RWMutex
	tag      BufferTag
}

func NewBufferPoolDescriptor(id int, next int) *BufferPoolDescriptor {
	return &BufferPoolDescriptor{
		bufId:    id,
		freeNext: next,
		tag:      -1,
	}
}
