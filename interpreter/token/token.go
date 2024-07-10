package token

/* the token type could be later changed to an int or a byte.
doing so would provide a lot of performance gains, however,
it would also be more difficult to debug and needs more boilerplate. */
type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	// specials
	ILLEGAL = "ILLEGAL"
	EOF			= "EOF"

	// identifiers and literals
	IDENT	= "IDENT" // add, foobar, x, y, my_var, etc
	INT		= "INT"		// 460, 69, 8675309, etc

	// operators
	ASSIGN	= "="
	PLUS		= "+"

	// delimiters
	COMMA			= ","
	SEMICOLON	= ";"
	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"

	// keywords
	FUNCTION			= "FUNCTION"
	TYPE_INTEGER	= "TYPE_INTEGER"
)
