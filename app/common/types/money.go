package types

type Money int64

// ToInt64 - money to int64
func (m Money) ToInt64() int64 {
	return int64(m)
}

// ToFloat64 - money to float64
func (m Money) ToFloat64() float64 {
	return float64(m) / 100
}
