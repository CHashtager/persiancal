# 🗓️ PersianCal - Persian (Jalali/Shamsi) Calendar for Go

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.24-blue)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight, dependency-free Go library and CLI tool for working with the Persian (Jalali/Shamsi) calendar. Built to be idiomatic and feel native to Go, similar to the standard `time` package.

## ✨ Features

- 🔄 **Accurate Conversion**: Gregorian ↔ Jalali using precise 2820-year cycle algorithm
- 📅 **Rich Date API**: Comprehensive `JalaliDate` struct with intuitive methods
- 🎨 **Flexible Formatting**: Custom layouts with Persian/English month names and digits
- 🔍 **Smart Parsing**: Parse dates with Persian/Latin digits and month names
- ⚡ **Zero Dependencies**: Core library uses only Go standard library
- 🛠️ **CLI Tool**: Powerful command-line interface for date operations
- 🧪 **Well Tested**: Comprehensive test coverage (coming soon)
- 📦 **Easy Integration**: Simple API for bots, scripts, and applications

## 📦 Installation

### As a Library

```bash
go get github.com/CHashtager/persiancal/pkg/persiancal
```

### As a CLI Tool

```bash
go install github.com/CHashtager/persiancal/cmd/persiancal@latest
```

Or build from source:

```bash
git clone https://github.com/CHashtager/persiancal.git
cd persiancal
go build -o persiancal ./cmd/persiancal
```

## 🚀 Quick Start

### Library Usage

```go
package main

import (
    "fmt"
    "time"
    "github.com/CHashtager/persiancal/pkg/persiancal"
)

func main() {
    // Get current Jalali date
    j := persiancal.Now()
    fmt.Println("Today:", j) // 1404/08/04

    // Convert Gregorian to Jalali
    j = persiancal.FromGregorianDate(time.Date(2025, 10, 26, 0, 0, 0, 0, time.UTC))
    fmt.Println("Jalali:", j.Format("yyyy/MM/dd")) // 1404/08/04

    // Format with month names
    fmt.Println(j.Format("dd MMMM yyyy")) // 04 آبان 1404
    fmt.Println(j.Format("dd MMM yyyy"))  // 04 Aban 1404

    // Convert back to Gregorian
    g := j.ToGregorian()
    fmt.Println("Gregorian:", g.Format("2006-01-02")) // 2025-10-26

    // Date arithmetic
    nextWeek := j.AddDays(7)
    fmt.Println("Next week:", nextWeek) // 1404/08/11

    // Parse dates
    d, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/04")
    fmt.Println("Parsed:", d)

    // Calculate differences
    a, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/04")
    b, _ := persiancal.Parse("yyyy/MM/dd", "1404/09/10")
    fmt.Println("Days between:", b.DaysBetween(a)) // 38

    // Leap year detection
    fmt.Println("Is leap year?", j.IsLeap()) // false
}
```

### CLI Usage

```bash
# Show current date
$ persiancal now
1404-08-04

# Show current date with time
$ persiancal now --time
1404-08-04 17:55:30

# Show current date in long format
$ persiancal now --long
04 آبان 1404

# Custom format
$ persiancal now --format "MMMM dd, yyyy"
آبان 04, 1404

# Use Persian digits
$ persiancal now --persian
۱۴۰۴-۰۸-۰۴

# Convert Gregorian to Jalali
$ persiancal convert 2025-10-26
1404-08-04

# Convert with custom format
$ persiancal convert 2025-10-26 --format "dd MMMM yyyy"
04 آبان 1404

# Convert Jalali to Gregorian (reverse)
$ persiancal convert 1404-08-04 --reverse
2025-10-26

# Calculate date difference
$ persiancal diff 1403-01-01 1404-01-01
365 days

# Verbose difference
$ persiancal diff 1404-01-01 1404-08-04 --verbose
7 months and 3 days
(Total: 217 days)
```

## 📖 API Documentation

### Core Types

#### `JalaliDate`

```go
type JalaliDate struct {
    Year  int // Jalali year
    Month int // 1-12
    Day   int // 1-31
}
```

### Main Functions

#### Creating Dates

```go
// Get current date
j := persiancal.Now()

// Create from Gregorian
j := persiancal.FromGregorianDate(time.Now())

// Create with validation
j, err := persiancal.New(1404, 8, 4)

// Parse from string
j, err := persiancal.Parse("yyyy/MM/dd", "1404/08/04")
```

