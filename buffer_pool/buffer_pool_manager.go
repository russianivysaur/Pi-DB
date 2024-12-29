package buffer_pool

type BufferPoolManager struct {
	pool *BufferPool
}

func NewBufferPoolManager() *BufferPoolManager {
	return &BufferPoolManager{}
}
