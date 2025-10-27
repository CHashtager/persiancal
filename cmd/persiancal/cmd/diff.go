package cmd

import (
	"fmt"
	"math"

	"github.com/CHashtager/persiancal/pkg/persiancal"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff [date1] [date2]",
	Short: "Calculate the difference between two Jalali dates",
	Long: `Calculate the difference between two dates in the Jalali calendar.
	
Both dates should be in Jalali format.

Supported input formats:
  - yyyy-MM-dd (e.g., 1404-08-04)
  - yyyy/MM/dd (e.g., 1404/08/04)
  - yyyy.MM.dd (e.g., 1404.08.04)`,
	Example: `  persiancal diff 1403-01-01 1404-01-01
  persiancal diff 1404-08-04 1404-09-10
  persiancal diff 1404-01-01 1404-12-29 --verbose
  persiancal diff 1403-01-01 1404-01-01 --persian`,
	Args: cobra.ExactArgs(2),
	RunE: runDiff,
}

var (
	diffVerbose  bool
	diffDaysOnly bool
)

func init() {
	rootCmd.AddCommand(diffCmd)

	diffCmd.Flags().BoolVarP(&diffVerbose, "verbose", "v", false, "Show detailed breakdown (years, months, days)")
	diffCmd.Flags().BoolVarP(&diffDaysOnly, "days-only", "d", false, "Show only the total number of days")
}

func runDiff(cmd *cobra.Command, args []string) error {
	usePersian, _ := cmd.Flags().GetBool("persian")

	date1Str := persiancal.ToLatinDigits(args[0])
	date2Str := persiancal.ToLatinDigits(args[1])

	j1, err := parseJalaliDate(date1Str)
	if err != nil {
		return fmt.Errorf("failed to parse first date: %w", err)
	}

	j2, err := parseJalaliDate(date2Str)
	if err != nil {
		return fmt.Errorf("failed to parse second date: %w", err)
	}

	// Calculate differences
	days := j2.DaysBetween(j1)
	absDays := int(math.Abs(float64(days)))

	if diffDaysOnly {
		output := fmt.Sprintf("%d", absDays)
		if usePersian {
			output = persiancal.ToPersianDigits(output)
		}
		fmt.Println(output)
		return nil
	}

	if !diffVerbose {
		output := fmt.Sprintf("%d days", absDays)
		if usePersian {
			output = persiancal.ToPersianDigits(output)
		}
		fmt.Println(output)
		return nil
	}

	var from, to persiancal.JalaliDate
	if days >= 0 {
		from = j1
		to = j2
	} else {
		from = j2
		to = j1
	}

	years := persiancal.YearsBetween(from, to)
	from = from.AddYears(years)

	months := persiancal.MonthsBetween(from, to)
	from = from.AddMonths(months)

	remainingDays := to.DaysBetween(from)

	var parts []string
	if years > 0 {
		yearStr := fmt.Sprintf("%d year", years)
		if years > 1 {
			yearStr += "s"
		}
		parts = append(parts, yearStr)
	}
	if months > 0 {
		monthStr := fmt.Sprintf("%d month", months)
		if months > 1 {
			monthStr += "s"
		}
		parts = append(parts, monthStr)
	}
	if remainingDays > 0 || len(parts) == 0 {
		dayStr := fmt.Sprintf("%d day", remainingDays)
		if remainingDays != 1 {
			dayStr += "s"
		}
		parts = append(parts, dayStr)
	}

	output := ""
	for i, part := range parts {
		if i > 0 {
			if i == len(parts)-1 {
				output += " and "
			} else {
				output += ", "
			}
		}
		output += part
	}

	if usePersian {
		output = persiancal.ToPersianDigits(output)
	}

	fmt.Printf("%s\n", output)
	fmt.Printf("(Total: %s)\n", func() string {
		s := fmt.Sprintf("%d days", absDays)
		if usePersian {
			return persiancal.ToPersianDigits(s)
		}
		return s
	}())

	return nil
}
