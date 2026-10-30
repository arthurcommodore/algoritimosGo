package main

import (
	"fmt"

	"github.com/arthurcommodore/algoritimosGo/algoritimo"
)

func main() {

	data := algoritimo.SelectionSort([]int{60, 20, 10, 70, 30, 60, 30, 20, 50})
	fmt.Println(data)
}
