package domain

//Amount 金额
type Amount struct {
	euros float64
}

//CreateAmount  构造器
func CreateAmount(euros, cents int) *Amount {
	return &Amount{euros: float64(100*euros + cents)}
}

//CreateAmountByEuros 构造器
func CreateAmountByEuros(euros int) *Amount {
	return &Amount{float64(100 * euros)}
}

//Plus add amount
func (r *Amount) Plus(amount *Amount) Amount {
	r.euros += amount.euros
	return *r
}

//Minus minus amount
func (r *Amount) Minus(amount Amount) *Amount {
	r.euros -= amount.euros
	return r
}

//Times times
func (r *Amount) Times(factor float64) *Amount {
	r.euros *= factor
	return r
}

//ToDouble return Double value
func (r *Amount) ToDouble() float64 {
	return r.euros
}
