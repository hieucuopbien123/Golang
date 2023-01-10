package main

import ( "fmt" )

// # Dùng struct 
// Receiver argument
type MyString string

// K có kiểu ký tự, đổi ngược string ta dùng mảng int của nó r đảo ngược bản r convert ngược lại
func (myStr MyString) reverse() string {
	s := string(myStr)
	// # Slices / Convert từ slice số về string và từ string sang slice số để thao tác 
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) // Convert ngược lại từ slice số về string
}

func main() {
	x := []rune("Hello")
	fmt.Println(x)

	// # Con trỏ
	var p *float64
	fmt.Println("p = ", p)
	var a = 5.67
	p = &a
	fmt.Println("Address of variable a = ", &a)

	fmt.Println("*p = ", *p)
	*p = 2000
	fmt.Println("a (after) = ", a)

	// Con trỏ tới con trỏ
	var pp = &p
	fmt.Println("pp = ", pp)
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)

	// Heap lưu giá trị, stack lưu địa chỉ. Cơ chế copy giá trị y hệt như C++. Khi truyền địa chỉ, nó copy address lưu trong stack và
	// đổi giá trị reference tới heap
	thayDoiTrenGiaTri(a)
	fmt.Println(a)
	thayDoiTrenDiaChi(&a)
	fmt.Println(a)
}
func thayDoiTrenGiaTri(a float64) {
	a = 100
}
func thayDoiTrenDiaChi(a *float64){
	*a = 200
}