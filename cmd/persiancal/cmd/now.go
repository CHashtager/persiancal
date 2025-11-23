package cmd

import (
	"fmt"
	"time"

	"github.com/CHashtager/persiancal/pkg/persiancal"
	"github.com/spf13/cobra"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Display the current date in Jalali calendar",
	Long:  `Display the current date and time in the Persian (Jalali) calendar.`,
	Example: `  persiancal now
  persiancal now --format "yyyy/MM/dd"
  persiancal now --format "MMMM dd, yyyy"
  persiancal now --persian`,
	Run: runNow,
}

var (
	nowFormat      string
	nowShowTime    bool
	nowLongFormat  bool
	nowEnglishName bool
)

func init() {
	rootCmd.AddCommand(nowCmd)

	nowCmd.Flags().StringVarP(&nowFormat, "format", "f", "", "Custom format layout (e.g., 'yyyy/MM/dd')")
	nowCmd.Flags().BoolVarP(&nowShowTime, "time", "t", false, "Show time along with date")
	nowCmd.Flags().BoolVarP(&nowLongFormat, "long", "l", false, "Use long format with month name")
	nowCmd.Flags().BoolVarP(&nowEnglishName, "english", "e", false, "Use English month names (with --long)")
}

func runNow(cmd *cobra.Command, args []string) {
	usePersian, _ := cmd.Flags().GetBool("persian")

	now := time.Now()
	j := persiancal.FromGregorianDate(now)
	var output string

	if nowFormat != "" {
		if usePersian {
			output = j.FormatPersian(nowFormat)
		} else {
			output = j.Format(nowFormat)
		}
	} else if nowLongFormat {
		if nowEnglishName {
			output = j.Format("dd MMM yyyy")
		} else {
			output = j.Format("dd MMMM yyyy")
		}
		if usePersian && !nowEnglishName {
			output = persiancal.ToPersianDigits(output)
		}
	} else {
		// Default format: yyyy-MM-dd HH:mm (for now, just date)
		output = j.Format("yyyy-MM-dd")
		if usePersian {
			output = persiancal.ToPersianDigits(output)
		}
	}

	// Add time if requested
	if nowShowTime {
		timeStr := now.Format("15:04:05")
		if usePersian {
			timeStr = persiancal.ToPersianDigits(timeStr)
		}
		output += " " + timeStr
	}

	fmt.Println(output)
}
