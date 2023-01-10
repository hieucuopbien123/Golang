package main
import ("fmt")

// # Basic
// constant
const (
  A int = 1
  B = 3.14
  C = "Hi!"
)
// # DÙng truct
type Person struct {
  name string
  age int
  job string
  salary int
}

type Student struct {
  Person
  id uint
}

type Rectangle struct {
	X, Y float64
}
// Method with receiver `Rectangle`
func (p Rectangle) Acreage() float64 {
	return p.Y * p.X
}
func (p *Rectangle) Translate(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
  const PI = 3.14

  var pers1 Person
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  fmt.Println("Name: ", pers1) // in cả
  fmt.Println("Age: ", pers1.age) // in từng thuộc tính
  fmt.Println("Job: ", pers1.job)
  fmt.Println("Salary: ", pers1.salary)

  stu := Student{pers1, 10}
  fmt.Println("Stu: ", stu)

  p3 := Person{name: "Robert"} 
  fmt.Println(p3)

  // # Biến và kiểu dữ liệu
  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v\n", y, y)

  var z float64 = 1.7e+308 // K xđ type thì float64 là default
  fmt.Printf("Type: %T, value: %v\n", z, z)

  // # Slices 
  // Khởi tạo slices với make
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
  
  // Nối 2 slices với nhau
  myslice1 := []int{1,2,3}
  myslice3 := []int{4,5,6}
  myslice4 := append(myslice1, myslice3...)
  fmt.Printf("myslice4=%v\n", myslice4)
  fmt.Printf("length=%d\n", len(myslice4))
  fmt.Printf("capacity=%d\n", cap(myslice4))

  // Tạo memory effiency khi dùng slice
  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  myslice := arr1[1:5] // Slice array
  fmt.Printf("myslice1 = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))

  myslice = arr1[1:3] // Change length by re-slicing the array, gán bằng chứ kp khởi tạo
  fmt.Printf("myslice1 = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))
  // Lúc này lenght của slice cũ sẽ là capacity of slice mới 

  myslice = append(myslice, 20, 21, 22, 23) // Change length by appending items, gán bằng chứ kp khởi tạo
  fmt.Printf("myslice1 = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))

  // # Array
  var arr2 = [...]int{1,2,3};
  arr3 := [3]int{1,2,3}

  fmt.Println(arr2)
  fmt.Println(arr3)
  fmt.Println(arr2 == arr3) // Chỉ ss được khi 2 mảng cùng số lương phần tử mà thôi. Nó ss giá trị và thứ tự phần tử

  var h [5]complex128 // dùng có type dữ liệu với mảng thì k được khai báo luôn
  fmt.Println(h);

  // Array 2 chièu
  k := [5][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(k)

  // # Map
  // Duyệt map có thứ tư
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  var b []string
  b = append(b, "one", "two", "three", "four") // Có thể sort nó tùy ý

  for k, v := range a { 
    fmt.Printf("%v : %v, ", k, v)
  }

  fmt.Println()

  for _, element := range b {  // loop with the defined order
    fmt.Printf("%v : %v, ", element, a[element])
  }

  // # Basic / Vòng for 
  i := 2
  for ;i <= 10; i += 2 {
    if i%2 != 0 {
      continue;
    }
    fmt.Printf("%d ", i)
  }
  fmt.Println("")
  
  // # Dùng struct
  // Receiver argument
  receiverArg := Rectangle{2,3}
  fmt.Println(receiverArg.Acreage())
  receiverArg.Translate(2,3)
  fmt.Println(receiverArg)

  
}