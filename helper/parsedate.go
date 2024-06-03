
package helper

import (
    "time"
    "errors"
)

// ParseDate parses a date string in the format "DD-MM-YYYY" to a time.Time object
func ParseDate(dateStr string) (time.Time, error) {
    date, err := time.Parse("02-01-2006", dateStr)
    if err != nil {
        return time.Time{}, errors.New("invalid date format, expected DD-MM-YYYY")
    }
    return date, nil
}
