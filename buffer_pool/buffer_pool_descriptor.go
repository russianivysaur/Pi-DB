package buffer_pool

import (
	"sync"
	"sync/atomic"
)

type BufferTag struct {
}

type BufferPoolDescriptor struct {
	state    atomic.Uint32 //tag state, containing flags,refcount,usagecount
	bufId    int
	freeNext int
	rwLock   sync.RWMutex
	tag      BufferTag
}
