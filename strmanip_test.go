package strmanip

import (
    "testing"
)

// TestIsNumeric calls strmanip.IsNumeric
func TestIsNumeric(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{"blank", "", false},
		{"edge of numeric", "0", true},
		{"numeric", "5", true},
		{"edge of numeric", "9", true},
		{"negative", "-1", true},
		{"not numeric", "a", false},
		{"2 numbers", "10", true},
		{"multiple numbers", "13579", true},
		{"multiple characters", "2k9199t", false},
		{"with decimal", "3.14159", true},
		{"with multiple decimals", "3.141.59", false},
		{"non alphanumeric", "%", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsNumeric(test.input)
			if actual != test.expected {
				t.Errorf("IsNumeric(%s) = %t; want %t", test.input, actual, test.expected)
			}
		})
	}
}

// TestIsLowercase calls strmanip.IsLowercase
func TestIsLowercase(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{"blank", "", false},
		{"edge of lowercase", "a", true},
		{"lowercase", "j", true},
		{"edge of lowercase", "z", true},
		{"uppercase", "A", false},
		{"numeric", "5", false},
		{"non alphanumeric", "%", false},
		{"2 characters", "az", true},
		{"multiple characters", "abcdefghijklmnopqrstuvwxyz", true},
		{"mixed with uppercase", "abcDefg", false},
		{"mixed with numeric", "abcd3fg", false},
		{"mixed", "zyx3Vu7", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsLowercase(test.input)
			if actual != test.expected {
				t.Errorf("IsLowercase(%s) = %t; want %t", test.input, actual, test.expected)
			}
		})
	}
}

// TestIsUppercase calls strmanip.IsUppercase
func TestIsUppercase(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{"blank", "", false},
		{"edge of uppercase", "A", true},
		{"uppercase", "G", true},
		{"edge of uppercase", "Z", true},
		{"lowercase", "a", false},
		{"numeric", "5", false},
		{"non alphanumeric", "%", false},
		{"2 characters", "AZ", true},
		{"multiple characters", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"mixed with lowercase", "ABCdEFG", false},
		{"mixed with numeric", "ABCD3FG", false},
		{"mixed", "ZYX3vU7", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsUppercase(test.input)
			if actual != test.expected {
				t.Errorf("IsUppercase(%s) = %t; want %t", test.input, actual, test.expected)
			}
		})
	}
}

// TestIsAlphaNumeric calls strmanip.IsAlphaNumeric
func TestIsAlphaNumeric(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{"blank", "", false},
		{"uppercase", "G", true},
		{"lowercase", "a", true},
		{"numeric", "5", true},
		{"2 characters", "AZ", true},
		{"multiple characters", "AbCdEfGhIjKl", true},
		{"mixed", "ZYX3vU7", true},
		{"non alphanumeric", "%", false},
		{"with decimal", "3.14159", false},
		{"with multiple decimals", "3.141.59", false},
		{"negative", "-3", false},
		{"negative with letters", "-3asdf132", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsAlphaNumeric(test.input)
			if actual != test.expected {
				t.Errorf("IsAlphaNumeric(%s) = %t; want %t", test.input, actual, test.expected)
			}
		})
	}
}

// TestToLowercase calls strmanip.ToLowercase
func TestToLowercase(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"blank", "", ""},
		{"single lowercase already", "a", "a"},
		{"all lowercase already", "abc", "abc"},
		{"one uppercase", "A", "a"},
		{"all uppercase", "AZ", "az"},
		{"mixed upper lower", "AbCdEfGhIjKl", "abcdefghijkl"},
		{"mixed with numbers", "ZyX3vu7", "zyx3vu7"},
		{"non alphanumeric", "%", "%"},
		{"non alphanumeric but has alphanumeric", "@USERNAME", "@username"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ToLowercase(test.input)
			if actual != test.expected {
				t.Errorf("ToLowercase(%s) = %s; want %s", test.input, actual, test.expected)
			}
		})
	}
}

// TestToUppercase calls strmanip.ToUppercase
func TestToUppercase(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"blank", "", ""},
		{"single uppercase already", "A", "A"},
		{"all uppercase already", "ABC", "ABC"},
		{"one lowercase", "a", "A"},
		{"all lowercase", "az", "AZ"},
		{"mixed upper lower", "AbCdEfGhIjKl", "ABCDEFGHIJKL"},
		{"mixed with numbers", "ZyX3vu7", "ZYX3VU7"},
		{"non alphanumeric", "%", "%"},
		{"non alphanumeric but has alphanumeric", "@username", "@USERNAME"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ToUppercase(test.input)
			if actual != test.expected {
				t.Errorf("ToUppercase(%s) = %s; want %s", test.input, actual, test.expected)
			}
		})
	}
}

