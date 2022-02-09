package main

import "fmt"

func main() {
	a := 9
	fmt.Printf("运算符&,读出变量a的内存地址为%p\n", &a)
	for {
		var name string
		fmt.Print("姓名: ")
		fmt.Scanln(&name)

		var weight float64
		fmt.Print("体重(KG) : ")
		fmt.Scanln(&weight)

		var tall float64
		fmt.Print("身高 (m) ")
		fmt.Scanln(&tall)
		//fmt包下的扫描函数,在go语言中使用&取值指针进行获值

		var bmi float64 = weight / (tall * tall)
		var age int
		fmt.Print("年龄: ")
		fmt.Scanln(&age)

		var sexWeight int
		var sex string = "男"
		fmt.Print("性别(male/female): ")
		fmt.Scanln(&sex)

		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
		fmt.Println("体脂率是: ", fatRate)

		if sex == "男" {
			// 编写男性的体脂率与体脂状态表
			if age >= 18 && age <= 39 {
				if fatRate <= 0.1 {
					fmt.Println("目前是:偏瘦,多多锻炼,增强体质.")

				} else if fatRate > 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是:标准.太棒了,Keep moving")

				} else if fatRate > 0.16 && fatRate <= 0.21 {
					fmt.Println("目前是:偏胖,少吃高热量,多锻炼")

				} else if fatRate > 0.21 && fatRate <= 0.26 {
					fmt.Println("目前是:肥胖,少吃,多锻炼")

				} else {
					fmt.Println("目前是非常肥胖,马上行动")
				}
			} else if age >= 40 && age <= 59 {
				//todo
			} else if age >= 60 {
				//todo
			} else {
				fmt.Println("我们不参考未成人的体脂率,因变动太大,无法评判.")
			}
		} else {
			//todo 编写女性的体脂率与体脂率
		}
		var whetherContinue string
		fmt.Printf("是否录入下一个(yes/no) ? ")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "yes" {
			break
		}
//暂停会导致栈溢出,Process finished with the exit code -1073741510 (0xC000013A: interrupted by Ctrl+C)
	}
}
