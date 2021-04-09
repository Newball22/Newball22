package stagety

import (
	"fmt"
	"testing"
)

func TestNewCashContext(t *testing.T) {
	c := NewCashContext("满100返20")
	data2 := c.GetMoney(120)
	fmt.Println(data2)
}
