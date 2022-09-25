package bigint

import (
	"errors"
	"strconv"
	"strings"
)

// Creating struct
type Bigint struct {
	Value string
}

// error
var (
	ErrorNotNumber = errors.New("input cannot be a number")
	ErrorLength    = errors.New("out of length")
)

// validation
func validation(num string) error {

	_, err := strconv.ParseInt(num, 10, 32)

	if err != nil {
		return ErrorLength
	}

	var allowed = "0123456789"

	if strings.HasPrefix(num, "-") {
		num = strings.Replace(num, "-", "", 1)

	}
	if strings.HasPrefix(num, "+") {
		num = strings.Replace(num, "+", "", 1)

	}

	str := strings.Split(num, "")

	for _, v := range str {
		if !strings.Contains(allowed, v) {
			return ErrorNotNumber
		}
	}

	return nil
}

// clean
func clean(num string) string {
	prefix := ""
	if strings.HasPrefix(num, "-") {
		prefix = "-"
		num = num[1:]
	}

	if strings.HasPrefix(num, "+") {
		num = num[1:]
	}

	for strings.HasPrefix(num, "0") {
		num = strings.TrimPrefix(num, "0")

	}
	if num == "" {
		return num + "0"
	}
	return prefix + num
}

// NewInt
func NewInt(num string) (Bigint, error) {
	err := validation(num)
	if err != nil {
		return Bigint{Value: ""}, err

	}
	return Bigint{
		Value: clean(num),
	}, err
}

// Set
func (z *Bigint) Set(num string) error {
	err := validation(num)
	if err != nil {
		return err
	}
	z.Value = clean(num)
	return err
}

// +
func Add(a, b Bigint) Bigint {
	A, _ := strconv.Atoi(a.Value)
	B, _ := strconv.Atoi(b.Value)
	c := strconv.Itoa(A + B)
	return Bigint{
		Value: c,
	}
}

// -
func Sub(a, b Bigint) Bigint {
	A, _ := strconv.Atoi(a.Value)
	B, _ := strconv.Atoi(b.Value)
	c := strconv.Itoa(A - B)
	return Bigint{
		Value: c,
	}
}

// *
func Multiply(a, b Bigint) Bigint {
	A, _ := strconv.Atoi(a.Value)
	B, _ := strconv.Atoi(b.Value)
	c := strconv.Itoa(A * B)
	return Bigint{
		Value: c,
	}
}

// %
func Mod(a, b Bigint) Bigint {
	A, _ := strconv.Atoi(a.Value)
	B, _ := strconv.Atoi(b.Value)
	if A == 0 || B == 0 {
		return Bigint{
			"Division by is undifined ",
		}
	}
	c := strconv.Itoa(A % B)
	return Bigint{
		Value: c,
	}
}

// /
func Divide(a, b Bigint) Bigint {
	A, _ := strconv.Atoi(a.Value)
	B, _ := strconv.Atoi(b.Value)
	if A == 0 || B == 0 {
		return Bigint{
			"Division by is undifined ",
		}
	}
	c := strconv.Itoa(A / B)
	return Bigint{
		Value: c,
	}
}

// ||
func (z Bigint) Abs() Bigint {
	c, _ := strconv.Atoi(z.Value)

	if c < 0 {
		c *= (-1)
		z.Value = strconv.Itoa(c)
	}

	return Bigint{
		Value: z.Value,
	}

}
