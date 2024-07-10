package token

/* the token type could be later changed to an int or a byte.
doing so would provide a lot of performance gains, however,
it would also be more difficult to debug and needs more boilerplate. */
type TokenType string

type Token struct {
	Type TokenType
	Literal string
	/* TODO: consider attaching filenames/line numbers/etc.
	this will make it easier to debug production code.
	note that it adds a lot of complexity,
	but will be needed prior to a 1.0 launch */
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
	FUNCTION				= "FUNCTION"
	RETURN					= "RETURN"
	VARTYPE_INT			= "VARTYPE_INT"
	RETTYPE_INT			= "RETTYPE_INT"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"@int": RETTYPE_INT,
	"int": VARTYPE_INT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
