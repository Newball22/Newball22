package demo2_struct

type person struct {
	Name string
	age  int
	addr string
}

//结构体构造函数
func NewPerson(name string, age int, addr string) *person {
	return &person{
		Name: name,
		age:  age,
		addr: addr,
	}
}

func (p *person) SetName(name string) {
	p.Name = name
}

func (p *person) GetName() string {
	return p.Name
}

func (p *person) SetAddr(str string) {
	p.addr = str
}

func (p *person) GetAddr() string {
	return p.addr
}
