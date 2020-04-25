package evaluator

import (
	"monkey-interpreter/object"
	"testing"
)

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"
	evaluted := testEval(input)
	fn, ok := evaluted.(*object.Function)
	if !ok {
		t.Fatal("object is not function, got = %T, (%+v)", evaluted, evaluted)
	}
	if len(fn.Parameters) != 1 {
		t.Fatal("wrong parameter, want = %d, got = %d", 1, len(fn.Parameters))
	}
	if fn.Parameters[0].String() != "x" {
		t.Fatal("parameter is not %s, got = %s", "x", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatal("body is not %q, got = %q", expectedBody, fn.Body.String())
	}
}
