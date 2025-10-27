package persiancal

import "time"

// Constants for Gregorian and Jalali calendars
const (
	gregorianEpoch = 1721426 // Julian day number of Gregorian epoch (0001-01-01)
	jalaliEpoch    = 1948321 // Julian day number of Jalali epoch (0001-01-01)
)

// FromGregorian converts a Gregorian time.Time to Jalali date components.
// Returns year, month (1-12), and day (1-31).
func FromGregorian(t time.Time) (jy, jm, jd int) {
	gy := t.Year()
	gm := int(t.Month())
	gd := t.Day()

	gDays := gregorianToJDN(gy, gm, gd)
	jy, jm, jd = jdnToJalali(gDays)
	return
}

// ToGregorian converts a Jalali date to Gregorian time.Time.
// Month should be 1-12, day should be 1-31.
func ToGregorian(jy, jm, jd int) time.Time {
	jdn := jalaliToJDN(jy, jm, jd)
	gy, gm, gd := jdnToGregorian(jdn)
	return time.Date(gy, time.Month(gm), gd, 0, 0, 0, 0, time.UTC)
}

// gregorianToJDN converts a Gregorian date to Julian Day Number
func gregorianToJDN(gy, gm, gd int) int {
	// Adjust for months before March
	if gm < 3 {
		gy--
		gm += 12
	}

	a := gy / 100
	b := 2 - a + (a / 4)

	jdn := int(365.25*float64(gy+4716)) + int(30.6001*float64(gm+1)) + gd + b - 1524
	return jdn
}

// jdnToGregorian converts a Julian Day Number to Gregorian date
func jdnToGregorian(jdn int) (gy, gm, gd int) {
	a := jdn + 32044
	b := (4*a + 3) / 146097
	c := a - (146097*b)/4

	d := (4*c + 3) / 1461
	e := c - (1461*d)/4
	m := (5*e + 2) / 153

	gd = e - (153*m+2)/5 + 1
	gm = m + 3 - 12*(m/10)
	gy = 100*b + d - 4800 + m/10

	return
}

// jalaliToJDN converts a Jalali date to Julian Day Number
func jalaliToJDN(jy, jm, jd int) int {
	epbase := jy - 474
	epyear := 474 + (epbase % 2820)

	var mdays int
	if jm <= 7 {
		mdays = (jm - 1) * 31
	} else {
		mdays = (jm-1)*30 + 6
	}

	jdn := jd + mdays +
		(epyear*682-110)/2816 +
		(epyear-1)*365 +
		(epbase/2820)*1029983 +
		jalaliEpoch - 1

	return jdn
}

// jdnToJalali converts a Julian Day Number to Jalali date
func jdnToJalali(jdn int) (jy, jm, jd int) {
	// Calculate depoch (days since Jalali epoch)
	depoch := jdn - jalaliToJDN(475, 1, 1)

	// Calculate the 2820-year cycle
	cycle := depoch / 1029983
	cyear := depoch % 1029983

	if cyear < 0 {
		cyear += 1029983
		cycle--
	}

	// Handle the last day of the cycle
	if cyear == 1029982 {
		ycycle := 2820
		jy = ycycle + 2820*cycle + 474
		if jy <= 0 {
			jy--
		}
		jm = 12
		jd = 30
		return
	}

	// Calculate year within the cycle
	aux1 := cyear / 366
	aux2 := cyear % 366

	ycycle := (2816*aux2 + 2134*aux1 + 2815) / 1028522
	ycycle += aux1 + 1

	jy = ycycle + 2820*cycle + 474
	if jy <= 0 {
		jy--
	}

	// Calculate day of year
	yday := jdn - jalaliToJDN(jy, 1, 1) + 1

	// Determine month and day from day of year
	if yday <= 186 {
		// First 6 months (31 days each)
		jm = 1 + (yday-1)/31
		jd = ((yday - 1) % 31) + 1
	} else {
		// Last 6 months (30 days each, except last month)
		jm = 7 + (yday-187)/30
		jd = ((yday - 187) % 30) + 1
	}

	return
}

// isJalaliLeap checks if a Jalali year is a leap year using the 2820-year cycle
func isJalaliLeap(jy int) bool {
	// Algorithm based on 2820-year cycle
	epbase := jy - 474
	epyear := 474 + (epbase % 2820)

	return ((epyear*682 - 110) % 2816) < 682
}

// daysInJalaliMonth returns the number of days in a given Jalali month
func daysInJalaliMonth(jy, jm int) int {
	if jm < 1 || jm > 12 {
		return 0
	}
	if jm <= 6 {
		return 31
	}
	if jm <= 11 {
		return 30
	}
	// Month 12
	if isJalaliLeap(jy) {
		return 30
	}
	return 29
}
