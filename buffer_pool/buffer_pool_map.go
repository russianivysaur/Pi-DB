package buffer_pool

type BufferPoolMap struct {
	poolMap map[BufferTag]*BufferPoolDescriptor
}
