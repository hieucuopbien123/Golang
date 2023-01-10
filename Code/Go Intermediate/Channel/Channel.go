package main

import (
	"fmt"
	"time"
)

func main() {
	// -> Dùng chuẩn: cho channel chờ bắt data và trong 1 go routine trước đó truyền data đó
	done := make(chan bool)

	fmt.Println("Application start")

	// Tạo 1 routine
	go func() {
		time.Sleep(time.Second)
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutines: ", i)
		}
		done <- true // gửi data đến routine, routine nhận được sẽ unlock
	}()

	fmt.Println("Application end")
	<-done // routine chờ nhận data, block phần code bên dưới

	// -> Tuy nhiên gây lỗi nếu: 
	// Channel chờ nhận data mà k có goroutine nào gửi data đến routine. Khi mọi routine kết thúc mà vẫn chờ sẽ báo lỗi
	// Channel gửi data mà chưa chờ nhận data. Kiểu gọi done<-true ko bất đồng bộ và thực hiện trước <-done là toang
	// TH 1 ok nhưng TH2 éo ok. Để phòng lỗi đó ta có thể dùng Channel Buffering
	done1 := make(chan string, 1)
	done1 <- "Done"
	// Dù done1 gửi data nhưng chưa từng nhận data nhưng k lỗi vì channel buffering tham số 2 là n báo rằng có thể có tới 
	// max là n giá trị mà channel gửi nhưng k cần đầu nhận tương ứng, nếu có đầu nhận thì cũng được.
}