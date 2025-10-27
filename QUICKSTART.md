# 🚀 Quick Start Guide

## Installation

### Install CLI Tool

```bash
go install github.com/CHashtager/persiancal/cmd/persiancal@latest
```

### Use as Library

```bash
go get github.com/CHashtager/persiancal/pkg/persiancal
```

## 5-Minute Tutorial

### CLI Usage

```bash
# Current date in Jalali calendar
$ persiancal now
1404-08-06

# With Persian digits
$ persiancal now --persian
۱۴۰۴-۰۸-۰۶

# Long format
$ persiancal now --long
06 آبان 1404

# Convert Gregorian to Jalali
$ persiancal convert 2025-03-20
1404-01-01

# Convert with custom format
$ persiancal convert 2025-03-20 --format "dd MMMM yyyy"
01 فروردین 1404

# Reverse conversion (Jalali to Gregorian)
$ persiancal convert 1404-01-01 --reverse
2025-03-20

# Calculate date difference
$ persiancal diff 1403-01-01 1404-01-01
365 days

# Detailed difference
$ persiancal diff 1403-01-01 1404-08-06 --verbose
1 year, 7 months and 5 days
(Total: 586 days)
```

### Library Usage

```go
package main

import (
    "fmt"
    "github.com/CHashtager/persiancal/pkg/persiancal"
)

func main() {
    // Get current date
    today := persiancal.Now()
    fmt.Println(today) // 1404/08/06

    // Format
    fmt.Println(today.Format("dd MMMM yyyy")) // 06 آبان 1404

    // Parse
    date, _ := persiancal.Parse("yyyy/MM/dd", "1404/01/01")
    
    // Date arithmetic
    nextWeek := today.AddDays(7)
    
    // Calculate difference
    days := nextWeek.DaysBetween(today)
    fmt.Println(days) // 7
}
```

## Common Use Cases

### Birthday Calculator

```go
birthday, _ := persiancal.Parse("yyyy/MM/dd", "1370/05/15")
age := persiancal.YearsBetween(birthday, persiancal.Now())
fmt.Printf("Age: %d years\n", age)
```

### Date Validation

```go
date, err := persiancal.New(1404, 12, 30)
if err != nil {
    fmt.Println("Invalid date:", err)
}
```

### Format Conversion

```go
// Gregorian to Jalali
g := time.Now()
j := persiancal.FromGregorianDate(g)
fmt.Println(j.Format("yyyy/MM/dd"))

// Jalali to Gregorian
j, _ = persiancal.Parse("yyyy/MM/dd", "1404/08/06")
g = j.ToGregorian()
fmt.Println(g.Format("2006-01-02"))
```

## Next Steps

- Read the full [README.md](README.md) for detailed documentation
- Check out [example/main.go](example/main.go) for more examples
- See [CONTRIBUTING.md](CONTRIBUTING.md) to contribute

## Get Help

```bash
persiancal --help
persiancal now --help
persiancal convert --help
persiancal diff --help
```

## Feedback

Found a bug? Have a feature request? [Open an issue](https://github.com/CHashtager/persiancal/issues)!

