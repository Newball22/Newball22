package demo1_oop

import (
	"fmt"
	"testing"
)

func TestLog_Sum(t *testing.T) {
	var a log=2
	res:=a.Sum(3)
	fmt.Println(res)
}
