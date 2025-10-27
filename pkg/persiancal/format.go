package persiancal

import (
	"fmt"
	"strconv"
	"strings"
)

// Format formats the JalaliDate according to the given layout.
// Supported tokens:
//   - yyyy: 4-digit year (e.g., 1404)
//   - yy: 2-digit year (e.g., 04)
//   - MM: 2-digit month (e.g., 08)
//   - M: month without leading zero (e.g., 8)
//   - MMMM: Persian month name (e.g., آبان)
//   - MMM: English month name (e.g., Aban)
//   - dd: 2-digit day (e.g., 04)
//   - d: day without leading zero (e.g., 4)
func (j JalaliDate) Format(layout string) string {
	result := layout

	// Replace year
	result = strings.ReplaceAll(result, "yyyy", fmt.Sprintf("%04d", j.Year))
	result = strings.ReplaceAll(result, "yy", fmt.Sprintf("%02d", j.Year%100))

	// Replace month (order matters: MMMM before MMM before MM before M)
	result = strings.ReplaceAll(result, "MMMM", GetMonthNamePersian(j.Month))
	result = strings.ReplaceAll(result, "MMM", GetMonthNameEnglish(j.Month))
	result = strings.ReplaceAll(result, "MM", fmt.Sprintf("%02d", j.Month))
	result = strings.ReplaceAll(result, "M", fmt.Sprintf("%d", j.Month))

	// Replace day
	result = strings.ReplaceAll(result, "dd", fmt.Sprintf("%02d", j.Day))
	result = strings.ReplaceAll(result, "d", fmt.Sprintf("%d", j.Day))

	return result
}

// FormatPersian formats the date with Persian digits
func (j JalaliDate) FormatPersian(layout string) string {
	formatted := j.Format(layout)
	return ToPersianDigits(formatted)
}

// Parse parses a date string according to the given layout.
// Supported tokens: yyyy, yy, MM, M, MMMM, MMM, dd, d
// Supports both Persian and Latin digits.
func Parse(layout, value string) (JalaliDate, error) {
	value = ToLatinDigits(value)

	var year, month, day int
	var err error

	tokens := []struct {
		token  string
		length int
		setter func(string) error
	}{
		{"yyyy", 4, func(s string) error {
			year, err = strconv.Atoi(s)
			return err
		}},
		{"yy", 2, func(s string) error {
			y, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			// Assume 2-digit years are in 1300-1399 range
			if y < 100 {
				year = 1300 + y
			} else {
				year = y
			}
			return nil
		}},
		{"MMMM", 0, func(s string) error {
			m := GetMonthFromPersianName(s)
			if m == 0 {
				return fmt.Errorf("unknown Persian month name: %s", s)
			}
			month = m
			return nil
		}},
		{"MMM", 0, func(s string) error {
			m := GetMonthFromEnglishName(s)
			if m == 0 {
				return fmt.Errorf("unknown English month name: %s", s)
			}
			month = m
			return nil
		}},
		{"MM", 2, func(s string) error {
			month, err = strconv.Atoi(s)
			return err
		}},
		{"M", 0, func(s string) error {
			month, err = strconv.Atoi(s)
			return err
		}},
		{"dd", 2, func(s string) error {
			day, err = strconv.Atoi(s)
			return err
		}},
		{"d", 0, func(s string) error {
			day, err = strconv.Atoi(s)
			return err
		}},
	}

	layoutIdx := 0
	valueIdx := 0

	for layoutIdx < len(layout) && valueIdx < len(value) {
		matched := false

		for _, token := range tokens {
			if strings.HasPrefix(layout[layoutIdx:], token.token) {
				var tokenValue string

				if token.length > 0 {
					if valueIdx+token.length > len(value) {
						return JalaliDate{}, fmt.Errorf("%w: insufficient characters for token %s", ErrParseFailure, token.token)
					}
					tokenValue = value[valueIdx : valueIdx+token.length]
					valueIdx += token.length
				} else {
					if token.token == "MMMM" {
						found := false
						for m := 1; m <= 12; m++ {
							monthName := GetMonthNamePersian(m)
							if strings.HasPrefix(value[valueIdx:], monthName) {
								tokenValue = monthName
								valueIdx += len(monthName)
								found = true
								break
							}
						}
						if !found {
							return JalaliDate{}, fmt.Errorf("%w: could not match Persian month name", ErrParseFailure)
						}
					} else if token.token == "MMM" {
						found := false
						for m := 1; m <= 12; m++ {
							monthName := GetMonthNameEnglish(m)
							if len(value[valueIdx:]) >= len(monthName) &&
								equalFold(value[valueIdx:valueIdx+len(monthName)], monthName) {
								tokenValue = monthName
								valueIdx += len(monthName)
								found = true
								break
							}
						}
						if !found {
							return JalaliDate{}, fmt.Errorf("%w: could not match English month name", ErrParseFailure)
						}
					} else {
						// Single-digit number (M or d)
						endIdx := valueIdx
						for endIdx < len(value) && endIdx < valueIdx+2 && value[endIdx] >= '0' && value[endIdx] <= '9' {
							endIdx++
						}
						if endIdx == valueIdx {
							return JalaliDate{}, fmt.Errorf("%w: expected digit for token %s", ErrParseFailure, token.token)
						}
						tokenValue = value[valueIdx:endIdx]
						valueIdx = endIdx
					}
				}

				if err := token.setter(tokenValue); err != nil {
					return JalaliDate{}, fmt.Errorf("%w: %v", ErrParseFailure, err)
				}

				layoutIdx += len(token.token)
				matched = true
				break
			}
		}

		if !matched {
			if layoutIdx < len(layout) && valueIdx < len(value) {
				if layout[layoutIdx] != value[valueIdx] {
					return JalaliDate{}, fmt.Errorf("%w: expected '%c' but got '%c'", ErrParseFailure, layout[layoutIdx], value[valueIdx])
				}
				layoutIdx++
				valueIdx++
			}
		}
	}

	if layoutIdx < len(layout) || valueIdx < len(value) {
		return JalaliDate{}, fmt.Errorf("%w: layout and value length mismatch", ErrParseFailure)
	}

	j := JalaliDate{Year: year, Month: month, Day: day}
	if err := j.Validate(); err != nil {
		return JalaliDate{}, fmt.Errorf("%w: %v", ErrInvalidDate, err)
	}

	return j, nil
}

// MustParse parses a date string and panics if parsing fails
func MustParse(layout, value string) JalaliDate {
	j, err := Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return j
}

// Common format layouts
const (
	// LayoutISO is the ISO-like format: yyyy-MM-dd
	LayoutISO = "yyyy-MM-dd"

	// LayoutSlash is the slash-separated format: yyyy/MM/dd
	LayoutSlash = "yyyy/MM/dd"

	// LayoutDot is the dot-separated format: yyyy.MM.dd
	LayoutDot = "yyyy.MM.dd"

	// LayoutLong is the long format with month name: dd MMMM yyyy
	LayoutLong = "dd MMMM yyyy"

	// LayoutLongEnglish is the long format with English month name: dd MMM yyyy
	LayoutLongEnglish = "dd MMM yyyy"

	// LayoutShort is the short format: yy/MM/dd
	LayoutShort = "yy/MM/dd"
)
