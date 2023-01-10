package main

import (  
    "fmt"
)

func calcSquares(number int, squareop chan int) {  
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    squareop <- sum
}

func calcCubes(number int, cubeop chan int) {  
    sum := 0 
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    cubeop <- sum
} 

func producer(chnl chan int) {  
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
	// -> Quy tắc nữa: khi ta dùng for để lấy data từ 1 channel, nó sẽ lấy toàn bộ data trong nó và kích hoạt chế độ bình luôn mở để luôn bắt 
    // multiple data đi thêm vào trong quá trình(mà ta phải tự đóng). Trong quá trình bắt data sẽ block cái routine hiện tại như là ta gọi <-pipe
    // bth(nhưng ở đây điểm khác là nó bắt multiple value truyền vào pipe). Mà như cơ chế đã biết là gọi <-pipe thì phải đảm bảo pipe đang bắt data 
    // thì phải có thêm data. Vì nếu chỉ bắt data mà k truyền thì sẽ lỗi. Buộc phải close channel thủ công để thông báo channel k nhận thêm dữ liệu
    // nào nữa mới dừng đk chờ tự động là <-pipe
}

func main() {  
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech	// -> Có thể gán giá trị nhận về vào biến
    fmt.Println("Final output", squares + cubes)

	ch := make(chan int)
    go producer(ch)
    for v := range ch { // Mở van và chờ và đồng thời lấy multiple data bên trong
        fmt.Println("Received ",v)
    }

	pipe := make(chan string, 2)
	go func() {
		pipe <- "water"
		pipe <- "water"
		close(pipe) // đóng van mới chạy tiếp bên dưới
	}()
	for receiver := range pipe {
		fmt.Println(receiver)
	}

	// Khi có buffer channel thì dùng đồng bộ hay bất đồng bộ đều được miễn k quá buffer
	pipe := make(chan string, 2)
	pipe <- "water"
	pipe <- "water"
	close(pipe)  // Dừng bơm, khoá van
	for receiver := range pipe {
		fmt.Println(receiver)
	}
}