package ringbuffer_test

import (
	"testing"

	"ringbuffer"
)

func BenchmarkRingBuffer_Add(b *testing.B) {
	rb := ringbuffer.New[Message](3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rb.Add(Message{i})
	}
}

func BenchmarkRingBuffer_Get(b *testing.B) {
	rb := ringbuffer.New[Message](3)
	for i := 0; i < 1024; i++ {
		rb.Add(Message{i})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rb.Get()
	}
}
