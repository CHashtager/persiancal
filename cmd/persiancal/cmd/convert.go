package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/CHashtager/persiancal/pkg/persiancal"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert [date]",
	Short: "Convert between Gregorian and Jalali dates",
	Long: `Convert a date between Gregorian and Jalali calendars.
	
By default, converts Gregorian to Jalali.
Use --reverse to convert Jalali to Gregorian.

Supported input formats:
  - yyyy-MM-dd (e.g., 2025-10-26)
  - yyyy/MM/dd (e.g., 2025/10/26)
  - yyyy.MM.dd (e.g., 2025.10.26)`,
	Example: `  persiancal convert 2025-10-26
  persiancal convert 1404-08-04 --reverse
  persiancal convert 2025-10-26 --format "MMMM dd, yyyy"
  persiancal convert 2025-10-26 --persian`,
	Args: cobra.ExactArgs(1),
	RunE: runConvert,
}

var (
	convertReverse bool
	convertFormat  string
)

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().BoolVarP(&convertReverse, "reverse", "r", false, "Convert from Jalali to Gregorian")
	convertCmd.Flags().StringVarP(&convertFormat, "format", "f", "", "Output format layout")
}

func runConvert(cmd *cobra.Command, args []string) error {
	usePersian, _ := cmd.Flags().GetBool("persian")
	dateStr := args[0]

	dateStr = persiancal.ToLatinDigits(dateStr)

	if convertReverse {
		j, err := parseJalaliDate(dateStr)
		if err != nil {
			return fmt.Errorf("failed to parse Jalali date: %w", err)
		}

		g := j.ToGregorian()

		var output string
		if convertFormat != "" {
			output = g.Format(convertFormat)
		} else {
			output = g.Format("2006-01-02")
		}

		if usePersian {
			output = persiancal.ToPersianDigits(output)
		}

		fmt.Println(output)
	} else {
		g, err := parseGregorianDate(dateStr)
		if err != nil {
			return fmt.Errorf("failed to parse Gregorian date: %w", err)
		}

		j := persiancal.FromGregorianDate(g)

		var output string
		if convertFormat != "" {
			if usePersian {
				output = j.FormatPersian(convertFormat)
			} else {
				output = j.Format(convertFormat)
			}
		} else {
			output = j.Format("yyyy-MM-dd")
			if usePersian {
				output = persiancal.ToPersianDigits(output)
			}
		}

		fmt.Println(output)
	}

	return nil
}

func parseJalaliDate(dateStr string) (persiancal.JalaliDate, error) {
	separators := []string{"-", "/", "."}

	for _, sep := range separators {
		if strings.Contains(dateStr, sep) {
			parts := strings.Split(dateStr, sep)
			if len(parts) != 3 {
				continue
			}

			year, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}

			month, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}

			day, err := strconv.Atoi(parts[2])
			if err != nil {
				continue
			}

			j, err := persiancal.New(year, month, day)
			if err != nil {
				return persiancal.JalaliDate{}, err
			}

			return j, nil
		}
	}

	return persiancal.JalaliDate{}, fmt.Errorf("unsupported date format: %s", dateStr)
}

func parseGregorianDate(dateStr string) (time.Time, error) {
	formats := []string{
		"2006-01-02",
		"2006/01/02",
		"2006.01.02",
		"02-01-2006",
		"02/01/2006",
		"02.01.2006",
	}

	for _, format := range formats {
		t, err := time.Parse(format, dateStr)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
}