#### Conversion

```go
// To Gregorian
t := j.ToGregorian() // returns time.Time

// From Gregorian (low-level)
year, month, day := persiancal.FromGregorian(time.Now())
```

#### Formatting

```go
// Standard format
s := j.Format("yyyy/MM/dd") // 1404/08/04

// With month names
s := j.Format("dd MMMM yyyy") // 04 آبان 1404
s := j.Format("dd MMM yyyy")  // 04 Aban 1404

// Persian digits
s := j.FormatPersian("yyyy/MM/dd") // ۱۴۰۴/۰۸/۰۴

// Pre-defined layouts
s := j.Format(persiancal.LayoutISO)   // yyyy-MM-dd
s := j.Format(persiancal.LayoutSlash) // yyyy/MM/dd
s := j.Format(persiancal.LayoutLong)  // dd MMMM yyyy
```

#### Format Tokens

| Token  | Description                      | Example |
|--------|----------------------------------|---------|
| `yyyy` | 4-digit year                     | 1404    |
| `yy`   | 2-digit year                     | 04      |
| `MM`   | 2-digit month                    | 08      |
| `M`    | Month without leading zero       | 8       |
| `MMMM` | Persian month name               | آبان    |
| `MMM`  | English month name               | Aban    |
| `dd`   | 2-digit day                      | 04      |
| `d`    | Day without leading zero         | 4       |

#### Date Arithmetic

```go
// Add/subtract days
future := j.AddDays(7)
past := j.AddDays(-7)

// Add/subtract months
next := j.AddMonths(1)
prev := j.AddMonths(-1)

// Add/subtract years
nextYear := j.AddYears(1)

// Calculate differences
duration := j.Sub(other) // time.Duration
days := j.DaysBetween(other) // int
months := persiancal.MonthsBetween(j, other)
years := persiancal.YearsBetween(j, other)
```

#### Comparison

```go
if j.Before(other) { }
if j.After(other) { }
if j.Equal(other) { }
```

#### Utility Methods

```go
// Leap year
isLeap := j.IsLeap()

// Day of week
weekday := j.DayOfWeek() // time.Weekday

// Day of year
doy := j.DayOfYear() // 1-365/366

// Month/Year boundaries
start := j.StartOfMonth()
end := j.EndOfMonth()
yearStart := j.StartOfYear()
yearEnd := j.EndOfYear()

// Month names
persian := j.MonthName()        // آبان
english := j.MonthNameEnglish() // Aban

// Validation
err := j.Validate()
```

### Standalone Functions

```go
// Days in month
days := persiancal.DaysInMonth(1404, 8) // 30

// Days in year
days := persiancal.DaysInYear(1404) // 365

// Leap year check
isLeap := persiancal.IsLeapYear(1404) // false

// Digit conversion
persian := persiancal.ToPersianDigits("1404") // ۱۴۰۴
latin := persiancal.ToLatinDigits("۱۴۰۴")      // 1404

// Month name lookup
name := persiancal.GetMonthNamePersian(8)  // آبان
name := persiancal.GetMonthNameEnglish(8)  // Aban
month := persiancal.GetMonthFromPersianName("آبان") // 8
month := persiancal.GetMonthFromEnglishName("Aban")  // 8
```

## 📅 Persian Calendar Reference

### Month Names

| # | Persian   | English     | Days |
|---|-----------|-------------|------|
| 1 | فروردین   | Farvardin   | 31   |
| 2 | اردیبهشت  | Ordibehesht | 31   |
| 3 | خرداد     | Khordad     | 31   |
| 4 | تیر       | Tir         | 31   |
| 5 | مرداد     | Mordad      | 31   |
| 6 | شهریور    | Shahrivar   | 31   |
| 7 | مهر       | Mehr        | 30   |
| 8 | آبان      | Aban        | 30   |
| 9 | آذر       | Azar        | 30   |
| 10| دی        | Dey         | 30   |
| 11| بهمن      | Bahman      | 30   |
| 12| اسفند     | Esfand      | 29/30|

### Leap Years

The Persian calendar uses a sophisticated 2820-year cycle for determining leap years, making it one of the most accurate solar calendars. In a leap year, Esfand has 30 days instead of 29.

## 🧪 Examples

### Example 1: Birthday Calculator

