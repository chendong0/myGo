package main

import "fmt"

func main() {
	//练习指针
	a := 2
	b := &a
	c := *b
	fmt.Printf("a的内存地址为:%p\n", &a)
	fmt.Printf("b变量的值为%v\n", b)
	fmt.Printf("指针b的值为%v\n", *b) //指针是地址所指向的值
	fmt.Printf("指针c的值为%p\n", c)
	fmt.Printf("%v\n", c)

	fmt.Println("1st 测试testing")
	//难以长期维护
	xqInfo := [3]string{"小强", "男", "在职"}
	xlInfo := [3]string{"小李", "男", "在职"}
	xsInfo := [3]string{"小苏", "男", "兼职"}
	fmt.Println(xqInfo, xlInfo, xsInfo)

	newPersonInfos2 := []string{"99"}

	for _, val := range newPersonInfos2 {
		fmt.Println(val)
	}
	//newPersonInfos2 = append(newPersonInfos2, [3]string{"大明 ", "不牛", "未知"})
	//在数组中,无法append,换成slice后可以更改
	newPersonInfos2 = append(newPersonInfos2, "大明 ", "不牛", "未知")
	//d2 := []int{1,2,3}
	fmt.Println("用降维方式输出: ")
	for d1, d1val := range newPersonInfos2 {
		for d2, d2val := range d1val {
			fmt.Println(d1, d1val, d2, "val:", d2val)
		}
	}
	add()
}

//练习Ctrl+NumLock+ 折叠快捷键
func add() {
	var a, b int = 2, 3
	for i := 1; i <= 5; i++ {
		if i <= a {
			fmt.Println(a + i)
		} else if i >= b {
			fmt.Println(b * b)

		}
	}

}
