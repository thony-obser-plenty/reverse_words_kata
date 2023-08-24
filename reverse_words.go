package main

import "fmt"

func main() {
	fmt.Println(reverseString("hello world"))           // Output: "world hello"
	fmt.Println(reverseString("  I love programming ")) // Output: "programming love I"
	fmt.Println(reverseString("  the sky is blue  "))   // Output: "blue is sky the"

	fmt.Println(reverseWords("hello world"))           // Output: "world hello"
	fmt.Println(reverseWords("  I love programming ")) // Output: "programming love I"
	fmt.Println(reverseWords("  the sky is blue  "))   // Output: "blue is sky the"
}

// Well i fucked up, the challenge was to write a function to reverse words, not strings.
func reverseString(s string) string {
	trimSpaces(&s)

	reversedString := ""
	for characterIndex := len(s) - 1; characterIndex >= 0; characterIndex-- {
		reversedString += string(s[characterIndex])
	}

	return reversedString
}

func trimSpaces(s *string) {
	trimLeadingWhitespace(s)
	trimTrailingWhitespaces(s)
}

func trimTrailingWhitespaces(s *string) {
	str := *s

	var stringOffset int
	for stringOffset = len(str) - 1; stringOffset >= 0; stringOffset-- {
		if !isWhitespace(rune(str[stringOffset])) {
			break
		}
	}

	*s = str[:stringOffset+1]
}

func trimLeadingWhitespace(s *string) {
	str := *s

	var stringOffset int
	for stringOffset = 0; stringOffset < len(str); stringOffset++ {
		if !isWhitespace(rune(str[stringOffset])) {
			break
		}
	}

	*s = str[stringOffset:]
}

func isWhitespace(ch rune) bool {
	whitespaceCharacters := []rune{' ', '\t', '\n'}

	for _, w := range whitespaceCharacters {
		if ch == w {
			return true
		}
	}

	return false
}

func reverseWords(sentence string) string {
	words := make([]string, 0)
	currentWord := ""

	for _, char := range sentence {
		if char == ' ' {
			if currentWord != "" {
				words = append([]string{currentWord}, words...)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		words = append([]string{currentWord}, words...)
	}

	reversedSentence := joinWords(words, " ")

	return reversedSentence
}

func joinWords(words []string, separator string) string {
	length := len(words)
	if length == 0 {
		return ""
	} else if length == 1 {
		return words[0]
	}

	reversedSentence := words[0]
	for i := 1; i < length; i++ {
		reversedSentence += separator + words[i]
	}

	return reversedSentence
}
