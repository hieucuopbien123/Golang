package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	// Các go routine chia sẻ data cho nhau và tự động k bị conflict khi cả 2 cùng thay đổi 1 biến rồi
	// Trong Go vẫn có mutex để ta có thể chủ động thực hiện 1 phần code mà phần khác k can thiệp vào, block kể cả read và write
	// Khi lock thì chỉ cho 1 routine truy cập vào các biến dùng bên dưới. Unlock sẽ trở lại bth, tức là ở trên kp lock là dừng đâu mà lock là
	// check chỉ cho phép 1 routine thực thi bên dưới, nếu routine này là routine thứ 2 sẽ phải chờ

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

	// Từ khóa defer sẽ push hàm hiện tại vào 1 cái queue và chờ khi function gọi nó(routine đó) kết thúc mới pop dần ra thực hiện
	// Có thể ta nghĩ rằng gọi ở cuối thì cũng tương tự nhưng nó còn có thể làm kiểu tự gọi sau hàm return
	defer fmt.Printf("World")
	defer fmt.Printf("One")
	defer fmt.Printf("Two")
	fmt.Printf("Hello")
}
