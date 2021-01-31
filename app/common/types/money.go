package types

type Money int64

func (m Money) GetInt64() int64 {
	return int64(m)
}

func (m Money) GetFloat64() float64 {
	return float64(m) / 100
}
