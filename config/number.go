package config

import (
	"fmt"
	"strings"
)

type Number string

func NewNumber(v string) *Number {
	n := Number(v)
	return &n
}

func (n *Number) Type() ValueType {
	return NumberType
}

func (n *Number) Int64() (int64, error) {
	var d int64
	_, err := fmt.Sscanf(string(*n), "%d", &d)
	if err != nil {
		return 0, err
	}
	return d, nil
}

func (n *Number) Float64() (float64, error) {
	var d float64
	_, err := fmt.Sscanf(string(*n), "%f", &d)
	if err != nil {
		return 0, err
	}
	return d, nil
}

func (n *Number) SetInt64(v int64) {
	*n = Number(fmt.Sprintf("%d", v))
}

func (n *Number) SetFloat64(v float64) {
	*n = Number(fmt.Sprintf("%f", v))
}

func (n *Number) String() string {
	return string(*n)
}

func (n *Number) Format(deep int) string {
	return strings.Repeat(" ", deep*4) + n.String()
}

func (n *Number) Reset() {
	n.SetInt64(0)
}
