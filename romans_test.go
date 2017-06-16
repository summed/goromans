package romans

import (
	"strings"
	"testing"
	"unicode"
)

var (
	values = map[string]uint{
		"I":          1,
		"V":          5,
		"X":          10,
		"L":          50,
		"C":          100,
		"D":          500,
		"M":          1000,
		"III":        3,
		"IV":         4,
		"VI":         6,
		"LXXIV":      74,
		"LXXXIX":     89,
		"MDCCC":      1800,
		"DCCCXC":     890,
		"MCCCCXXVI":  1426,
		"MMDCCCLVII": 2857,
		"MMVL":       2045,
	}
)

func TestIsRomanNumerals(t *testing.T) {
	for k := range values {
		if !IsRomanNumerals(k) {
			t.Errorf("'%s' did not match expected as roman numeral value", k)
		}
	}
	var (
		from  = rune('a')
		to    = rune('Z')
		check rune
	)
	for i := from; i <= to; i++ {
		check = unicode.ToUpper(i)
		if _, ok := romanNumerals[check]; !ok {
			if IsRomanNumerals(string(i)) {
				t.Errorf("'%c' unexpectedly match as roman numeral value", i)
			}
		}
	}
}

func TestRtoA(t *testing.T) {
	for k, v := range values {
		if r, err := RtoA(k); r != v || err != nil {
			t.Errorf("'%s' did not match expected value of '%d', but '%d' instead", k, v, r)
		}
	}
}

func TestAtoR(t *testing.T) {
	for k, v := range values {
		if r := AtoR(v); strings.Compare(r, k) != 0 {
			t.Errorf("'%d' did not match expected value of '%s', but '%s' instead", v, k, r)
		}
	}
}
