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
	ASSIGN		= "="
	PLUS			= "+"
	MINUS			= "-"
	BANG			= "!"
	ASTERISK	= "*"
	FSLASH		= "/"
	LT				= "<"
	GT				= ">"

	// boolean logic
	EQ			= "=="
	NOT_EQ	= "!="

	// delimiters
	COMMA			= ","
	SEMICOLON	= ";"
	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"
	FN_RETURN	= "@"

	// keywords
	FUNCTION				= "FUNCTION"
	RETURN					= "RETURN"
	TRUE						= "TRUE"
	FALSE						= "FALSE"
	IF							= "IF"
	ELSE						= "ELSE"
	TYPE_INT				= "TYPE_INT"
	TYPE_VOID				= "TYPE_VOID"
)

var keywords = map[string]TokenType{
	// basic keywords
	"fn": FUNCTION,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,

	// types
	"int": TYPE_INT,
	"void": TYPE_VOID,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
