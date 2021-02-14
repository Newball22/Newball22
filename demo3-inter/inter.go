package demo3_inter

import (
	"fmt"
)

type USB interface {
	start()
	stop()
}

type Phone struct {
	name string
}

func (p *Phone) start() {
	fmt.Println(p.name, "start")
}

func (p *Phone) stop() {
	fmt.Println(p.name, "stop")
}

func NewPhone(name string) *Phone {
	return &Phone{name}
}

type IPad struct {
	name string
}

func NewIPad(name string) *IPad {
	return &IPad{name}
}

func (ip *IPad) start() {
	fmt.Println(ip.name, "start")
}

func (ip *IPad) stop() {
	fmt.Println(ip.name, "stop")
}

type Machine struct {
}

func (m *Machine) Work(u USB) {
	u.start()
	u.stop()
}
