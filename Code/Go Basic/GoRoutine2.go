package main
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Application start")

	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutines: ", i)
		}

		wg.Done()
	}()

	fmt.Println("Application end")
	wg.Wait()
	// wg.wait cho đến khi goroutine thực hiện xong => cái này kiểu semaphore. Hàm wg.Add(n) báo cần Wait n lần gọi wg.Done() mới xong
}