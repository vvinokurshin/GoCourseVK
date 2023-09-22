package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var calcTestsPositive = []struct {
	name   string
	input  string
	output float64
}{
	{
		"Simple sum test",
		"1+1",
		2,
	},
	{
		"Simple diff test",
		"1-2",
		-1,
	},
	{
		"Simple mul test",
		"1*3",
		3,
	},
	{
		"Simple div test",
		"9/2",
		4.5,
	},
	{
		"Test with all operations",
		"1+9/3-4*2",
		-4,
	},
	{
		"Text with brackets",
		"(1+8)/3",
		3,
	},
	{
		"Text with nested brackets",
		"((2+8)-4)*3",
		18,
	},
	{
		"Text with only one number",
		"12",
		12,
	},
}

var calcTestsNegative = []struct {
	name   string
	input  string
	output float64
}{
	{
		"Letters in the expression",
		"a",
		0,
	},
	{
		"Several operations in a row",
		"1++1",
		0,
	},
	{
		"Empty expression",
		"",
		0,
	},
	{
		"Division by zero",
		"1/0",
		0,
	},
	{
		"Only one operation without numbers",
		"*",
		0,
	},
}

func TestPositive(t *testing.T) {
	for _, tt := range calcTestsPositive {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.input)

			require.NoError(t, err)
			require.Equal(t, tt.output, result)
		})
	}
}

func TestNegative(t *testing.T) {
	for _, tt := range calcTestsNegative {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.input)

			require.Error(t, err)
			require.Equal(t, tt.output, result)
		})
	}
}