```go
package main

import (
    "fmt"
    "github.com/CHashtager/persiancal/pkg/persiancal"
)

func main() {
    birthday, _ := persiancal.Parse("yyyy/MM/dd", "1370/05/15")
    today := persiancal.Now()
    
    age := persiancal.YearsBetween(birthday, today)
    days := today.DaysBetween(birthday)
    
    fmt.Printf("Age: %d years (%d days)\n", age, days)
    
    // Next birthday
    nextBirthday := persiancal.JalaliDate{
        Year:  today.Year,
        Month: birthday.Month,
        Day:   birthday.Day,
    }
    if nextBirthday.Before(today) {
        nextBirthday = nextBirthday.AddYears(1)
    }
    
    daysUntil := nextBirthday.DaysBetween(today)
    fmt.Printf("Days until next birthday: %d\n", daysUntil)
}
```

### Example 2: Date Range Validator

```go
package main

import (
    "fmt"
    "github.com/CHashtager/persiancal/pkg/persiancal"
)

func isInRange(date, start, end persiancal.JalaliDate) bool {
    return (date.Equal(start) || date.After(start)) && 
           (date.Equal(end) || date.Before(end))
}

func main() {
    start, _ := persiancal.Parse("yyyy/MM/dd", "1404/01/01")
    end, _ := persiancal.Parse("yyyy/MM/dd", "1404/12/29")
    check, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/04")
    
    if isInRange(check, start, end) {
        fmt.Println("Date is in range!")
    }
}
```

### Example 3: Working Days Calculator

```go
package main

import (
    "fmt"
    "time"
    "github.com/CHashtager/persiancal/pkg/persiancal"
)

func countWorkingDays(start, end persiancal.JalaliDate) int {
    count := 0
    current := start
    
    for current.Before(end) || current.Equal(end) {
        // Friday is weekend in Iran
        if current.DayOfWeek() != time.Friday {
            count++
        }
        current = current.AddDays(1)
    }
    
    return count
}

func main() {
    start, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/01")
    end, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/30")
    
    working := countWorkingDays(start, end)
    fmt.Printf("Working days in Aban: %d\n", working)
}
```

## 🔧 CLI Commands

### `persiancal now`

Display the current date in Jalali calendar.

**Flags:**
- `-f, --format`: Custom format layout
- `-t, --time`: Show time along with date
- `-l, --long`: Use long format with month name
- `-e, --english`: Use English month names
- `-p, --persian`: Use Persian digits (global flag)

**Examples:**
```bash
persiancal now
persiancal now --format "MMMM dd, yyyy"
persiancal now --long --persian
persiancal now --time
```

### `persiancal convert`

Convert between Gregorian and Jalali dates.

**Usage:** `persiancal convert [date]`

**Flags:**
- `-r, --reverse`: Convert from Jalali to Gregorian
- `-f, --format`: Output format layout
- `-p, --persian`: Use Persian digits (global flag)

**Examples:**
```bash
persiancal convert 2025-10-26
persiancal convert 1404-08-04 --reverse
persiancal convert 2025-10-26 --format "dd MMMM yyyy"
```

### `persiancal diff`

Calculate the difference between two Jalali dates.

**Usage:** `persiancal diff [date1] [date2]`

**Flags:**
- `-v, --verbose`: Show detailed breakdown (years, months, days)
- `-d, --days-only`: Show only the total number of days
- `-p, --persian`: Use Persian digits (global flag)

**Examples:**
```bash
persiancal diff 1403-01-01 1404-01-01
persiancal diff 1404-01-01 1404-08-04 --verbose
persiancal diff 1404-01-01 1404-12-29 --days-only
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Based on the astronomical algorithms for the Persian calendar
- Uses the accurate 2820-year cycle for leap year calculations
- Inspired by Go's `time` package design philosophy

## 📚 Additional Resources

- [Persian Calendar on Wikipedia](https://en.wikipedia.org/wiki/Solar_Hijri_calendar)
- [Astronomical Algorithms](https://en.wikipedia.org/wiki/Astronomical_algorithms)

## 🐛 Known Issues

None currently. Please report any issues on the GitHub issue tracker.

## 🗺️ Roadmap

- [x] Core conversion logic
- [x] JalaliDate struct and methods
- [x] Formatting and parsing
- [x] Date arithmetic
- [x] CLI tool
- [ ] Comprehensive test suite
- [ ] Benchmarks
- [ ] Additional calendar systems (Hijri)
- [ ] Timezone support
- [ ] JSON marshaling/unmarshaling


---

Made with ❤️ for the Persian-speaking Go community

