package demo2_struct

import (
	"fmt"
	"testing"
)

func TestNewPerson(t *testing.T) {
	p := NewPerson("ball", 22, "GZ")
	fmt.Println(p)

	p.SetName("New")
	fmt.Println(p.GetName())

	p.SetAddr("SZ")
	fmt.Println(p.GetAddr())
}
