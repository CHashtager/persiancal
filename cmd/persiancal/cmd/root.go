package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "persiancal",
	Short: "Persian (Jalali) Calendar CLI tool",
	Long: `A command-line tool for working with the Persian (Jalali/Shamsi) calendar.
	
Features:
  - Convert between Gregorian and Jalali dates
  - Display current date in Jalali calendar
  - Calculate date differences
  - Format dates in various styles
  
Examples:
  persiancal now
  persiancal convert 2025-10-26
  persiancal diff 1403-01-01 1404-01-01`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("persian", "p", false, "Use Persian digits in output")
}
