package passwords

// ValidPasswords finds all passwords in the range that match all the rules
func ValidPasswords(min, max int) (result []int) {
	for password := min; password < max; password++ {
		if ValidPassword(password) {
			result = append(result, password)
		}
	}

	return
}

// ValidPassword checks the rules (a double, no descending) and returns whether
// the given password is valid
func ValidPassword(password int) bool {
	if noDoubles(password) {
		return false
	}

	if anyDescending(password) {
		return false
	}

	return true
}

func noDoubles(password int) bool {
	for index := 0; index < 5; index++ {
		digit := (password / Pow(10, index)) % 10
		prevDigit := (password / Pow(10, index+1)) % 10

		if digit == prevDigit {
			return false
		}
	}

	return true
}

func anyDescending(password int) bool {
	for index := 0; index < 5; index++ {
		digit := (password / Pow(10, index)) % 10
		prevDigit := (password / Pow(10, index+1)) % 10

		if digit < prevDigit {
			return true
		}
	}

	return false
}

func Pow(number, exponent int) int {
	result := 1

	for index := 0; index < exponent; index++ {
		result *= number
	}

	return result
}
