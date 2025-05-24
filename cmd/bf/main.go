package main

import (
	"log"
	"os"

	"github.com/MikhailoSafronov/Project-lab4/internal/lexer"
	"github.com/MikhailoSafronov/Project-lab4/internal/parser"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: bf <file.bf>")
	}

	filename := os.Args[1]
	src, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("read: %v", err)
	}

	lex := lexer.New(src)
	if _, err := parser.Parse(lex); err != nil {
		log.Fatalf("syntax: %v", err)
	}

	log.Println("lexing + parsing OK") // тимчасове підтвердження
}
