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
            //outside the range of ASCII numbers
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
            //outside the range of ASCII lowercase letters
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
            //outside the range of ASCII lowercase letters
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

    //find the difference between the ASCII values of lowercase and uppercase
    //this will be what we need to add to the uppercase value to make lowercase
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

    //find the difference between the ASCII values of lowercase and uppercase
    //this will be what we need to subtract from the lowercase value to make uppercase
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

func Split(s string, c rune) []string {
    sLength := len(s)
    result := []string{}

    //if string is empty, return an empty slice
    if sLength == 0 {
        return result
    }

    //create a rune slice to process each letter into a word
    word := []rune{}

    //loop through the string
    for _, val := range s {
        if val == c {
            //we've found the split character, so add the current word to the result slice
            result = append(result, string(word))
            //reset the word for the next set of processing
            word = word[:0]
        } else {
            //we've not found the split character yet, so add this rune to the word
            word = append(word, val)
        }
    }

    //whatever word is left at the end of the string gets added to the result
    result = append(result, string(word))

    return result
}

func Join(arr []string, s string) string {
    arrLength := len(arr)

    //if we're given an empty slice, return an empty string
    if arrLength == 0 {
        return ""
    }

    //build the result one character at a time
    result := []rune{}
    for i, val := range arr {
        //loop over the characters in each string in arr & add to the result
        for _, r := range val {
            result = append(result, r)
        }

        //we add the joining string BETWEEN the words in arr, no need to add at the end
        if (i < arrLength - 1) {
            //loop over the characters in the joining string & add to the result
            for _, r := range s {
                result = append(result, r)
            }
        }
    }

    //return the rune result slice as a string
    return string(result)
}
