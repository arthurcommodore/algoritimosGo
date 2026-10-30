package main

import (
	"fmt"

	"github.com/arthurcommodore/algoritimosGo/algoritimo"
)

func main() {

	data := algoritimo.Selection_sort([]int{60, 20, 10, 70, 30, 60, 30, 20, 50})
	fmt.Println(algoritimo.Binary_search(20, data))
}
