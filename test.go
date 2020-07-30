package main

import (
	"fmt"

	"github.com/lierbai/project/util"
)

func main() {
	intSet := util.New()
	for i := 0; i < 99999; i++ {
		intSet.Add(string(i))
	}
	for i := 90000; i < 99999; i++ {
		intSet.Remove(string(i))
	}
	fmt.Println(intSet.Len())
}
