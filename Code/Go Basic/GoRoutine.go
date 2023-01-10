package main 

import (
	"fmt" 
	"time"
)

func main(){
	go g1()
	go g2()
	time.Sleep(time.Second) // Chương trình main sleep 1 giây chờ 2 hàm g1(), và g2() kết thúc. Nếu k thì mainthread kết thúc sẽ 
	// dừng luôn tất cả dù g1, g2 chưa thực hiện xong. Nhược điểm là ta kb chờ bao lâu thì dủ để g1, g2 thực hiện xong nên mới phải
	// dùng sync.WaitGroup
}

func g1(){
	for i:=1; i< 10;i++ {
		go fmt.Println(i)
	}
}
func g2(){
	for i:=10;i< 20;i++ {
		go fmt.Println(i)
	}
}