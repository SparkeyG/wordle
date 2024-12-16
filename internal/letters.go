package letters

import (
	"strings"
)

type Letter struct {
	IsExact     bool
	ExactLetter string
	LetterGuess []string
	ThisLetter  []string
}

func (l Letter) MakeRegexString() string {
	var this strings.Builder
	if l.IsExact {
		return l.ExactLetter
	}
	this.WriteString("[^")
	for _, char := range l.LetterGuess {
		this.WriteString(char)
	}
	this.WriteString("]")
	return this.String()
}
