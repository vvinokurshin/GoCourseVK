package uniq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var uniqTestsPositive = []struct {
	name    string
	input   []string
	options Options
	output  []string
}{
	{
		"Simple test",
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		Options{},
		[]string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
	{
		"Count test",
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		Options{
			count: true,
		},
		[]string{
			"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik.",
		},
	},
	{
		"Duplicated test",
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		Options{
			duplicated: true,
		},
		[]string{
			"I love music.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
	},
	{
		"Unique test",
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		Options{
			unique: true,
		},
		[]string{
			"",
			"Thanks.",
		},
	},
	{
		"Insensitive test",
		[]string{
			"I LOVE MUSIC.",
			"I love music.",
			"I LoVe MuSiC.",
			"",
			"I love MuSIC of Kartik.",
			"I love music of kartik.",
			"Thanks.",
			"I love music of kartik.",
			"I love MuSIC of Kartik.",
		},
		Options{
			insensitive: true,
		},
		[]string{
			"I LOVE MUSIC.",
			"",
			"I love MuSIC of Kartik.",
			"Thanks.",
			"I love music of kartik.",
		},
	},
	{
		"Fields test",
		[]string{
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		Options{
			field: 1,
		},
		[]string{
			"We love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	{
		"Chars test",
		[]string{
			"I love music.",
			"A love music.",
			"C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		Options{
			chars: 1,
		},
		[]string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
	},
	{
		"Several flags test",
		[]string{
			"I love music.",
			"W love music.",
			"B love music.",
			"",
			"I love music of Kartik.",
			"Im love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		Options{
			count: true,
			chars: 1,
		},
		[]string{
			"3 I love music.",
			"1 ",
			"1 I love music of Kartik.",
			"1 Im love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik.",
		},
	},
	{
		"Empty test",
		[]string{},
		Options{
			chars: 1,
		},
		[]string{},
	},
}

var uniqTestsNegative = []struct {
	name    string
	input   []string
	options Options
	output  []string
}{
	{
		"Interchangeable flags test",
		[]string{},
		Options{
			count:      true,
			duplicated: true,
		},
		[]string{},
	},
}

func TestPositive(t *testing.T) {
	for _, tt := range uniqTestsPositive {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Uniq(tt.input, tt.options)

			require.NoError(t, err)
			require.Equal(t, result, tt.output)
		})
	}
}

func TestNegative(t *testing.T) {
	for _, tt := range uniqTestsNegative {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Uniq(tt.input, tt.options)

			require.Error(t, err)
			require.Equal(t, result, tt.output)
		})
	}
}
