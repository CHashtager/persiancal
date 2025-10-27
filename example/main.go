package main

import (
	"fmt"
	"time"

	"github.com/CHashtager/persiancal/pkg/persiancal"
)

func main() {
	fmt.Println("PersianCal Library Example")
	fmt.Println()

	// Example 1: Get current date
	fmt.Println("1. Current Date:")
	j := persiancal.Now()
	fmt.Printf("   Today: %s\n", j)
	fmt.Printf("   Formatted: %s\n", j.Format("dd MMMM yyyy"))
	fmt.Printf("   Persian: %s\n\n", j.FormatPersian("yyyy/MM/dd"))

	// Example 2: Convert Gregorian to Jalali
	fmt.Println("2. Gregorian to Jalali Conversion:")
	g := time.Date(2025, 3, 20, 0, 0, 0, 0, time.UTC)
	j = persiancal.FromGregorianDate(g)
	fmt.Printf("   %s => %s (Nowruz!)\n\n", g.Format("2006-01-02"), j)

	// Example 3: Formatting options
	fmt.Println("3. Various Formatting Options:")
	j, _ = persiancal.Parse("yyyy/MM/dd", "1404/08/04")
	fmt.Printf("   ISO format: %s\n", j.Format(persiancal.LayoutISO))
	fmt.Printf("   Slash format: %s\n", j.Format(persiancal.LayoutSlash))
	fmt.Printf("   Long format: %s\n", j.Format(persiancal.LayoutLong))
	fmt.Printf("   English: %s\n", j.Format(persiancal.LayoutLongEnglish))
	fmt.Printf("   Custom: %s\n\n", j.Format("yy/M/d"))

	// Example 4: Parse dates
	fmt.Println("4. Parsing Dates:")
	j1, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/04")
	j2, _ := persiancal.Parse("yyyy-MM-dd", "1404-09-10")
	fmt.Printf("   Parsed: %s and %s\n\n", j1, j2)

	// Example 5: Date arithmetic
	fmt.Println("5. Date Arithmetic:")
	j = persiancal.Now()
	fmt.Printf("   Today: %s\n", j)
	fmt.Printf("   Next week: %s\n", j.AddDays(7))
	fmt.Printf("   Last month: %s\n", j.AddMonths(-1))
	fmt.Printf("   Next year: %s\n\n", j.AddYears(1))

	// Example 6: Calculate differences
	fmt.Println("6. Date Differences:")
	start, _ := persiancal.Parse("yyyy/MM/dd", "1404/01/01")
	end, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/06")
	fmt.Printf("   From %s to %s\n", start, end)
	fmt.Printf("   Days: %d\n", end.DaysBetween(start))
	fmt.Printf("   Months: %d\n", persiancal.MonthsBetween(start, end))
	fmt.Printf("   Duration: %v\n\n", end.Sub(start))

	// Example 7: Leap year detection
	fmt.Println("7. Leap Year Detection:")
	years := []int{1403, 1404, 1407, 1408}
	for _, year := range years {
		j := persiancal.JalaliDate{Year: year, Month: 12, Day: 29}
		fmt.Printf("   Year %d: %v (Esfand has %d days)\n",
			year, j.IsLeap(), persiancal.DaysInMonth(year, 12))
	}
	fmt.Println()

	// Example 8: Comparison
	fmt.Println("8. Date Comparison:")
	d1, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/04")
	d2, _ := persiancal.Parse("yyyy/MM/dd", "1404/08/10")
	fmt.Printf("   %s before %s: %v\n", d1, d2, d1.Before(d2))
	fmt.Printf("   %s after %s: %v\n", d1, d2, d1.After(d2))
	fmt.Printf("   %s equals %s: %v\n\n", d1, d2, d1.Equal(d2))

	// Example 9: Utility functions
	fmt.Println("9. Utility Functions:")
	j = persiancal.Now()
	fmt.Printf("   Day of week: %s\n", j.DayOfWeek())
	fmt.Printf("   Day of year: %d\n", j.DayOfYear())
	fmt.Printf("   Start of month: %s\n", j.StartOfMonth())
	fmt.Printf("   End of month: %s\n", j.EndOfMonth())
	fmt.Printf("   Month name (Persian): %s\n", j.MonthName())
	fmt.Printf("   Month name (English): %s\n\n", j.MonthNameEnglish())

	// Example 10: Convert back to Gregorian
	fmt.Println("10. Round-trip Conversion:")
	j, _ = persiancal.Parse("yyyy/MM/dd", "1404/08/04")
	g = j.ToGregorian()
	j2 = persiancal.FromGregorianDate(g)
	fmt.Printf("   Jalali: %s => Gregorian: %s => Jalali: %s\n",
		j, g.Format("2006-01-02"), j2)
	fmt.Printf("   Match: %v\n", j.Equal(j2))
}
