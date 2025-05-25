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
			return nil, fmt.Errorf("unexpected ']' at %d", tok.Offset)
		case lexer.EOF:
			if stopOnBracket {
				return nil, fmt.Errorf("unmatched '[' at EOF")
			}
			return nodes, nil
		case lexer.INVALID:
			return nil, fmt.Errorf("invalid character '%c' at %d", tok.Char, tok.Offset)
		default:
			return nil, fmt.Errorf("unknown token kind: %v", tok.Kind)
		}
	}
}

// Interpret виконує обробку AST
func Interpret(nodes []ast.Node) {
	const tapeSize = 30000
	tape := make([]byte, tapeSize)
	ptr := 0

	var exec func(nodes []ast.Node)
	exec = func(nodes []ast.Node) {
		for _, n := range nodes {
			switch node := n.(type) {
			case *ast.Command:
				switch node.Kind {
				case lexer.GT:
					ptr++
					if ptr >= tapeSize {
						ptr = 0 // wrap around
					}
				case lexer.LT:
					if ptr == 0 {
						ptr = tapeSize - 1
					} else {
						ptr--
					}
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
