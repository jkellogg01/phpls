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
			`[](){}$\:;,`,
			[]TestToken{
				{token.LSquare, "["},
				{token.RSquare, "]"},
				{token.LParen, "("},
				{token.RParen, ")"},
				{token.LBrace, "{"},
				{token.RBrace, "}"},
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
			`< <= << <<= <=> <?= <?php`,
			[]TestToken{
				{token.Less, "<"},
				{token.LessEq, "<="},
				{token.TwoLess, "<<"},
				{token.TwoLessEq, "<<="},
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
		{
			"single-line comments",
			`// this is a comment
# this is also a comment
# comments can end by exiting php-world as well ?>`,
			[]TestToken{
				{token.Comment, "// this is a comment\n"},
				{token.Comment, "# this is also a comment\n"},
				{token.Comment, "# comments can end by exiting php-world as well "},
				{token.QuestionMore, "?>"},
			},
		},
		{
			"multi-line comments",
			`/*
a multi-line comment starts with 'fslash-star'
and only ends with 'star-fslash'.

it should be treated like whitespace.
*/`,
			[]TestToken{
				{token.Comment, "/*\na multi-line comment starts with 'fslash-star'\nand only ends with 'star-fslash'.\n\nit should be treated like whitespace.\n*/"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.Input)

			idx := 0
			for {
				tok := l.NextToken()
				t.Logf("got token: %+v", tok)
				if tok.Type == token.EOF {
					break
				} else if idx >= len(tc.Expect) {
					t.Fatalf("test input produced at least %d non-EOF tokens, expect %d exactly", idx+1, len(tc.Expect))
				}
				et := tc.Expect[idx]
				idx += 1

				if tok.Type != et.Type {
					t.Errorf("wrong TokenType:\texpect %q\tgot %q",
						et.Type, tok.Type)
				}

				if tok.Literal != et.Literal {
					t.Errorf("wrong Literal:\texpect %q\tgot %q",
						et.Literal, tok.Literal)
				}
			}
		})
	}
}
