package strmanip

/*
import (
    "fmt"
)
*/

var decimal = '.'

func processIsNumeric(s string, allowDecimal bool, allowNegative bool) bool {
    length := len(s)
    if length == 0 {
        return false
    }

    foundDecimal := false
    for i, val := range s {
        if allowNegative && i == 0 && val == rune('-') {
            //the first value is "-" and we've allowed negative numbers, so continue
            continue
        } else if !allowNegative && i == 0 && val == rune('-') {
            //the first value is "-" but we've not allowed negative numbers
            return false
        }

        if allowDecimal && val == rune(decimal) {
            if !foundDecimal {
                //we've not found a decimal before, so set true & continue
                foundDecimal = true
                continue
            } else {
                //we've found a decimal before and now we have a second--not numeric
                return false
            }
        } else if !allowDecimal && val == rune(decimal) {
            //we're not allowing decimals but we've found one
            return false
        }

        if val < rune('0') || val > rune('9') {
            //outside the range of ascii numbers
            return false
        }
    }

    return true
}

func IsNumeric(s string) bool {
    return processIsNumeric(s, true, true)
}

func IsLowercase(s string) bool {
    length := len(s)
    if length == 0 {
        return false
    }

    for _, val := range s {
        if val < rune('a') || val > rune('z') {
            //outside the range of ascii lowercase letters
            return false
        }
    }

    return true
}

func IsUppercase(s string) bool {
    length := len(s)
    if length == 0 {
        return false
    }

    for _, val := range s {
        if val < rune('A') || val > rune('Z') {
            //outside the range of ascii lowercase letters
            return false
        }
    }

    return true
}

func IsAlphaNumeric(s string) bool {
    length := len(s)
    if length == 0 {
        return false
    }

    for _, val := range s {
        strVal := string(val)
        if !processIsNumeric(strVal, false, false) && !IsLowercase(strVal) && !IsUppercase(strVal) {
            //character is neither numeric without negative/decimal, lowercase, nor uppercase,
            //therefore the string as a whole is not alphanumeric
            return false
        }
    }

    return true
}

func ToLowercase(s string) string {
    length := len(s)
    if length == 0 {
        return s
    }

    difference := rune('a') - rune('A')

    runes := []rune(s)

    for i, val := range s {
        strVal := string(val)

        if IsUppercase(strVal) {
            val += difference
        }
        runes[i] = val
    }

    return string(runes)
}

func ToUppercase(s string) string {
    length := len(s)
    if length == 0 {
        return s
    }

    difference := rune('a') - rune('A')

    runes := []rune(s)

    for i, val := range s {
        strVal := string(val)

        if IsLowercase(strVal) {
            val -= difference
        }
        runes[i] = val
    }

    return string(runes)
}
