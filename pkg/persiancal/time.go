package persiancal

import (
	"fmt"
	"time"
)

// JalaliDate represents a date in the Jalali (Persian) calendar
type JalaliDate struct {
	Year  int
	Month int // 1-12
	Day   int // 1-31
}

// Now returns the current date in the Jalali calendar
func Now() JalaliDate {
	return FromGregorianDate(time.Now())
}

// FromGregorianDate converts a Gregorian time.Time to JalaliDate
func FromGregorianDate(t time.Time) JalaliDate {
	y, m, d := FromGregorian(t)
	return JalaliDate{Year: y, Month: m, Day: d}
}

// New creates a new JalaliDate with validation
func New(year, month, day int) (JalaliDate, error) {
	j := JalaliDate{Year: year, Month: month, Day: day}
	if err := j.Validate(); err != nil {
		return JalaliDate{}, err
	}
	return j, nil
}

// ToGregorian converts the JalaliDate to a Gregorian time.Time
func (j JalaliDate) ToGregorian() time.Time {
	return ToGregorian(j.Year, j.Month, j.Day)
}

// IsLeap returns true if the year is a leap year in the Jalali calendar
func (j JalaliDate) IsLeap() bool {
	return isJalaliLeap(j.Year)
}

// String returns a string representation of the date in yyyy/MM/dd format
func (j JalaliDate) String() string {
	return fmt.Sprintf("%04d/%02d/%02d", j.Year, j.Month, j.Day)
}

// AddDays adds n days to the date and returns a new JalaliDate
func (j JalaliDate) AddDays(n int) JalaliDate {
	// Convert to Gregorian, add days, convert back
	t := j.ToGregorian()
	t = t.AddDate(0, 0, n)
	return FromGregorianDate(t)
}

// AddMonths adds n months to the date and returns a new JalaliDate
func (j JalaliDate) AddMonths(n int) JalaliDate {
	year := j.Year
	month := j.Month + n
	day := j.Day

	// Handle month overflow/underflow
	for month > 12 {
		month -= 12
		year++
	}
	for month < 1 {
		month += 12
		year--
	}

	// Adjust day if it exceeds the number of days in the new month
	maxDay := daysInJalaliMonth(year, month)
	if day > maxDay {
		day = maxDay
	}

	return JalaliDate{Year: year, Month: month, Day: day}
}

// AddYears adds n years to the date and returns a new JalaliDate
func (j JalaliDate) AddYears(n int) JalaliDate {
	year := j.Year + n
	month := j.Month
	day := j.Day

	// Adjust day if it's Feb 30 in a leap year becoming non-leap
	maxDay := daysInJalaliMonth(year, month)
	if day > maxDay {
		day = maxDay
	}

	return JalaliDate{Year: year, Month: month, Day: day}
}

// Sub returns the duration between two JalaliDates
func (j JalaliDate) Sub(other JalaliDate) time.Duration {
	t1 := j.ToGregorian()
	t2 := other.ToGregorian()
	return t1.Sub(t2)
}

// DaysBetween returns the number of days between two JalaliDates
func (j JalaliDate) DaysBetween(other JalaliDate) int {
	duration := j.Sub(other)
	days := int(duration.Hours() / 24)
	return days
}

// Before returns true if j is before other
func (j JalaliDate) Before(other JalaliDate) bool {
	if j.Year != other.Year {
		return j.Year < other.Year
	}
	if j.Month != other.Month {
		return j.Month < other.Month
	}
	return j.Day < other.Day
}

// After returns true if j is after other
func (j JalaliDate) After(other JalaliDate) bool {
	if j.Year != other.Year {
		return j.Year > other.Year
	}
	if j.Month != other.Month {
		return j.Month > other.Month
	}
	return j.Day > other.Day
}

// Equal returns true if j equals other
func (j JalaliDate) Equal(other JalaliDate) bool {
	return j.Year == other.Year && j.Month == other.Month && j.Day == other.Day
}

// Validate checks if the JalaliDate is valid
func (j JalaliDate) Validate() error {
	if j.Month < 1 || j.Month > 12 {
		return ErrInvalidMonth
	}

	maxDay := daysInJalaliMonth(j.Year, j.Month)
	if j.Day < 1 || j.Day > maxDay {
		return ErrInvalidDay
	}

	return nil
}

// DayOfWeek returns the day of the week (0 = Sunday, 6 = Saturday)
func (j JalaliDate) DayOfWeek() time.Weekday {
	return j.ToGregorian().Weekday()
}

// DayOfYear returns the day of the year (1-365 or 1-366)
func (j JalaliDate) DayOfYear() int {
	days := 0
	for m := 1; m < j.Month; m++ {
		days += daysInJalaliMonth(j.Year, m)
	}
	days += j.Day
	return days
}

// WeekNumber returns the ISO week number (1-53)
func (j JalaliDate) WeekNumber() int {
	// Calculate based on Gregorian conversion
	_, week := j.ToGregorian().ISOWeek()
	return week
}

// StartOfMonth returns the first day of the month
func (j JalaliDate) StartOfMonth() JalaliDate {
	return JalaliDate{Year: j.Year, Month: j.Month, Day: 1}
}

// EndOfMonth returns the last day of the month
func (j JalaliDate) EndOfMonth() JalaliDate {
	maxDay := daysInJalaliMonth(j.Year, j.Month)
	return JalaliDate{Year: j.Year, Month: j.Month, Day: maxDay}
}

// StartOfYear returns the first day of the year (1 Farvardin)
func (j JalaliDate) StartOfYear() JalaliDate {
	return JalaliDate{Year: j.Year, Month: 1, Day: 1}
}

// EndOfYear returns the last day of the year (29 or 30 Esfand)
func (j JalaliDate) EndOfYear() JalaliDate {
	maxDay := daysInJalaliMonth(j.Year, 12)
	return JalaliDate{Year: j.Year, Month: 12, Day: maxDay}
}

// MonthName returns the Persian name of the month
func (j JalaliDate) MonthName() string {
	return GetMonthNamePersian(j.Month)
}

// MonthNameEnglish returns the English transliteration of the month name
func (j JalaliDate) MonthNameEnglish() string {
	return GetMonthNameEnglish(j.Month)
}
