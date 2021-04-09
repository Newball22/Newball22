package stagety

type cashSuper interface {
	AcceptMoney(money float64) float64
}

//正常价格
type cashNormal struct {
}

func newCashNormal() *cashNormal {
	instance := new(cashNormal)
	return instance
}

func (c *cashNormal) AcceptMoney(money float64) float64 {
	return money
}

//打折
type cashRebate struct {
	Rebate float64
}

func newCashRebate(rebate float64) *cashRebate {
	return &cashRebate{
		Rebate: rebate,
	}
}

func (bate *cashRebate) AcceptMoney(money float64) float64 {
	return money * bate.Rebate
}

//直接返利(满减)
type cashReturn struct {
	MoneyCondition float64
	MoneyReturn    float64
}

func newCashReturn(moneyCondition, moneyReturn float64) *cashReturn {
	return &cashReturn{
		MoneyCondition: moneyCondition,
		MoneyReturn:    moneyReturn,
	}
}

func (ct *cashReturn) AcceptMoney(money float64) float64 {
	if money >= ct.MoneyCondition {
		min := int(money / ct.MoneyCondition)
		return money - float64(min)*ct.MoneyReturn
	}
	return money
}

type CashContext struct {
	Strategy cashSuper
}

func NewCashContext(cashType string) CashContext {
	c := new(CashContext)
	switch cashType {
	case "打八折":
		c.Strategy = newCashRebate(0.8)
	case "满100返20":
		c.Strategy = newCashReturn(100.0, 20.0)
	default:
		c.Strategy = newCashNormal()
	}
	return *c
}

func (c *CashContext) GetMoney(money float64) float64 {
	return c.Strategy.AcceptMoney(money)
}
