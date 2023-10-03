package utils

import "fmt"

func ToDecimalString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
