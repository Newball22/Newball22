package demo4_struct

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func NewUser(id int, name string, age int) *User {
	return &User{
		ID:   id,
		Name: name,
		Age:  age,
	}
}

func (u *User) Run() {
	fmt.Printf("%s在跑\n", u.Name)
}

type Person interface {
	Eat()
}

type XiaoMing struct {
	name string
}

func (x *XiaoMing) Eat() {
	fmt.Printf("%s在吃饭\n", x.name)
}

type Tom struct {
	name string
}

func (t *Tom) Eat() {
	fmt.Printf("%s在吃饭\n", t.name)
}

type PersonTye int

func (p *PersonTye) PersonFactory(num int) Person {
	switch num {
	case 1:
		return &XiaoMing{name: "小明"}
	case 2:
		return &Tom{name: "Tom"}
	default:
		fmt.Println("Error Request")
		return nil
	}
}
