package lexer

// Kind описує тип BF-токена.
type Kind byte

const (
	GT    Kind = iota // '>'
	LT                // '<'
	PLUS              // '+'
	MINUS             // '-'
	DOT               // '.'
	COMMA             // ','
	LBR               // '['
	RBR               // ']'
	EOF
	INVALID // новий тип для некоректних символів
)

var mapByteToKind = map[byte]Kind{
	'>': GT, '<': LT, '+': PLUS, '-': MINUS,
	'.': DOT, ',': COMMA, '[': LBR, ']': RBR,
}

type Token struct {
	Kind   Kind
	Offset int
	Char   byte // оригінальний символ для відладки
}

type Lexer struct {
	src []byte
	pos int
}

func New(src []byte) *Lexer {
	return &Lexer{src: src}
}

// Next повертає наступний значущий токен, пропускаючи всі інші символи.
func (l *Lexer) Next() Token {
	for {
		if l.pos >= len(l.src) {
			return Token{Kind: EOF, Offset: l.pos}
		}
		ch := l.src[l.pos]
		offset := l.pos
		l.pos++
		if k, ok := mapByteToKind[ch]; ok {
			return Token{Kind: k, Offset: offset, Char: ch}
		}
		// будь-які інші символи трактується як коментар або недійсні
		return Token{Kind: INVALID, Offset: offset, Char: ch}
	}
}
