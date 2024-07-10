package lexer

type Lexer struct {
	input string
	position int			// current position in input (points to current char)
	readPosition int	// current reading position in input (after current char)
	ch byte						// current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/* gives the next character and advance our position in the input string.
check whether we reach the end of input, if so set l.ch to 0, otherwise set l.ch to the next character then increment by one.

with the current implementation, Omega only supports ASCII instead of full unicode.
in order to fix this, l.ch needs to be a rune instead of a byte since characters could be multiple bytes wide, and then other functions would need to change to accomodate this.*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	}	else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
