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
		}
	}
}
