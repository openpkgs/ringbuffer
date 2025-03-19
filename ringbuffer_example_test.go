package ringbuffer_test

import (
	"fmt"

	"ringbuffer"
)

func ExampleRingBuffer() {
	rb := ringbuffer.New[Message](3)

	rb.Add(Message{1})
	rb.Add(Message{2})
	rb.Add(Message{3})

	val, _ := rb.Get()
	fmt.Println(val.Id)

	rb.Add(Message{4})
	discarded, _ := rb.Add(Message{5})
	fmt.Println(discarded.Id)

	// Output:
	// 1
	// 2
}
