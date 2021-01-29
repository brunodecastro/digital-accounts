package types

type Money int64

func (m Money) GetInt() int64 {
	return int64(m)
}

func (m Money) GetFloat() float64 {
	return float64(m) / 100
}
