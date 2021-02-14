package demo4_struct

import (
	"fmt"
	"testing"
)

func TestUser_NewUser(t *testing.T) {
	user := NewUser(22, "ball", 33)
	fmt.Println(user)
	user.Run()
}

func TestPersonFactory(t *testing.T) {
	var tp PersonTye
	p := tp.PersonFactory(2)
	p.Eat()
}
