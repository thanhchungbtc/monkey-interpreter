package parser

import (
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"testing"
)

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T",
			stmt.Expression)
	}

	if !testIdentifier(t, exp.Function, "add") {
		return
	}

	if len(exp.Arguments) != 3 {
		t.Fatalf("wrong length of arguments. got=%d", len(exp.Arguments))
	}

	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func TestCallExpressionParametersParsing(t *testing.T) {
	var tests = []struct {
		input         string
		expectedIdent string
		expectedArgs  []string
	}{
		{"add();", "add", []string{}},
		{"add(1);", "add", []string{"1"}},
		{"add(1, 2 + 3);", "add", []string{"1", "(2 + 3)"}},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements, got = %d", 1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.CallExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.CallExpression, got = %T", stmt.Expression)
		}
		if !testIdentifier(t, exp.Function, "add") {
			return
		}
		if len(exp.Arguments) != len(tt.expectedArgs) {
			t.Errorf("wrong length of arguments, want = %d, got = %d", len(tt.expectedArgs), len(exp.Arguments))
		}
		for i, arg := range tt.expectedArgs {
			if exp.Arguments[i].String() != arg {
				t.Errorf("wrong argument, want = %s, got = %s", arg, exp.Arguments[i].String())
			}
		}
	}

}
