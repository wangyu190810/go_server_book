package main

import "fmt"

func map_main() {
	classMates1 := make(map[int]string)
	classMates1[0] = "hello"
	classMates1[1] = "你好"
	classMates1[2] = "buybuy"
	classMates1[3] = "thinks"
	fmt.Printf("id %v is %v\n", 1, classMates1[1])
	classMates2 := map[int]string{
		0: "小红",
		1: "小明",
		2: "小绿",
	}
	fmt.Printf("id %v is %v\n", 2, classMates2[1])

}

func vct_main() {
	nums := [...]int{1, 2, 3, 56, 9}
	for k, v := range nums {
		fmt.Println(k, v, " ")
	}
	slis := []int{1, 3, 4, 6, 7}
	for k, v := range slis {
		fmt.Println(k, v, " ")

	}

	tmpMap := map[int]string{
		0: "1212",
		1: "asdfas",
		2: "fasdfa",
	}
	for k, v := range tmpMap {
		fmt.Println(k, v, " ")
	}

}

func main() {
	map_main()
	vct_main()
}
