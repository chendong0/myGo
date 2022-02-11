package main

import "fmt"

func main() {
	a := []int{} //a是切片
	fmt.Println("添加元素到a切片")
	a = append(a, 123, 22)
	fmt.Println(a)

	b := [0]int{} // b是数组
	fmt.Println(b)
	fmt.Println("添加元素到数组b")
	//b = append(00,01)
	//first argument to append must be slice; have untyped int
	//fmt.Println(b )

	xqInfo := []string{"小强", "男", "在职"}
	fmt.Println(xqInfo)
	for i, v := range xqInfo {
		fmt.Println(i, v)
	}
	fmt.Println(xqInfo[0], xqInfo[1])

	fmt.Println("----->")
	fmt.Println("1删除切片中的元素")
	a = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("2删除之前:", a)

	a = append(a[1:])
	fmt.Println("3删除第一个元素", a)
	a = append(a[:1])
	fmt.Println("4只留index为1的元素", a)
	fmt.Println(a) //切片是引用类型,变更后的切片下传.
	g := []int{11, 12, 13}
	g = append(g[1:1]) //g:= arr[stratIndex:endIndex]
	fmt.Println("5在arr中下标startIndex到endIndex-1下的元素创建一个新的切片.", g)

	a = append(a[:1], a[2:4]...)
	fmt.Println("6删除后: ", a)
	a = append(a, a[:]...) //利用切片截取和append()函数实现切片删除元素

	fmt.Println("7double后:", a) //切片中添加切片

	backup := append([]int{}, a[1:]...)
	a = append(a[:1], 999)
	a = append(a, backup...)
	fmt.Println(a)
}
