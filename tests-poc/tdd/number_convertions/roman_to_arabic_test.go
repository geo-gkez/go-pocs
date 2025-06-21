package numberconvertions

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	arabicNumber uint16
	romanNumber  string
}{
	{arabicNumber: 1, romanNumber: "I"},
	{arabicNumber: 2, romanNumber: "II"},
	{arabicNumber: 3, romanNumber: "III"},
	{arabicNumber: 4, romanNumber: "IV"},
	{arabicNumber: 5, romanNumber: "V"},
	{arabicNumber: 6, romanNumber: "VI"},
	{arabicNumber: 7, romanNumber: "VII"},
	{arabicNumber: 8, romanNumber: "VIII"},
	{arabicNumber: 9, romanNumber: "IX"},
	{arabicNumber: 10, romanNumber: "X"},
	{arabicNumber: 14, romanNumber: "XIV"},
	{arabicNumber: 18, romanNumber: "XVIII"},
	{arabicNumber: 20, romanNumber: "XX"},
	{arabicNumber: 39, romanNumber: "XXXIX"},
	{arabicNumber: 40, romanNumber: "XL"},
	{arabicNumber: 47, romanNumber: "XLVII"},
	{arabicNumber: 49, romanNumber: "XLIX"},
	{arabicNumber: 50, romanNumber: "L"},
	{arabicNumber: 90, romanNumber: "XC"},
	{arabicNumber: 100, romanNumber: "C"},
	{arabicNumber: 400, romanNumber: "CD"},
	{arabicNumber: 500, romanNumber: "D"},
	{arabicNumber: 900, romanNumber: "CM"},
	{arabicNumber: 1000, romanNumber: "M"},
	{arabicNumber: 1984, romanNumber: "MCMLXXXIV"},
	{arabicNumber: 1991, romanNumber: "MCMXCI"},
	{arabicNumber: 2023, romanNumber: "MMXXIII"},
	{arabicNumber: 3549, romanNumber: "MMMDXLIX"},
	{arabicNumber: 3999, romanNumber: "MMMCMXCIX"},
}

func TestConvertingToRoman(t *testing.T) {
	for _, testcase := range cases {
		t.Run(fmt.Sprintf("convert %d to roman %q", testcase.arabicNumber, testcase.romanNumber),
			func(t *testing.T) {
				got := ConvertToRoman(testcase.arabicNumber)
				want := testcase.romanNumber
				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, testcase := range cases {
		t.Run(fmt.Sprintf("convert %q to arabic %d", testcase.romanNumber, testcase.arabicNumber),
			func(t *testing.T) {
				got := ConvertToArabic(testcase.romanNumber)
				want := testcase.arabicNumber
				if got != want {
					t.Errorf("got %d, want %d", got, want)
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

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
