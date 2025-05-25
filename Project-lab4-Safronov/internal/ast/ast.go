package ast

import "github.com/MikhailoSafronov/Project-lab4/internal/lexer"

// Node — базовий інтерфейс.
type Node interface{}

// Command — одинична BF-команда.
type Command struct{ Kind lexer.Kind }

// Loop — набір вузлів усередині [... ].
type Loop struct{ Body []Node }
