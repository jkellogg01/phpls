package lexer

import (
	"testing"

	"github.com/jkellogg01/phpls/token"
)

// redundant now but will definitely be needed later
type TestToken struct {
	Type    token.TokenType
	Literal string
}

func TestNextToken(t *testing.T) {
	testCases := []struct {
		Name   string
		Input  string
		Expect []TestToken
	}{
		{
			"double-character punctuators",
			`[](){}`,
			[]TestToken{
				{token.LSquare, "["},
				{token.RSquare, "]"},
				{token.LParen, "("},
				{token.RParen, ")"},
				{token.LBrace, "{"},
				{token.RBrace, "}"},
			},
		},
		{
			"single-character operators",
			`$\:;,`,
			[]TestToken{
				{token.Dollar, "$"},
				{token.BSlash, "\\"},
				{token.Colon, ":"},
				{token.Semi, ";"},
				{token.Comma, ","},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.Input)

			for _, et := range tc.Expect {
				tok := l.NextToken()

				if tok.Type != et.Type {
					t.Fatalf("wrong TokenType:\texpect %q\tgot %q",
						et.Type, tok.Type)
				}

				if tok.Literal != et.Literal {
					t.Fatalf("wrong Literal:\texpect %q\tgot %q",
						et.Literal, tok.Literal)
				}
			}
		})
	}
}
