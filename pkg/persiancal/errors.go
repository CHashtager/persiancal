package persiancal

import "errors"

// Common errors
var (
	// ErrInvalidDate is returned when a date is invalid
	ErrInvalidDate = errors.New("invalid date")

	// ErrInvalidFormat is returned when a format string is invalid
	ErrInvalidFormat = errors.New("invalid format")

	// ErrInvalidLayout is returned when a layout string is invalid
	ErrInvalidLayout = errors.New("invalid layout")

	// ErrParseFailure is returned when parsing fails
	ErrParseFailure = errors.New("failed to parse date")

	// ErrInvalidMonth is returned when month is out of range (1-12)
	ErrInvalidMonth = errors.New("invalid month: must be between 1 and 12")

	// ErrInvalidDay is returned when day is out of range for the given month
	ErrInvalidDay = errors.New("invalid day for the given month")
)
