package numberconvertions

import "strings"

type RomanNumeral struct {
	value  uint16
	symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabicNumber uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabicNumber >= numeral.value {
			result.WriteString(numeral.symbol)
			arabicNumber -= numeral.value
		}
	}

	return result.String()
}

func ConvertToArabic(romanNumber string) uint16 {
	var arabic uint16

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(romanNumber, numeral.symbol) {
			arabic += numeral.value
			romanNumber = strings.TrimPrefix(romanNumber, numeral.symbol)
		}
	}

	return arabic
}
