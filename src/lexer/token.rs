use std::fmt::{Display, Formatter, Result};

#[derive(Debug)]
pub enum TokenKind {
    // utility tokens
    Illegal(String),
    Comment,

    // identifiers & literals
    Ident(String),
    Integer(String),
    Float(String),
    SQString(String),
    DQString(String),
    HDString(String),
    NDString(String),

    // keywords
    Abstract,
    And,
    Array,
    As,
    Break,
    Callable,
    Case,
    Catch,
    Class,
    Clone,
    Const,
    Continue,
    Declare,
    Default,
    Die,
    Do,
    Echo,
    Else,
    ElseIf,
    Empty,
    EndDeclare,
    EndFor,
    EndForEach,
    EndIf,
    EndSwitch,
    EndWhile,
    Eval,
    Exit,
    Extends,
    Final,
    Finally,
    For,
    ForEach,
    Funtion,
    Global,
    Goto,
    If,
    Implements,
    Include,
    IncludeOnce,
    InstanceOf,
    InsteadOf,
    Interface,
    IsSet,
    List,
    Namespace,
    New,
    Or,
    Print,
    Private,
    Protected,
    Public,
    Readonly,
    Require,
    RequireOnce,
    Return,
    Static,
    Switch,
    Throw,
    Trait,
    Try,
    Unset,
    Use,
    Var,
    While,
    Xor,
    Yield,
    YieldFrom,

    // one-character tokens
    LSquare,
    RSquare,
    LParen,
    RParen,
    LBrace,
    RBrace,
    Dot,
    Plus,
    Dash,
    Star,
    Tilde,
    Bang,
    Dollar,
    FSlash,
    BSlash,
    Percent,
    Less,
    More,
    Eq,
    Caret,
    Pipe,
    Amper,
    Question,
    Colon,
    Semi,
    Comma,

    // two-character tokens
    Arrow,
    TwoPlus,
    TwoDash,
    TwoStar,
    TwoLess,
    TwoMore,
    LessEq,
    MoreEq,
    TwoEq,
    BangEq,
    TwoPipe,
    TwoAmper,
    StarEq,
    FSlashEq,
    PercentEq,
    PlusEq,
    DashEq,
    DotEq,
    AmperEq,
    CaretEq,
    PipeEq,
    TwoQuestion,
    QuestionMore,
    PoundLSquare,

    // three-character tokens
    ThreeEq,
    BangTwoEq,
    TwoStarEq,
    EchoOpen,
    TwoLessEq,
    ThreeLess,
    TwoMoreEq,
    LessEqMore,
    TwoQuestionEq,
    ThreeDot,

    // many-character tokens
    Open,
}

pub struct Token {
    kind: TokenKind,
    start: (usize, usize),
    end: (usize, usize),
}

impl Token {
    pub fn emit() -> Token {
        Token {
            kind: TokenKind::Illegal(String::from("test")),
            start: (0, 0),
            end: (0, 0),
        }
    }
}

impl Display for Token {
    fn fmt(&self, f: &mut Formatter) -> Result {
        write!(
            f,
            "{:?}, [{}:{}] - [{}:{}]",
            self.kind, self.start.0, self.start.1, self.end.0, self.end.1
        )
    }
}
