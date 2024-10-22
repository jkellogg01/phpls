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
			"one character punctuators",
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
			"one character operators",
			`$\:;,`,
			[]TestToken{
				{token.Dollar, "$"},
				{token.BSlash, "\\"},
				{token.Colon, ":"},
				{token.Semi, ";"},
				{token.Comma, ","},
			},
		},
		{
			"two character operators",
			`- -> -- -= + ++ += | || |= & && &= / /= % %= ^ ^=`,
			[]TestToken{
				{token.Dash, "-"},
				{token.Arrow, "->"},
				{token.TwoDash, "--"},
				{token.DashEq, "-="},
				{token.Plus, "+"},
				{token.TwoPlus, "++"},
				{token.PlusEq, "+="},
				{token.Pipe, "|"},
				{token.TwoPipe, "||"},
				{token.PipeEq, "|="},
				{token.Amper, "&"},
				{token.TwoAmper, "&&"},
				{token.AmperEq, "&="},
				{token.FSlash, "/"},
				{token.FSlashEq, "/="},
				{token.Percent, "%"},
				{token.PercentEq, "%="},
				{token.Caret, "^"},
				{token.CaretEq, "^="},
			},
		},
		{
			"three character operators",
			`= == === ! != !== * *= ** **= > >= >> >>= ? ?> ?? ??= . .= ...`,
			[]TestToken{
				{token.Eq, "="},
				{token.TwoEq, "=="},
				{token.ThreeEq, "==="},
				{token.Bang, "!"},
				{token.BangEq, "!="},
				{token.BangTwoEq, "!=="},
				{token.Star, "*"},
				{token.StarEq, "*="},
				{token.TwoStar, "**"},
				{token.TwoStarEq, "**="},
				{token.More, ">"},
				{token.MoreEq, ">="},
				{token.TwoMore, ">>"},
				{token.TwoMoreEq, ">>="},
				{token.Question, "?"},
				{token.QuestionMore, "?>"},
				{token.TwoQuestion, "??"},
				{token.TwoQuestionEq, "??="},
				{token.Dot, "."},
				{token.DotEq, ".="},
				{token.ThreeDot, "..."},
			},
		},
		{
			"many character operators",
			`< <= << <<= <<< <=> <?= <?php`,
			[]TestToken{
				{token.Less, "<"},
				{token.LessEq, "<="},
				{token.TwoLess, "<<"},
				{token.TwoLessEq, "<<="},
				{token.ThreeLess, "<<<"},
				{token.LessEqMore, "<=>"},
				{token.EchoOpen, "<?="},
				{token.Open, "<?php"},
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

				t.Logf("parsed token \"%s\" correctly", tok.Literal)
			}
		})
	}
}
