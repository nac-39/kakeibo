package entity

type Money int

func (m Money) IsNegative() bool {
	return m < 0
}
