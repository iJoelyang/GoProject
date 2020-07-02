package main

import "fmt"

func updateSlice(s []int)  {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]: ", arr[2:6])
	fmt.Println("arr[1:4]: ", arr[1:4])
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("update s1, s2")

	updateSlice(s1)
	updateSlice(s2)

	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)

	//slice 扩展
	fmt.Println()
	s1 = arr[2:6]
	fmt.Println("s1: ", s1)
	s2 = s1[3:5]
	fmt.Println("s2: ", s2)

	fmt.Println("cap s1:", cap(s1))
	fmt.Println("cap s2:", cap(s2))

	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)

	fmt.Println("s3, s4, s5 =", s3, s4, s5)

	fmt.Printf("&arr=%p\n &s2=%p\n &s3=%p\n &s4=%p", &arr, &s2, &s3, &s4)



}
