package propertybased

import (
	"fmt"
	"testing"
	"testing/quick"
)

type Case struct {
	Arabic uint16
	Roman  string
}

func TestRomanNumerals(t *testing.T) {
	cases := []Case{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{100, "C"},
		{90, "XC"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
	}

	for _, c := range cases {
		name := fmt.Sprintf("converts %d to %s", c.Arabic, c.Roman)
		t.Run(name, func(t *testing.T) {
			got := ConvertToRoman(c.Arabic)
			want := c.Roman

			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
		t.Run(name, func(t *testing.T) {
			got := ConvertToArabic(c.Roman)
			want := c.Arabic

			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)

		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}
