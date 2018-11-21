package decimal

import (
	"github.com/shopspring/decimal"
)

//If a < b return -1, a == b return 0, a > b return 1
func CompareStr(a, b string) int{
	decimal_a, _ := decimal.NewFromString(a)
	decimal_b, _ := decimal.NewFromString(b)
	return decimal_a.Cmp(decimal_b)
}

//Sum strings
func SumStr(first string, rest ...string) string{
	decimal_f, _ := decimal.NewFromString(first)
	for _, r := range rest{
		decimal_r, _ := decimal.NewFromString(r)
		decimal_f = decimal_f.Add(decimal_r)
	}
	return decimal_f.String()
}

//Add string
func AddStr(a, b string) string{
	decimal_a, _ := decimal.NewFromString(a)
	decimal_b, _ := decimal.NewFromString(b)

	return decimal_a.Add(decimal_b).String()
}

//Max string
func MaxStr(first string, rest...string) string{
	cotainer := make([]decimal.Decimal,1)
	decimal_f, _ := decimal.NewFromString(first)
	for _, r := range rest{
		decimal_r, _ := decimal.NewFromString(r)
		cotainer = append(cotainer, decimal_r)
	}
	max := decimal.Max(decimal_f, cotainer...)
	return max.String()
}

//Min string
func MinStr(first string, rest...string) string{
	cotainer := make([]decimal.Decimal,1)
	decimal_f, _ := decimal.NewFromString(first)
	for _, r := range rest{
		decimal_r, _ := decimal.NewFromString(r)
		cotainer = append(cotainer, decimal_r)
	}
	min := decimal.Min(decimal_f, cotainer...)
	return min.String()
}