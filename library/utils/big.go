package utils

import (
	"math"
	"math/big"
	"strings"
)

const (
	COIN_BASE float64 = 1000000000000000000
)

func EmptyStr(s string) string {
	if s == "" {
		return "0"
	}
	return s
}
func Add(a string, b string) (string, bool) {

	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	rb, ok := new(big.Int).SetString(EmptyStr(b), 0)
	if !ok {
		return "", ok
	}

	rb = new(big.Int).Add(ra, rb)

	return rb.String(), true
}

func Sub(a string, b string) (string, bool) {

	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	rb, ok := new(big.Int).SetString(EmptyStr(b), 0)
	if !ok {
		return "", ok
	}

	rb = new(big.Int).Sub(ra, rb)

	return rb.String(), true
}

func Mul(a string, b string) (string, bool) {

	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	rb, ok := new(big.Int).SetString(EmptyStr(b), 0)
	if !ok {
		return "", ok
	}

	rb = new(big.Int).Mul(ra, rb)
	return rb.String(), true
}

func Divide(a string, b string, prec int) (string, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(a))
	if !ok {
		return "", ok
	}

	rb, ok := new(big.Float).SetString(EmptyStr(b))
	if !ok {
		return "", ok
	}

	rb = rb.SetPrec(64).Quo(ra, rb)

	return rb.Text('f', prec), true
}

// IsLessThenZero if a < 0 then true, else false
func IsLessThenZero(a string) bool {
	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return ok
	}
	return ra.Cmp(big.NewInt(0)) == -1
}

// MulBase 剩以基数
func MulBase(a string) (string, bool) {
	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	return new(big.Int).Mul(ra, big.NewInt(int64(COIN_BASE))).String(), true
}

func MulByDecimal(a string, decimal int, prec int) (string, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(a))
	if !ok {
		return "", ok
	}

	rb := big.NewFloat(math.Pow10(decimal))

	rb = rb.SetPrec(64).Mul(ra, rb)

	return rb.Text('f', prec), true
}

// DivideBase 除以基数
func DivideBase(a string, prec int) (string, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(a))
	if !ok {
		return "", ok
	}

	rb := big.NewFloat(COIN_BASE)

	rb = rb.SetPrec(64).Quo(ra, rb)

	return rb.Text('f', prec), true
}

func DivideByDecimal(a string, decimal int, prec int) (string, bool) {
	return ToDecimalPrecious(a, int32(decimal), int32(prec)).String(), true
}

func DivideByDecimalForLimit(a string, decimal int, prec int) (string, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(a))
	if !ok {
		return "", ok
	}

	rb := big.NewFloat(math.Pow10(decimal))
	rb = rb.SetPrec(64).Quo(ra, rb)

	v := rb.Text('f', prec)
	v2 := strings.TrimRight(v, "0")
	v2Len := len(v2)
	if v2[v2Len-1] == '.' {
		v2 = v2[:v2Len-1]
	}
	return v2, true
}

func DivideForLimit(a string, prec int) (string, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(a))
	if !ok {
		return "", ok
	}

	rb := big.NewFloat(1)
	rb = rb.SetPrec(64).Quo(ra, rb)

	v := rb.Text('f', prec)
	v2 := strings.TrimRight(v, "0")
	v2Len := len(v2)
	if v2[v2Len-1] == '.' {
		v2 = v2[:v2Len-1]
	}
	return v2, true
}

func ToFloat(s string) (float64, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(s))
	if !ok {
		return 0, ok
	}

	rb := big.NewFloat(COIN_BASE)

	rb = rb.SetPrec(64).Quo(ra, rb)
	f, _ := rb.Float64()
	return f, true
}

func ToInt(s string) (int64, bool) {
	ra, ok := new(big.Float).SetString(EmptyStr(s))
	if !ok {
		return 0, ok
	}

	rb := big.NewFloat(COIN_BASE)

	rb = rb.SetPrec(64).Quo(ra, rb)
	f, _ := rb.Int64()
	return f, true
}

// 比较大小
// a<b -1; a=b 0;a>b 1
func Cmp(a, b string) (int, bool) {
	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return -2, ok
	}

	rb, ok := new(big.Int).SetString(EmptyStr(b), 0)
	if !ok {
		return -2, ok
	}

	res := ra.Cmp(rb)
	return res, true
}

func FloatMul(a string, b string, prec int) (string, bool) {

	ra, ok := new(big.Float).SetString(EmptyStr(a))
	if !ok {
		return "", ok
	}
	rb, ok := new(big.Float).SetString(EmptyStr(b))
	if !ok {
		return "", ok
	}

	r := new(big.Float).Mul(ra, rb)

	return r.Text('f', prec), true
}
