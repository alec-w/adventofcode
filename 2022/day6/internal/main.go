package internal

import "strings"

func CharactersBeforeXDistinct(input string, distinctCharsRequired int) int {
	for i := 0; ; i++ {
		chars := input[i : i+distinctCharsRequired]
		count := 0
		for i := 0; i < distinctCharsRequired; i++ {
			count += strings.Count(chars, string(chars[i]))
		}
		if count == distinctCharsRequired {
			return i + distinctCharsRequired
		}
	}
}