// TestSplit calls strmanip.Split
func TestSplit(t *testing.T) {
	var tests = []struct {
		name      string
		inputA    string
		inputB    rune
		expected  []string
	}{
		{"empty", "", ',', []string{}},
		{"by space", "split by space", ' ', []string{"split", "by", "space"}},
		{"by letter", "beekeeper needed", 'e', []string{"b", "", "k", "", "p", "r n", "", "d", "d"}},
		{"filepath", "/home/./..//Documents/", '/', []string{"", "home", ".", "..", "", "Documents", ""}},
		{"same as split char", "A", 'A', []string{"", ""}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Split(test.inputA, test.inputB)
            if len(actual) != len(test.expected) {
				t.Errorf("Split(%s, %c) = %s; want %s", test.inputA, test.inputB, actual, test.expected)
            } else {
                for i, val := range test.expected {
                    if val != actual[i] {
                        t.Errorf("Split(%s, %c) = %s; want %s", test.inputA, test.inputB, actual, test.expected)
                    }
                }
            }
		})
	}
}

// TestJoin calls strmanip.Join
func TestJoin(t *testing.T) {
	var tests = []struct {
		name      string
		inputA    []string
		inputB    string
		expected  string
	}{
		{"empty", []string{}, "", ""},
		{"by space", []string{"join", "by", "space"}, " ", "join by space"},
		{"by letter", []string{"b", "", "k", "", "p", "r n", "", "d", "d!!"}, "ee", "beeeekeeeepeer neeeedeed!!" },
		{"filepath", []string{"", "home", ".", "..", "", "Documents", ""}, "/", "/home/./..//Documents/" },
		{"by blank", []string{"join", "by", "blank"}, "", "joinbyblank"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Join(test.inputA, test.inputB)
			if actual != test.expected {
				t.Errorf("Join(%v, %s) = %s; want %s", test.inputA, test.inputB, actual, test.expected)
			}
		})
	}
}

// TestIsPalindrome calls strmanip.IsPalindrome
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		name      string
		input     string
		expected  bool
	}{
		{"empty", "", true},
		{"even letter count", "abccba", true},
		{"odd letter count", "abcba", true},
		{"not palindrome", "abcde", false },
		{"close to but not palindrome even", "abcdecba", false },
		{"close to but not palindrome odd", "abcdeecba", false },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsPalindrome(test.input)
			if actual != test.expected {
				t.Errorf("IsPalindrome(%s) = %t; want %t", test.input, actual, test.expected)
			}
		})
	}
}

// TestStripNonAlphanumeric calls strmanip.StripNonAlphanumeric
func TestStripNonAlphanumeric(t *testing.T) {
	var tests = []struct {
		name      string
		input     string
		expected  string
	}{
		{"empty", "", ""},
		{"no non alphanumeric", "aBcDe", "aBcDe"},
		{"only non alphanumeric", "!@ #$", ""},
		{"combo", "\"I wish it need not have happened in my time,\" said Frodo. \"So do I,\" said Gandalf, \"and so do all who live to see such times. But that is not for them to decide. All we have to decide is what to do with the time that is given us.\"", "IwishitneednothavehappenedinmytimesaidFrodoSodoIsaidGandalfandsodoallwholivetoseesuchtimesButthatisnotforthemtodecideAllwehavetodecideiswhattodowiththetimethatisgivenus"},
    }

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := StripNonAlphanumeric(test.input)
			if actual != test.expected {
				t.Errorf("StripNonAlphanumeric(%s) = %s; want %s", test.input, actual, test.expected)
			}
		})
	}
}

func TestIsPalindromicSentence(t *testing.T) {
    var tests = []struct {
        name      string
        input     string
        expected  bool
    }{
		{"empty", "", true},
        {"palindrome", "Bob wondered, 'Now, Bob?'", true},
        {"not palindrome", "race a car", false},
		{"single letter", "a", true},
		{"palindrome various cases", "A man, a plan, a canal: Panama", true},
		{"palindrome", "Was it a car or a cat I saw?", true},
		{"not palindrome", "hello", false},
		{"only punctuation", ".,?!'", true},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            actual := IsPalindromicSentence(test.input)
            if actual != test.expected {
				t.Errorf("IsPalindromicSentence(%s) = %t; want %t", test.input, actual, test.expected)
            }
        })
    }
}
