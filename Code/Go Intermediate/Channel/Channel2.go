package main

import (
	"fmt"
	"time"
)

func sendValue(number string, channel chan<- string) {
	for { // vòng for push data vào channel sẽ tự động thực hiện cho đến khi channel đầy 
		channel <- number
	}
}

func receiveValue(channel <-chan string) {
	// -> Tháo bình và lấy toàn bộ data
	for v := range channel {
		fmt.Println(v)
	}
}

func main() {
	var channel chan string // -> Tạo biến nhưng k dùng được
    if channel == nil {
        fmt.Println("channel a is nil, going to define it")
        channel = make(chan string, 2)
        fmt.Printf("Type of a is %T", channel)
    }

	go sendValue("Hello", channel)
	go sendValue("Xin chao", channel)

	go receiveValue(channel)
	time.Sleep(time.Second)
	// Chuyện gì xảy ra ở bên trên. Thực tế ta gửi vào channel Hello và Xin chao với max buffer là 2, nhưng khi vào xong nó lại pop ra ngay 
	// vì hàm receiveValue cùng thực hiện lúc đó là bất đồng bộ. Mà vừa vào đã ra dẫn đến channel k thể đầy và cứ thế tiếp tục mãi do vòng 
	// for ta k set đk kết thúc => kbh dùng như v
	// Có lẽ để bất đồng bộ ở cái for lấy data khiến cho nó làm nhanh or chậm hơn mà có thể thoát khỏi vòng for vô tận chứ để bth nó éo thoát
	// đc vòng
	// Giải thích: Ta đã biết vòng for + channel sẽ khiến channel chờ và bắt data liên tục, phải chủ động đóng channel mới được. Nhưng ở đây
	// ta k cần đóng channel vì nó nằm trong 1 go routine khác. Nên khi nó thực tế vẫn block cái routine đó k thực hiện bên dưới nhưng hàm main 
	// vẫn chạy tiếp bth k sao. Nếu viết ở hàm main và nó block hàm main mà k có ai close van lại thì sẽ deadlock ngay

	// Khi dùng buffer có thể k cần bất đồng bộ nữa, thực hiện push pop bất cứ lúc nào miễn đủ buffer và pop đủ lượng push. Do k dùng vòng
	// for nên k cần close
    pipe := make(chan string, 2)
	pipe <- "water"
    pipe <- "water"
    receiver := <-pipe
    fmt.Println(receiver)
    receiver = <-pipe
    fmt.Println(receiver)
}