package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancyNumber, isType := fnb.(FancyNumber)
	if !isType {
		return 0
	}

	value, err := strconv.Atoi(fancyNumber.Value())
	if err != nil {
		return 0
	}
	return value
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	number := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(number))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case float64:
		return DescribeNumber(i.(float64))
	case int:
		return DescribeNumber(float64(i.(int)))
	case NumberBox:
		return DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		return DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		return "Return to sender"
	}
}
