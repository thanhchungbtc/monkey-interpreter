package evaluator

import (
	"monkey-interpreter/object"
	"testing"
)

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 > 2", false},
		{"2 > 1", true},
		{"1 < 2", true},
		{"2 < 1", false},
		{"1 == 2", false},
		{"1 == 1", true},
		{"1 != 2", true},
		{"1 != 1", false},
		{"true != true", false},
		{"true == true", true},
		{"true != true", false},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}

}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean, got=%T(%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value, want=%t, got=%t", expected, result.Value)
		return false
	}
	return true
}
