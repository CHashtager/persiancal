package persiancal

// DaysInMonth returns the number of days in a given Jalali month and year
func DaysInMonth(year, month int) int {
	return daysInJalaliMonth(year, month)
}

// IsLeapYear checks if a Jalali year is a leap year
func IsLeapYear(year int) bool {
	return isJalaliLeap(year)
}

// DaysInYear returns the number of days in a Jalali year (365 or 366)
func DaysInYear(year int) int {
	if isJalaliLeap(year) {
		return 366
	}
	return 365
}

// MonthsBetween calculates the number of months between two dates
func MonthsBetween(from, to JalaliDate) int {
	months := (to.Year-from.Year)*12 + (to.Month - from.Month)

	// Adjust if the day hasn't been reached yet
	if to.Day < from.Day {
		months--
	}

	return months
}

// YearsBetween calculates the number of complete years between two dates
func YearsBetween(from, to JalaliDate) int {
	years := to.Year - from.Year

	// Adjust if the birthday hasn't occurred yet this year
	if to.Month < from.Month || (to.Month == from.Month && to.Day < from.Day) {
		years--
	}

	return years
}
