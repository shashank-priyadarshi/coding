package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type Token struct {
	TokenType TokenType
	Text      string
}
type Element interface {
	Value() int
}
type Integer struct {
	value int
}

func (i Integer) Value() int {
	return i.value
}

type BinaryOperationType int

const (
	Addition BinaryOperationType = iota
	Subtraction
)

type BinaryOperation struct {
	Type        BinaryOperationType
	Left, Right Element
}

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Subtraction:
		return b.Left.Value() - b.Right.Value()
	default:
		panic("unsupported")
	}
}
func Parse(tokens []Token) Element {
	b := BinaryOperation{}
	var haveLhs bool
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].TokenType {
		case Int:
			n, _ := strconv.Atoi(tokens[i].Text)
			integer := &Integer{
				value: n,
			}
			if !haveLhs {
				b.Left = integer
				haveLhs = true
			} else {
				b.Right = integer
			}
		case Minus:
			b.Type = Subtraction
		case Plus:
			b.Type = Addition

		case Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].TokenType == Rparen {
					break
				}
			}
			var subExp []Token
			for k := i + 1; k < j; k++ {
				subExp = append(subExp, tokens[k])
			}
			element := Parse(subExp)
			if !haveLhs {
				b.Left = element
				haveLhs = true
			} else {
				b.Right = element
			}
			i = j
		}
	}
	return &b
}

func Lex(input string) []Token {
	var tokens []Token
	for i := 0; i < len(input); i++ {

		switch input[i] {
		case '+':
			tokens = append(tokens, Token{
				TokenType: Plus,
				Text:      "+",
			})
		case '-':
			tokens = append(tokens, Token{
				TokenType: Minus,
				Text:      "-",
			})
		case '(':
			tokens = append(tokens, Token{
				TokenType: Lparen,
				Text:      "(",
			})
		case ')':
			tokens = append(tokens, Token{
				TokenType: Rparen,
				Text:      ")",
			})
		default:
			sb := strings.Builder{}
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteByte(input[j])
					i++
				} else {
					i--
					tokens = append(tokens, Token{
						TokenType: Int,
						Text:      sb.String(),
					})
					break
				}
			}
		}

	}
	return tokens
}

func main() {
	input := "(12+10)-(20-18)"
	tokens := Lex(input)
	fmt.Printf("tokens: %v\n", tokens)
	element := Parse(tokens)
	fmt.Printf("%v: %v\n", input, element.Value())
}
