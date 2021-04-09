package main

import (
	"fmt"
	"time"
)

func Find(slice []string, val string) bool {
	if len(slice) <= 0 {
		return false
	}
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(time.Now().Unix())
}
