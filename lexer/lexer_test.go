package lexer

import (
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		src            string
		expectedTokens []string
	}{
		{
			"x, y => x if (x > 100) && y < 2",
			[]string{
				"ident(x)",
				"comma",
				"ident(y)",
				"reactionOp",
				"ident(x)",
				"if",
				"openBracket",
				"ident(x)",
				"greaterThan",
				"number(100)",
				"closeBracket",
				"and",
				"ident(y)",
				"lessThan",
				"number(2)",
			},
		},
		{
			"x,y=>x",
			[]string{
				"ident(x)",
				"comma",
				"ident(y)",
				"reactionOp",
				"ident(x)",
			},
		},
		{
			"x,y=>   x  ",
			[]string{
				"ident(x)",
				"comma",
				"ident(y)",
				"reactionOp",
				"ident(x)",
			},
		},
		{
			"{1,2,3}",
			[]string{
				"openCurlyBracket",
				"number(1)",
				"comma",
				"number(2)",
				"comma",
				"number(3)",
				"closeCurlyBracket",
			},
		},
		{
			"{1,23}",
			[]string{
				"openCurlyBracket",
				"number(1)",
				"comma",
				"number(23)",
				"closeCurlyBracket",
			},
		},
		{
			"{1,2 3}",
			[]string{
				"openCurlyBracket",
				"number(1)",
				"comma",
				"number(2)",
				"number(3)",
				"closeCurlyBracket",
			},
		},
	}

	for testNo, test := range tests {
		actual, err := FromString(test.src).RemainingTokens()
		expected := test.expectedTokens

		if err != nil {
			t.Errorf("error lexing test %d. error=%v", testNo, err)
			return
		}

		if len(actual) != len(expected) {
			t.Errorf("test %d, incorrect number of tokens. expected=%d, got=%d", testNo, len(expected), len(actual))
			return
		}

		for i := range expected {
			et := expected[i]
			at := actual[i].String()

			if et != at {
				t.Errorf("test %d, incorrect token at index %d. expected=%q, got=%q", testNo, i, et, at)
				return
			}
		}
	}
}

func TestLexError(t *testing.T) {
	src := "x, y => x if (x > 100) & y < 2"
	_, err := FromString(src).RemainingTokens()

	if err == nil {
		t.Error("expected an error")
		return
	}

	expectedError := "unknown token '&' at repl:1:25"

	if err.Error() != expectedError {
		t.Errorf("incorrect error. expected=%q, got=%q", expectedError, err.Error())
	}
}
