package main

func main() {
	reverseWords("hello world")           // Output: "world hello"
	reverseWords("  I love programming ") // Output: "programming love I"
	reverseWords("  the sky is blue  ")   // Output: "blue is sky the"
}

func reverseWords(s string) string {
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
