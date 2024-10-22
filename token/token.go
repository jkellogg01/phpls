package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// utility tokens
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	// identifiers & literals
	Ident    = "IDENTIFIER"
	Integer  = "INTEGER"
	Float    = "FLOAT"
	SQString = "STRING(SQ)"
	DQString = "STRING(DQ)"

	// keywords
	Abstract    = "abstract"
	And         = "and"
	Array       = "array"
	As          = "as"
	Break       = "break"
	Callable    = "callable"
	Case        = "case"
	Catch       = "catch"
	Class       = "class"
	Clone       = "clone"
	Const       = "const"
	Continue    = "continue"
	Declare     = "declare"
	Default     = "default"
	Die         = "die"
	Do          = "do"
	Echo        = "echo"
	Else        = "else"
	ElseIf      = "elseif"
	Empty       = "empty"
	EndDeclare  = "enddeclare"
	EndFor      = "endfor"
	EndForEach  = "endforeach"
	EndIf       = "endif"
	EndSwitch   = "endswitch"
	EndWhile    = "endwhile"
	Eval        = "eval"
	Exit        = "exit"
	Extends     = "extends"
	Final       = "final"
	Finally     = "finally"
	For         = "for"
	ForEach     = "foreach"
	Funtion     = "function"
	Global      = "global"
	Goto        = "goto"
	If          = "if"
	Implements  = "implements"
	Include     = "include"
	IncludeOnce = "include_once"
	InstanceOf  = "instanceof"
	InsteadOf   = "insteadof"
	Interface   = "interface"
	IsSet       = "isset"
	List        = "list"
	Namespace   = "namespace"
	New         = "new"
	Or          = "or"
	Print       = "print"
	Private     = "private"
	Protected   = "protected"
	Public      = "public"
	Require     = "require"
	RequireOnce = "require_once"
	Return      = "return"
	Static      = "static"
	Switch      = "switch"
	Throw       = "throw"
	Trait       = "trait"
	Try         = "try"
	Unset       = "unset"
	Use         = "use"
	Var         = "var"
	While       = "while"
	Xor         = "xor"
	Yield       = "yield"
	YieldFrom   = "yield from"

	// one-character punctuators
	LSquare = "["
	RSquare = "]"
	LParen  = "("
	RParen  = ")"
	LBrace  = "{"
	RBrace  = "}"

	// one-character operators
	Dot      = "."
	Plus     = "+"
	Dash     = "-"
	Star     = "*"
	Tilde    = "~"
	Bang     = "!"
	Dollar   = "$"
	FSlash   = "/"
	BSlash   = "\\"
	Percent  = "%"
	Less     = "<"
	More     = ">"
	Eq       = "="
	Caret    = "^"
	Pipe     = "|"
	Amper    = "&"
	Question = "?"
	Colon    = ":"
	Semi     = ";"
	Comma    = ","

	// two-character operators
	Arrow       = "->"
	TwoPlus     = "++"
	TwoDash     = "--"
	TwoStar     = "**"
	TwoLess     = "<<"
	TwoMore     = ">>"
	LessEq      = "<="
	MoreEq      = ">="
	TwoEq       = "=="
	BangEq      = "!="
	TwoPipe     = "||"
	TwoAmper    = "&&"
	StarEq      = "*="
	FSlashEq    = "/="
	PercentEq   = "%="
	PlusEq      = "+="
	DashEq      = "-="
	DotEq       = ".="
	AmperEq     = "&="
	CaretEq     = "^="
	PipeEq      = "|="
	TwoQuestion = "??"

	// three-character operators
	ThreeEq    = "==="
	BangTwoEq  = "!=="
	TwoStarEq  = "**="
	TwoLessEq  = "<<="
	ThreeLess  = "<<<"
	TwoMoreEq  = ">>="
	LessEqMore = "<=>"
	Spread     = "..."
	// NOTE: I think `??=` is missing from this
)
