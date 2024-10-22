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
			"one character tokens",
			`[](){}#$\:;,`,
			[]TestToken{
				{token.LSquare, "["},
				{token.RSquare, "]"},
				{token.LParen, "("},
				{token.RParen, ")"},
				{token.LBrace, "{"},
				{token.RBrace, "}"},
				{token.Pound, "#"},
				{token.Dollar, "$"},
				{token.BSlash, "\\"},
				{token.Colon, ":"},
				{token.Semi, ";"},
				{token.Comma, ","},
			},
		},
		{
			"two character tokens",
			// NOTE: started adding spaces here so I don't have to worry about
			// maximal munches when putting these in order
			`- -> -- -= + ++ += | || |= & && &= / /= /* % %= ^ ^=`,
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
				{token.FSlashStar, "/*"},
				{token.Percent, "%"},
				{token.PercentEq, "%="},
				{token.Caret, "^"},
				{token.CaretEq, "^="},
			},
		},
		{
			"three character operators",
			`= == === ! != !== * */ *= ** **= > >= >> >>= ? ?> ?? ??= . .= ...`,
			[]TestToken{
				{token.Eq, "="},
				{token.TwoEq, "=="},
				{token.ThreeEq, "==="},
				{token.Bang, "!"},
				{token.BangEq, "!="},
				{token.BangTwoEq, "!=="},
				{token.Star, "*"},
				{token.StarFSlash, "*/"},
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
		{
			"single-quoted string literals",
			`'use \\ to escape a backslash'
            'use \' to escape a single quote'
            '\any \other \character \can \be \escaped, \with \no \effect'
            b'binary strings should be treated as equivalent to regular strings'
            B'and can be prefixed with a lowercase or uppercase b'
            'single-quoted
strings can also
accomodate line breaks'`,
			[]TestToken{
				{token.SQString, "'use \\\\ to escape a backslash'"},
				{token.SQString, "'use \\' to escape a single quote'"},
				{token.SQString, "'\\any \\other \\character \\can \\be \\escaped, \\with \\no \\effect'"},
				{token.SQString, "b'binary strings should be treated as equivalent to regular strings'"},
				{token.SQString, "B'and can be prefixed with a lowercase or uppercase b'"},
				{token.SQString, "'single-quoted\nstrings can also\naccomodate line breaks'"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.Input)

			for _, et := range tc.Expect {
				tok := l.NextToken()

				if tok.Type != et.Type {
					t.Errorf("wrong TokenType:\texpect %q\tgot %q",
						et.Type, tok.Type)
				}

				if tok.Literal != et.Literal {
					t.Errorf("wrong Literal:\texpect %q\tgot %q",
						et.Literal, tok.Literal)
				}

				if !t.Failed() {
					t.Logf("parsed token \"%s\" correctly", tok.Literal)
				}
			}
		})
	}
}
