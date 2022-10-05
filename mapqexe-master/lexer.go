package mapq

import (
	_ "errors"
	"fmt"
	_ "fmt"
	"strings"
)

const (
	TYPE_PLUS      = iota // "+"0
	TYPE_SUB              // "-"0
	TYPE_MUL              // "*"0
	TYPE_DIV              // "/"0
	TYPE_LP               // "("0
	TYPE_RP               // ")"0
	TYPE_VAR              // "([a-z]|[A-Z])([a-z]|[A-Z]|[0-9])*"0
	TYPE_RES_TRUE         // "true"
	TYPE_RES_FALSE        // "false"
	TYPE_AND              // "&&"0
	TYPE_OR               // "||"
	TYPE_EQ               // "=="
	TYPE_LG               // ">"0
	TYPE_SM               // "<"0
	TYPE_LEQ              // ">="0
	TYPE_SEQ              // "<="0
	TYPE_NEQ              // "!="0
	TYPE_STR              // a quoted string(单引号)0
	TYPE_INT              // an integer
	TYPE_FLOAT            // 小数，x.y这种
	TYPE_UNKNOWN          // 未知的类型
	TYPE_NOT              // "!"0
	TYPE_DOT              // "."0
	TYPE_RES_NULL         // "null"
)

// Lexer 词法分析器
type Lexer struct {
	input string
	pos   int //开始是0
	runes []rune
}

var loc int

func (l *Lexer) Peek() (ch rune, end bool) {
	loc += 1
	return
}

// some finction maybe useful for your implementation
func isLetter(ch rune) bool {

	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}
func isLetterOrUnderscore(ch rune) bool {
	return isLetter(ch) || ch == '_' || ch == '-'
}
func isNum(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
func (l *Lexer) fin() bool {
	return loc >= len(l.runes)-1
}

func (l *Lexer) way() (code int, token string, eos bool) {
	for l.pos <= len(l.runes) {
		var char = l.runes[loc]
		switch char {
		case '(': //左括号
			code = TYPE_LP
			token = string('(')
			eos = true
			l.Peek()
			return
		case ')': //右括号
			code = TYPE_RP
			token = string(')')
			eos = true
			l.Peek()
			return
		case '+':
			code = TYPE_PLUS
			token = string('+')
			eos = true
			l.Peek()
			return
		case '-':
			code = TYPE_SUB
			token = string('-')
			eos = true
			l.Peek()
			return
		case '*':
			code = TYPE_MUL
			token = string('*')
			eos = true
			l.Peek()
			return
		case '/':
			code = TYPE_DIV
			token = string('/')
			eos = true
			l.Peek()
			return
		case '<':
			if l.fin() {
				l.Peek()
				char = l.runes[loc]
				switch char {
				case '=':
					code = TYPE_SEQ
					token = string("<=")
					eos = true
				default:
					code = TYPE_SM
					token = string('<')
					eos = true
				}
			} else {
				code = TYPE_SM
				token = string('<')
				eos = true
			}
			return

		case '>':
			if l.fin() {
				l.Peek()
				char = l.runes[loc]
				switch char {
				case '=':
					code = TYPE_LEQ
					token = string(">=")
					eos = true
				default:
					code = TYPE_LG
					token = string('>')
					eos = true
				}
			} else {
				code = TYPE_LG
				token = string('<')
				eos = true
			}
			return

		case '!':
			if l.fin() {
				l.Peek()
				char = l.runes[loc]
				switch char {
				case '=':
					code = TYPE_NEQ
					token = "!="
					eos = true
				default:
					code = TYPE_NOT
					token = "!"
					eos = true
				}
			} else {
				code = TYPE_NOT
				token = string('>')
				eos = true
			}
			return
		case '=':
			if l.fin() {
				l.Peek()
				char = l.runes[loc]
				if char == '=' {
					code = TYPE_EQ
					token = "=="
					eos = true
				}
			} else {
				code = TYPE_EQ
				token = string('>')
				eos = true
			}

			return
		case '&':
			l.Peek()
			char = l.runes[loc]
			if char == '&' {
				code = TYPE_AND
				token = "&&"
				eos = true
			}
			return
		case '|':
			l.Peek()
			char = l.runes[loc]
			if char == '&' {
				code = TYPE_OR
				token = "||"
				eos = true
			}
			return
		case '.':
			l.Peek()
			code = TYPE_DOT
			token = "."
			eos = true
			return
		case 39: // string

			l.Peek()
			char = l.runes[loc]
			for char != 39 {

				token += string(char)
				code = TYPE_STR
				eos = true

				l.Peek()
				char = l.runes[loc]
			}
			if char == 39 {

				return
			}
		default: // integers or identifiers
			if isLetterOrUnderscore(char) {
				for isLetterOrUnderscore(char) || isNum(char) {
					if l.fin() {
						token += string(char)
						if token == "true" {
							code = TYPE_RES_TRUE
							eos = true
							return
						} else if token == "false" {
							code = TYPE_RES_FALSE
							eos = true
							return
						} else {
							code = TYPE_VAR
							eos = true
							return
						}
					}

					token += string(char)
					code = TYPE_VAR
					eos = true
					l.Peek()
					char = l.runes[loc]
					if !isLetterOrUnderscore(char) && !isNum(char) {
						fmt.Println(loc)
						return
					}
				}

				return
			}

			if isNum(char) {
				for isNum(char) {
					if l.fin() {
						if !strings.Contains(token, ".") {
							code = TYPE_INT
							eos = true
						}
						token += string(char)
						return
					}
					token += string(char)
					l.Peek()
					char = l.runes[loc]
				}

				if char == '.' {
					code = TYPE_FLOAT
					eos = true
					token += string(char)
					l.Peek()
					char = l.runes[loc]
					for isNum(char) {
						if l.fin() {
							token += string(char)
							return
						}
						token += string(char)
						l.Peek()
						char = l.runes[loc]
					}
				}
			}
		}
	}
	return
}
