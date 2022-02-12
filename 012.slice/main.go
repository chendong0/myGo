package main

import "fmt"

func main() {
	李老板 := "1"
	fmt.Println("testing测试中文变量", 李老板)

	names := []string{"小强", "光头强", "熊大"}
	fr := []float64{28, 5, 3}

	names = append(names, "熊二")
	fr = append(fr, 19)

	for i, name := range names {
		if name == "光头强" {
			fmt.Printf("%s 的体脂率是 %f\n", name, fr[i])

		}

	}
	fmt.Println("定义Map")

	var m1 map[string]int = nil
	//m1["a"] = 1 // panic on nil map//assignment to entry in nil map
	delete(m1, "a")
	fmt.Println("m1 没有实例化,直接取数: ", m1["a"])

	m2 := map[string]float64{}
	m3 := map[string]int{"梨子": 66, "芒果": 88, "苹果": 99}
	fmt.Println(m1, m2, m3)

	fmt.Println("梨子的分数", m3["梨子"])
	fmt.Println("芒果的分数", m3["芒果"])
	fmt.Println("苹果的分数", m3["苹果"])

	guoScore, ok := m3["苹果"]
	fmt.Println(guoScore, ">>>>", ok)

	m3["苹果"] = 77
	fmt.Println("苹果的分数: ", m3["小强"])
	guoScore, ok = m3["苹果"]
	fmt.Println(guoScore, "<<<<<", ok)

	for name,score := range m3{
		fmt.Println(name, "=",score )
	}
	for i,v := range m3{//practice for....range
		fmt.Println(i, "=",v )
	}
	for j,k := range m3{
		fmt.Println(j, "=",k )
	}

}
