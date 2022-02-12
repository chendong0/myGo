package main

import "fmt"

func main() {
	//leftMap, rightMap := map[string]int{}, map[string]int{}
	leftMap := map[string]int{}
	rightMap := map[string]int{}

	//leftMap{"语文"} = 80//leftMap is not a type
	//rightMap{"math"} = 45//missing key in map literal
	//犯的错是用错花括号{},应是方括号

	leftMap["语文"] = 99
	rightMap["math"] = 40
	rightMap["psychology"] = 80
	for k,v := range rightMap{
		leftMap[k] = v
	}
	fmt.Println(leftMap)//map类型中的key和value一样,可以说他们是同一个map类型

	for k, v := range rightMap {
		rightMap[k] = v

	}
	fmt.Println(rightMap)

	for i, j := range leftMap {
		leftMap[i] = j

	}
	fmt.Println(leftMap)

}
