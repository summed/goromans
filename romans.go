package romans

import (
	"bytes"
	"fmt"
	"sort"
	"unicode"
)

type numeralByArabic []numeral

func (v numeralByArabic) Len() int            { return len(v) }
func (v numeralByArabic) Swap(this, that int) { v[this], v[that] = v[that], v[this] }
func (v numeralByArabic) Less(this, that int) bool {
	return v[this].arabic > v[that].arabic
}

type numeral struct {
	roman  rune
	arabic uint
}

var (
	sets = []numeral{
		numeral{'M', 1000},
		numeral{'D', 500},
		numeral{'C', 100},
		numeral{'L', 50},
		numeral{'X', 10},
		numeral{'V', 5},
		numeral{'I', 1},
	}

	initialized   bool
	romanNumerals = make(map[rune]numeral)
)

func initialize() {
	if !initialized {
		for _, s := range sets {
			romanNumerals[s.roman] = s
		}
		sort.Sort(numeralByArabic(sets)) // Strictly not required, since the correct (DESC) order is set at initialization.
		initialized = true
	}
}

// IsRomanNumerals returns true if able to parse string as roman numerals
func IsRomanNumerals(romans string) bool {
	if _, err := RtoA(romans); err != nil {
		return false
	}
	return true
}

// RtoA converts a string of roman numerals to arabic numerals
func RtoA(romans string) (out uint, err error) {
	initialize()
	var last uint
	if len(romans) == 0 {
		return out, fmt.Errorf("Empty string when parsing to roman numerals")
	}
	for i := 0; i < len(romans); i++ {
		if s, ok := romanNumerals[unicode.ToUpper(rune(romans[i]))]; ok {
			if s.arabic > last && last > 0 {
				out -= 2 * last
			}
			out += s.arabic
			last = s.arabic
		} else {
			return out, fmt.Errorf("Unable to '%s' to roman numerals, because of character '%s'", romans, string(romans[i]))
		}
	}
	return out, nil
}

// AtoR converts arabic numerals to roman numerals
func AtoR(arabic uint) (romans string) {
	initialize()
	var out bytes.Buffer
	var major, minor numeral
	for arabic > 0 {
		for i := 0; i < len(sets); i++ {
			major = sets[i]

			if arabic == major.arabic {
				out.WriteRune(major.roman)
				arabic -= major.arabic
				goto loopEnd
			}

			if i < len(sets) {
				for j := i + 1; j < len(sets); j++ {
					minor = sets[j]
					if major.arabic/minor.arabic == 2 { // if minor is half of major (M&D, C&L, X&V), then skip - and let it be 'handled' as major later
						continue
					}
					if arabic-(major.arabic-minor.arabic) == 0 {
						out.WriteRune(minor.roman)
						out.WriteRune(major.roman)
						arabic -= major.arabic - minor.arabic
						goto loopEnd
					}
				}
			}

			if arabic > major.arabic {
				divs := uint(arabic / major.arabic)
				arabic -= major.arabic * divs
				for ; divs > 0; divs-- {
					out.WriteRune(major.roman)
				}
				goto loopEnd
			}
		}
	loopEnd:
	}

	return out.String()
}
