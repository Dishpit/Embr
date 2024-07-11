package parser

import (
	"testing"
	"interpreter/ast"
	"interpreter/lexer"
)

func TestVarTypeIntStatements(t *testing.T) {
	input := `
int x = 420;
int y = 69;
int foobar = 8675309;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVarTypeIntStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testVarTypeIntStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "int" {
		t.Errorf("s.TokenLiteral not 'int'. got=%q", s.TokenLiteral())
		return false
	}

	intStmt, ok := s.(*ast.VarTypeInt)
	if !ok {
		t.Errorf("s not *ast.VarTypeInt. got=%T", s)
		return false
	}

	if intStmt.Name.Value != name {
		t.Errorf("intStmt.Name.Value not '%s'. got=%s", name, intStmt.Name.Value)
		return false
	}

	if intStmt.Name.TokenLiteral() != name {
		t.Errorf("intStmt.Name.TokenLiteral() not '%s'. got=%s", name, intStmt.Name.TokenLiteral())
		return false
	}
	return true
}
