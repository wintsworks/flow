package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Token string

// enums
const (
	MathOp     Token = "+-*/"
	Identifier Token = "ID"
	Number     Token = "NUM"
	OpenBrac   Token = "{"
	CloseBrac  Token = "}"
	LeftPar    Token = "("
	RightPar   Token = ")"
)

// Lexer tokenizes an input string.
type Lexer struct {
	input *bytes.Buffer
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: bytes.NewBufferString(input)}
	return lexer
}

// Tokenize returns the next token/input buffer, or an empty string if none left.
func (lex *Lexer) Tokenize() Token {
	ch, _, err := lex.input.ReadRune()
	if err != nil {
		if err == io.EOF {
			return "" // End of input
		}
		return "" // Error
	}

	switch ch {
	case ' ': // First Delimiter
		return lex.Tokenize() // Skip whitespace
	case ';': // Delimeter // issue with double white space.
		return lex.Tokenize() // Skip whitespace
	case '(':
		return LeftPar
	case ')':
		return RightPar
	case '{':
		return OpenBrac
	case '}':
		return CloseBrac
	case '+':
		return Token(string(ch))
	case '-':
		return Token(string(ch))
	case '*':
		return Token(string(ch))
	case '/':
		return Token(string(ch))
	default: // tokens that are not identifiers or scalars.
		if is_letter(ch) {
			lex.input.UnreadRune()
			return lex.readIdentifier()
		} else if is_digit(ch) {
			lex.input.UnreadRune()
			return lex.readNumber()
		}
		return Token(string(ch))
	}
}

// ReadIdentifier reads an identifier from the input buffer.
func (lex *Lexer) readIdentifier() Token {
	var buffer bytes.Buffer
	for {
		ch, _, err := lex.input.ReadRune()
		if err != nil {
			break
		}
		if !is_letter(ch) && !is_digit(ch) {
			lex.input.UnreadRune()
			break
		}
		buffer.WriteRune(ch)
	}
	return Token(buffer.String())
}

// ReadNumber reads a number from the input buffer.
func (lex *Lexer) readNumber() Token {
	var buffer bytes.Buffer
	for {
		ch, _, err := lex.input.ReadRune()
		if err != nil {
			break
		}
		if !is_digit(ch) {
			lex.input.UnreadRune()
			break
		}
		buffer.WriteRune(ch)
	}
	numberStr := buffer.String()
	if _, err := strconv.Atoi(numberStr); err == nil {
		return Token(numberStr)
	}
	return Token(buffer.String())
}

// Need to implement a scalar method.

// checks if the character is a letter or a beginning of an identifier.
func is_letter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '$'
}

// checks if the character is a digit or a beginning of a number.
func is_digit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

// entry point
func main() {
	input := "$test_var + $varTest * 10;"
	lexer := NewLexer(input)

	var tokens []Token

	for {
		token := lexer.Tokenize()
		if token == "" {
			break
		}
		tokens = append(tokens, token)
		fmt.Printf("%s ", token)
	}

	fmt.Println("\n", tokens)
}
