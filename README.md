# RingBuffer

golang RingBuffer impl

## Installation

To start using RingBuffer, install Go and run go get:

```sh
go get -u github.com/openpkgs/ringbuffer
```

## Usage

Import Copier into your application to access its copying capabilities

```go
import (
	"fmt"
	
	"github.com/openpkgs/ringbuffer"
)

func main() {
	rb := ringbuffer.New[int](1024)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)

	// 取出元素
	if val, ok := rb.Get(); ok {
		fmt.Println("Got:", val) // Output: Got: 1
	}

	// 添加新元素，触发覆盖
	discarded, discardedValid := rb.Add(4)
	if discardedValid {
		fmt.Println("Discarded:", discarded) // Output: Discarded: 2
	}
}
```

