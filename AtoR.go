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
		sort.Sort(numeralByArabic(sets))
		initialized = true
	}
}

// IsRomanNumeral returns true if able to parse string as roman numerals
func IsRomanNumeral(romans string) bool {
	if _, err := RToI(romans); err != nil {
		return true
	}
	return false
}

// RToI converts a string of roman numeral to arabic numeral
func RToI(romans string) (sum uint, err error) {
	initialize()
	var last uint
	if len(romans) == 0 {
		return sum, fmt.Errorf("Empty string when parsing to roman numerals")
	}
	for i := 0; i < len(romans); i++ {
		if s, ok := romanNumerals[unicode.ToUpper(rune(romans[i]))]; ok {
			if s.arabic > last && last > 0 {
				sum -= 2 * last
			}
			sum += s.arabic
			last = s.arabic
		} else {
			return sum, fmt.Errorf("Unable to '%s' to roman numerals, because of character '%s'", romans, string(romans[i]))
		}
	}
	return sum, nil
}

// IToR converts a arabic numeral to roman numeral
func IToR(arabic uint) (romans string) {
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
