package main

import (
	"container/list"
	"fmt"
)

func main() {
	var integerList list.List

	integerList.PushBack(11)
	integerList.PushBack(23)
	integerList.PushBack(34)

	for element := integerList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
