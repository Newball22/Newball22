package demo3_inter

import "testing"

//测试
func TestMachine_Work(t *testing.T) {
	p := NewPhone("APPLE")
	m := Machine{}
	m.Work(p)

	ip := NewIPad("HUAWEI")
	m.Work(ip)
}
