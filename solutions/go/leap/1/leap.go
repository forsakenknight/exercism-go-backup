package leap

// IsLeapYear calculates if the given year is leap or not
func IsLeapYear(year int) bool {
	if year%100 == 0 {
        return year%400 == 0
    }

    return year%4 == 0
}
