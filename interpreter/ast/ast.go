package ast

import "interpreter/token"

// variable types
type VarTypeInt struct {
	Token token.Token	// the token.VARTYPE_INT token
	Name *Identifier
	Value Expression
}
func (vti *VarTypeInt) statementNode() {}
func (vti *VarTypeInt) TokenLiteral() string { return vti.Token.Literal }

// abstract syntax tree nodes
type Node interface {
	TokenLiteral() string
}

type Statement interface{
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token	// the token.IDENT token
	Value string
}
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}	else {
		return ""
	}
}
