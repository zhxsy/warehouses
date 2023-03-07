package example

import (
	"fmt"
	"github.com/cfx/warehouses/library/utils"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

func Test_value(b *testing.T) {
	gasFee := "0.00"
	//gasFee := "1000000000000000000"
	//gasFee = "10000000000000000"
	decimal := int(18)
	value, _ := utils.DivideByDecimalForLimit(gasFee, decimal, 8)
	//value, _ := utils.DivideByDecimal(gasFee, decimal, 8)
	fmt.Println(value)
}

func Test_value2(b *testing.T) {
	gasFee := "1.223456"
	//gasFee := "1000000000000000000"
	//gasFee = "10000000000000000"
	value, _ := utils.DivideForLimit(gasFee, 4)
	//value, _ := utils.DivideByDecimal(gasFee, decimal, 8)
	fmt.Println(value)
}

func Test_value3(b *testing.T) {
	gasFee := 1.0003

	strScore := strconv.FormatFloat(gasFee, 'f', 2, 64)
	fmt.Println(strScore)

}

func Test_value4(t *testing.T) {
	V := "10000000000000000000000000000000000000000000000"
	//V = "10000000000000000000000000000"
	//V = "20000000 000000000000000000"
	value, _ := DivideByDecimalForLimit(V, 36, 8)
	fmt.Println(value)
	//100000000000 00000000000000000
}

func EmptyStr(s string) string {
	if s == "" {
		return "0"
	}
	return s
}

func DivideByDecimalForLimit(a string, decimal int, prec int) (string, bool) {
	ra, ok := new(big.Int).SetString(EmptyStr(a), 0)
	if !ok {
		return "", ok
	}
	str := "1"
	for i := 0; i < decimal; i++ {
		str = str + "0"
	}

	rb, ok := new(big.Int).SetString(EmptyStr(str), 0)
	if !ok {
		return "", ok
	}

	value := new(big.Int).Div(ra, rb)
	really := value.String()

	fmt.Println(ra)
	fmt.Println(rb)
	fmt.Println(value)
	fmt.Println(really)

	rr, ok := new(big.Float).SetString(EmptyStr(really))
	v := rr.Text('f', prec)

	fmt.Println(v)
	v2 := strings.TrimRight(v, "0")
	v2Len := len(v2)
	if v2[v2Len-1] == '.' {
		v2 = v2[:v2Len-1]
	}
	return v2, true
}
