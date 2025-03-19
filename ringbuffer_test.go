package ringbuffer_test

import (
	"testing"

	"ringbuffer"

	"github.com/stretchr/testify/assert"
)

type Message struct{ Id int }

func TestRingBuffer(t *testing.T) {
	tests := []struct {
		name       string
		size       int
		operations func(rb *ringbuffer.RingBuffer[Message]) Message
		expectId   int
	}{
		{
			"Add 4 messages, expect Id 2",
			3,
			func(rb *ringbuffer.RingBuffer[Message]) Message {
				rb.Add(Message{Id: 1})
				rb.Add(Message{Id: 2})
				rb.Add(Message{Id: 3})
				rb.Add(Message{Id: 4})
				msg, _ := rb.Get()
				return msg
			},
			2,
		},
		{
			"Add one more, expect Id 3",
			3,
			func(rb *ringbuffer.RingBuffer[Message]) Message {
				rb.Add(Message{Id: 1})
				rb.Add(Message{Id: 2})
				rb.Add(Message{Id: 3})
				rb.Add(Message{Id: 4})
				rb.Get()
				rb.Add(Message{Id: 5})
				msg, _ := rb.Get()
				return msg
			},
			3,
		},
		{
			"Add multiple messages, expect Id 6",
			3,
			func(rb *ringbuffer.RingBuffer[Message]) Message {
				rb.Add(Message{Id: 1})
				rb.Add(Message{Id: 2})
				rb.Add(Message{Id: 3})
				rb.Add(Message{Id: 4})
				rb.Get()
				rb.Add(Message{Id: 5})
				rb.Get()
				rb.Add(Message{Id: 6})
				rb.Add(Message{Id: 7})
				rb.Add(Message{Id: 8})
				msg, _ := rb.Get()
				return msg
			},
			6,
		},
		{
			"Add two more, expect Id 8",
			3,
			func(rb *ringbuffer.RingBuffer[Message]) Message {
				rb.Add(Message{Id: 1})
				rb.Add(Message{Id: 2})
				rb.Add(Message{Id: 3})
				rb.Add(Message{Id: 4})
				rb.Get()
				rb.Add(Message{Id: 5})
				rb.Get()
				rb.Add(Message{Id: 6})
				rb.Add(Message{Id: 7})
				rb.Add(Message{Id: 8})
				rb.Get()
				rb.Add(Message{Id: 9})
				rb.Add(Message{Id: 10})
				msg, _ := rb.Get()
				return msg
			},
			8,
		},
		{
			"Clear and add new messages, expect Id 11",
			3,
			func(rb *ringbuffer.RingBuffer[Message]) Message {
				rb.Clear()
				rb.Add(Message{Id: 11})
				rb.Add(Message{Id: 12})
				msg, _ := rb.Get()
				return msg
			},
			11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rb := ringbuffer.New[Message](tt.size)
			msg := tt.operations(rb)
			assert.Equal(t, tt.expectId, msg.Id)
		})
	}
}
