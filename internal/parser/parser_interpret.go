package parser

import (
	"fmt"

	"github.com/MikhailoSafronov/Project-lab4/internal/ast"
	"github.com/MikhailoSafronov/Project-lab4/internal/lexer"
)

// Parse будує AST через рекурсивний спуск.
func Parse(l *lexer.Lexer) ([]ast.Node, error) {
	return parseSeq(l, false)
}

// parseSeq читає послідовність токенів, доки не зустріне потрібну “зупинку”.
// stopOnBracket true → вихід, коли натрапили на ']'.
func parseSeq(l *lexer.Lexer, stopOnBracket bool) ([]ast.Node, error) {
	var nodes []ast.Node

	for {
		tok := l.Next()

		switch tok.Kind {
		case lexer.GT, lexer.LT, lexer.PLUS, lexer.MINUS,
			lexer.DOT, lexer.COMMA:
			nodes = append(nodes, &ast.Command{Kind: tok.Kind})

		case lexer.LBR:
			body, err := parseSeq(l, true)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, &ast.Loop{Body: body})

		case lexer.RBR:
			if stopOnBracket {
				return nodes, nil
			}
			return nil, fmt.Errorf("unexpected ']' at offset %d", tok.Offset)

		case lexer.EOF:
			if stopOnBracket {
				return nil, fmt.Errorf("unmatched '[' (reached EOF)")
			}
			return nodes, nil

		case lexer.INVALID:
			return nil, fmt.Errorf("invalid character %q at offset %d", tok.Char, tok.Offset)

		default:
			return nil, fmt.Errorf("unknown token kind %d", tok.Kind)
		}
	}
}

// Interpret виконує обробку AST.
func Interpret(nodes []ast.Node) {
	const tapeSize = 30_000 // стандартна довжина “стрічки”
	tape := make([]byte, tapeSize)
	ptr := 0

	var exec func([]ast.Node)
	exec = func(seq []ast.Node) {
		for _, n := range seq {
			switch node := n.(type) {
			case *ast.Command:
				switch node.Kind {
				case lexer.GT:
					ptr = (ptr + 1) % tapeSize               // цикл по колу
				case lexer.LT:
					ptr = (ptr - 1 + tapeSize) % tapeSize    // цикл по колу
				case lexer.PLUS:
					tape[ptr]++
				case lexer.MINUS:
					tape[ptr]--
				case lexer.DOT:
					fmt.Printf("%c", tape[ptr])
				case lexer.COMMA:
					var input byte
					fmt.Scanf("%c", &input)
					tape[ptr] = input
				}

			case *ast.Loop:
				for tape[ptr] != 0 {
					exec(node.Body)
				}
			}
		}
	}

	exec(nodes)
}
