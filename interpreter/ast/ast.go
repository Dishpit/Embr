package ast

import (
	"interpreter/token"
	"bytes"
	"strings"
)

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}
func (hl *HashLiteral) expressionNode() {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left Expression
	Index Expression
}
func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

type ArrayLiteral struct {
	Token token.Token
	Elements []Expression
}
func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}
func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string { return sl.Token.Literal }

type CallExpression struct {
	Token token.Token
	Function Expression
	Arguments []Expression
}
func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type FunctionLiteral struct {
	Token token.Token
	Name *Identifier
	Parameters []*Identifier
	ReturnType *Identifier
	Body *BlockStatement
}
func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) ReturnTypeLiteral() string {
	if fl.ReturnType != nil {
		return fl.ReturnType.Literal()
	}
	return ""
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral() + " ")
	out.WriteString(fl.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	
	if fl.ReturnType != nil {
		out.WriteString(fl.ReturnTypeLiteral())
	}

	out.WriteString(fl.Body.String())

	return out.String()
}

type BlockStatement struct {
	Token token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type IfExpression struct {
	Token token.Token
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}
func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string { return b.Token.Literal }

type InfixExpression struct {
	Token token.Token
	Left Expression
	Operator string
	Right Expression
}
func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type PrefixExpression struct {
	Token token.Token
	Operator string
	Right Expression
}
func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// BEGIN VARIABLE TYPES
type TypeInt struct {
	Token token.Token	// the token.TYPE_INT token
	Name *Identifier
	Value Expression
}
func (ti *TypeInt) statementNode() {}
func (ti *TypeInt) TokenLiteral() string { return ti.Token.Literal }

type TypeBool struct {
	Token token.Token
	Name *Identifier
	Value Expression
}
func (tb *TypeBool) statementNode() {}
func (tb *TypeBool) TokenLiteral() string { return tb.Token.Literal }

type TypeString struct {
	Token token.Token
	Name *Identifier
	Value Expression
}
func (ts *TypeString) statementNode() {}
func (ts *TypeString) TokenLiteral() string { return ts.Token.Literal }

type TypeArray struct {
	Token token.Token
	Name *Identifier
	Value Expression
}
func (ta *TypeArray) statementNode() {}
func (ta *TypeArray) TokenLiteral() string { return ta.Token.Literal }
// END VARIABLE TYPES

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string { return il.Token.Literal }

// abstract syntax tree nodes
type Node interface {
	TokenLiteral() string
	String() string
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
func (i *Identifier) String() string { return i.Value }
func (i *Identifier) Literal() string {
	return i.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}	else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (tb *TypeBool) String() string {
	var out bytes.Buffer

	out.WriteString(tb.TokenLiteral() + " ")
	out.WriteString(tb.Name.String())
	out.WriteString(" = ")
	if tb.Value != nil {
		out.WriteString(tb.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (ti *TypeInt) String() string {
	var out bytes.Buffer

	out.WriteString(ti.TokenLiteral() + " ")
	out.WriteString(ti.Name.String())
	out.WriteString(" = ")
	if ti.Value != nil {
		out.WriteString(ti.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (ts *TypeString) String() string {
	var out bytes.Buffer

	out.WriteString(ts.TokenLiteral() + " ")
	out.WriteString(ts.Name.String())
	out.WriteString(" = ")
	if ts.Value != nil {
		out.WriteString(ts.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (ta *TypeArray) String() string {
	var out bytes.Buffer

	out.WriteString(ta.TokenLiteral() + " ")
	out.WriteString(ta.Name.String())
	out.WriteString(" = ")
	if ta.Value != nil {
		out.WriteString(ta.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
