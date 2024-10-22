package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// utility tokens
	Illegal TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// identifiers & literals
	Ident    TokenType = "IDENTIFIER"
	Integer  TokenType = "INTEGER"
	Float    TokenType = "FLOAT"
	SQString TokenType = "STRING(SQ)"
	DQString TokenType = "STRING(DQ)"

	// keywords
	Abstract    TokenType = "abstract"
	And         TokenType = "and"
	Array       TokenType = "array"
	As          TokenType = "as"
	Break       TokenType = "break"
	Callable    TokenType = "callable"
	Case        TokenType = "case"
	Catch       TokenType = "catch"
	Class       TokenType = "class"
	Clone       TokenType = "clone"
	Const       TokenType = "const"
	Continue    TokenType = "continue"
	Declare     TokenType = "declare"
	Default     TokenType = "default"
	Die         TokenType = "die"
	Do          TokenType = "do"
	Echo        TokenType = "echo"
	Else        TokenType = "else"
	ElseIf      TokenType = "elseif"
	Empty       TokenType = "empty"
	EndDeclare  TokenType = "enddeclare"
	EndFor      TokenType = "endfor"
	EndForEach  TokenType = "endforeach"
	EndIf       TokenType = "endif"
	EndSwitch   TokenType = "endswitch"
	EndWhile    TokenType = "endwhile"
	Eval        TokenType = "eval"
	Exit        TokenType = "exit"
	Extends     TokenType = "extends"
	Final       TokenType = "final"
	Finally     TokenType = "finally"
	For         TokenType = "for"
	ForEach     TokenType = "foreach"
	Funtion     TokenType = "function"
	Global      TokenType = "global"
	Goto        TokenType = "goto"
	If          TokenType = "if"
	Implements  TokenType = "implements"
	Include     TokenType = "include"
	IncludeOnce TokenType = "include_once"
	InstanceOf  TokenType = "instanceof"
	InsteadOf   TokenType = "insteadof"
	Interface   TokenType = "interface"
	IsSet       TokenType = "isset"
	List        TokenType = "list"
	Namespace   TokenType = "namespace"
	New         TokenType = "new"
	Or          TokenType = "or"
	Print       TokenType = "print"
	Private     TokenType = "private"
	Protected   TokenType = "protected"
	Public      TokenType = "public"
	Require     TokenType = "require"
	RequireOnce TokenType = "require_once"
	Return      TokenType = "return"
	Static      TokenType = "static"
	Switch      TokenType = "switch"
	Throw       TokenType = "throw"
	Trait       TokenType = "trait"
	Try         TokenType = "try"
	Unset       TokenType = "unset"
	Use         TokenType = "use"
	Var         TokenType = "var"
	While       TokenType = "while"
	Xor         TokenType = "xor"
	Yield       TokenType = "yield"
	YieldFrom   TokenType = "yield from"

	// one-character punctuators
	LSquare TokenType = "["
	RSquare TokenType = "]"
	LParen  TokenType = "("
	RParen  TokenType = ")"
	LBrace  TokenType = "{"
	RBrace  TokenType = "}"

	// one-character operators
	Dot      TokenType = "."
	Plus     TokenType = "+"
	Dash     TokenType = "-"
	Star     TokenType = "*"
	Tilde    TokenType = "~"
	Bang     TokenType = "!"
	Dollar   TokenType = "$"
	FSlash   TokenType = "/"
	BSlash   TokenType = "\\"
	Percent  TokenType = "%"
	Less     TokenType = "<"
	More     TokenType = ">"
	Eq       TokenType = "="
	Caret    TokenType = "^"
	Pipe     TokenType = "|"
	Amper    TokenType = "&"
	Question TokenType = "?"
	Colon    TokenType = ":"
	Semi     TokenType = ";"
	Comma    TokenType = ","

	// two-character operators
	Arrow        TokenType = "->"
	TwoPlus      TokenType = "++"
	TwoDash      TokenType = "--"
	TwoStar      TokenType = "**"
	TwoLess      TokenType = "<<"
	TwoMore      TokenType = ">>"
	LessEq       TokenType = "<="
	MoreEq       TokenType = ">="
	TwoEq        TokenType = "=="
	BangEq       TokenType = "!="
	TwoPipe      TokenType = "||"
	TwoAmper     TokenType = "&&"
	StarEq       TokenType = "*="
	FSlashEq     TokenType = "/="
	PercentEq    TokenType = "%="
	PlusEq       TokenType = "+="
	DashEq       TokenType = "-="
	DotEq        TokenType = ".="
	AmperEq      TokenType = "&="
	CaretEq      TokenType = "^="
	PipeEq       TokenType = "|="
	TwoQuestion  TokenType = "??"
	QuestionMore TokenType = "?>"

	// three-character operators
	ThreeEq    TokenType = "==="
	BangTwoEq  TokenType = "!=="
	TwoStarEq  TokenType = "**="
	EchoOpen   TokenType = "<?="
	TwoLessEq  TokenType = "<<="
	ThreeLess  TokenType = "<<<"
	TwoMoreEq  TokenType = ">>="
	LessEqMore TokenType = "<=>"
	Spread     TokenType = "..."
	// NOTE: I think `??=` is missing from this

	Open TokenType = "<?php"
)
