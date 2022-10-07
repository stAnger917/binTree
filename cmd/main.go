package main

import (
	"fmt"

	"binTree/pkg/tree"
)

func main() {
	oak := tree.Node[map[string]interface{}]{}
	err := oak.Insert(4, map[string]interface{}{
		"some_key": fmt.Sprintf("some_value %v", 4),
	})
	if err != nil {
		fmt.Println(err)
	}
	err = oak.Insert(13, map[string]interface{}{
		"some_key": fmt.Sprintf("some_value %v", 13),
	})
	if err != nil {
		fmt.Println(err)
	}
	err = oak.Insert(19, map[string]interface{}{
		"some_key": fmt.Sprintf("some_value %v", 19),
	})
	if err != nil {
		fmt.Println(err)
	}
	err = oak.Insert(5, map[string]interface{}{
		"some_key": fmt.Sprintf("some_value %v", 5),
	})
	if err != nil {
		fmt.Println(err)
	}
	err = oak.Insert(2, map[string]interface{}{
		"some_key": fmt.Sprintf("some_value %v", 2),
	})
	if err != nil {
		fmt.Println(err)
	}
	val, err := oak.GetValueByIndex(13)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	fmt.Println("Min: ", oak.FindMin())
	fmt.Println("Max:", oak.FindMax())
}
