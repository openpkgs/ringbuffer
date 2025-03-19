package ringbuffer

import "sync"

// RingBuffer 定义环形缓冲区结构体
type RingBuffer[T any] struct {
	buffer []T
	size   int
	head   int
	tail   int
	count  int
	mu     sync.Mutex // 互斥锁，用于保护共享资源
}

// New 创建一个新的环形缓冲区
func New[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer: make([]T, size),
		size:   size,
		head:   0,
		tail:   0,
		count:  0,
	}
}

// Add 添加新元素到环形缓冲区，并返回被丢弃的数据（如果有）
func (rb *RingBuffer[T]) Add(msg T) (discarded T, discardedValid bool) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	// 变量 discardedValid 表示是否有被丢弃的数据
	discardedValid = rb.count == rb.size

	// 如果缓冲区已满，保存被覆盖的元素
	if discardedValid {
		discarded = rb.buffer[rb.head]
		// 移动 head 指针，覆盖最早的元素
		rb.head = (rb.head + 1) % rb.size
	} else {
		// 缓冲区未满，元素数量加 1
		rb.count++
	}

	// 将新元素添加到 tail 位置
	rb.buffer[rb.tail] = msg
	// 移动 tail 指针
	rb.tail = (rb.tail + 1) % rb.size

	// 返回被丢弃的数据（如果有）
	return discarded, discardedValid
}

// Get 获取最早添加的数据并从缓冲区中删除
func (rb *RingBuffer[T]) Get() (msg T, ok bool) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	// 如果缓冲区为空，返回 false
	if rb.count == 0 {
		return msg, false
	}

	// 获取最早添加的元素
	msg = rb.buffer[rb.head]
	// 移动 head 指针
	rb.head = (rb.head + 1) % rb.size
	// 元素数量减 1
	rb.count--

	return msg, true
}

// Clear 清空环形缓冲区
func (rb *RingBuffer[T]) Clear() {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	rb.head = 0
	rb.tail = 0
	rb.count = 0
	// 可选：将缓冲区中的指针置为零值，帮助垃圾回收
	var zero T
	for i := range rb.buffer {
		rb.buffer[i] = zero
	}
}

// IsEmpty 检查环形缓冲区是否为空
func (rb *RingBuffer[T]) IsEmpty() bool {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	// 如果 count 为 0，说明缓冲区为空
	return rb.count == 0
}
