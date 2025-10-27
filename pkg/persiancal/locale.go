package persiancal

// MonthName represents a month name in different languages
type MonthName struct {
	Persian string
	English string
}

// Persian month names (1-indexed, so 0 is placeholder)
var persianMonthNames = []MonthName{
	{Persian: "", English: ""},                    // placeholder
	{Persian: "فروردین", English: "Farvardin"},    // 1
	{Persian: "اردیبهشت", English: "Ordibehesht"}, // 2
	{Persian: "خرداد", English: "Khordad"},        // 3
	{Persian: "تیر", English: "Tir"},              // 4
	{Persian: "مرداد", English: "Mordad"},         // 5
	{Persian: "شهریور", English: "Shahrivar"},     // 6
	{Persian: "مهر", English: "Mehr"},             // 7
	{Persian: "آبان", English: "Aban"},            // 8
	{Persian: "آذر", English: "Azar"},             // 9
	{Persian: "دی", English: "Dey"},               // 10
	{Persian: "بهمن", English: "Bahman"},          // 11
	{Persian: "اسفند", English: "Esfand"},         // 12
}

// GetMonthNamePersian returns the Persian name of a month (1-12)
func GetMonthNamePersian(month int) string {
	if month < 1 || month > 12 {
		return ""
	}
	return persianMonthNames[month].Persian
}

// GetMonthNameEnglish returns the English transliteration of a month name (1-12)
func GetMonthNameEnglish(month int) string {
	if month < 1 || month > 12 {
		return ""
	}
	return persianMonthNames[month].English
}

// GetMonthFromPersianName returns the month number (1-12) from a Persian name
func GetMonthFromPersianName(name string) int {
	for i := 1; i <= 12; i++ {
		if persianMonthNames[i].Persian == name {
			return i
		}
	}
	return 0
}

// GetMonthFromEnglishName returns the month number (1-12) from an English name (case-insensitive)
func GetMonthFromEnglishName(name string) int {
	for i := 1; i <= 12; i++ {
		// Case-insensitive comparison
		if equalFold(persianMonthNames[i].English, name) {
			return i
		}
	}
	return 0
}

// equalFold is a simple case-insensitive string comparison
func equalFold(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		c1, c2 := s1[i], s2[i]
		// Convert to lowercase if uppercase ASCII
		if c1 >= 'A' && c1 <= 'Z' {
			c1 += 'a' - 'A'
		}
		if c2 >= 'A' && c2 <= 'Z' {
			c2 += 'a' - 'A'
		}
		if c1 != c2 {
			return false
		}
	}
	return true
}

// persianDigits maps ASCII digits to Persian digits
var persianDigits = map[rune]rune{
	'0': '۰',
	'1': '۱',
	'2': '۲',
	'3': '۳',
	'4': '۴',
	'5': '۵',
	'6': '۶',
	'7': '۷',
	'8': '۸',
	'9': '۹',
}

// latinDigits maps Persian digits to ASCII digits
var latinDigits = map[rune]rune{
	'۰': '0',
	'۱': '1',
	'۲': '2',
	'۳': '3',
	'۴': '4',
	'۵': '5',
	'۶': '6',
	'۷': '7',
	'۸': '8',
	'۹': '9',
}

// ToPersianDigits converts Latin digits to Persian digits
func ToPersianDigits(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if persian, ok := persianDigits[r]; ok {
			runes[i] = persian
		}
	}
	return string(runes)
}

// ToLatinDigits converts Persian digits to Latin digits
func ToLatinDigits(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if latin, ok := latinDigits[r]; ok {
			runes[i] = latin
		}
	}
	return string(runes)
}
