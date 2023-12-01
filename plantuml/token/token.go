package token

import (
	"strconv"
)

type (
	// Token token is an identifier for various syntactic features supported by PlantUML.
	Token  int
	Lexeme string
)

const (
	// Special Tokens
	INVALID   Token = iota // invalid token with index 0 should always be the fist one
	START_UML              // startuml
	END_UML                // enduml
	EMPTY                  //

	// Identifiers and basic type literals
	literalBeg
	IDENT // Identifier token
	DUMMY // Empty identifier or token that will be thrown away later
	literalEnd

	// Operators and delimiters
	operatorBeg
	PLUS   // +
	MINUS  // -
	STAR   // *
	EQ     // =
	AT     // @
	QUO    // /
	REM    // %
	LSS    // <
	GTR    // >
	COMMA  // ,
	DOT    // .
	SEMI   // ;
	COLON  // :
	AND    // &
	OR     // |
	XOR    // ^
	TILDE  // ~
	POUND  // #
	DOLLAR // $

	LPAREN // (
	LBRACE // {
	LBRACK // [

	RPAREN // )
	RBRACE // }
	RBRACK // ]
	operatorEnd

	NL // \n

	// Keyword Tokens
	keywordBeg
	// Diagram Structure Elements
	diagramStructElementBeg
	PACKAGE // package
	MODULE  // module
	FOLDER  // folder
	diagramStructElementEnd

	// Class Diagram Entities
	classBeg
	ABSTRACT      // abstract
	ANNOTATION    // annotation
	CLASS         // class
	CIRCLE        // circle
	CIRCLE_SHORT  // ()
	DIAMOND       // diamond
	DIAMOND_SHORT // <>
	ENTITY        // entity
	ENUM          // enum
	EXCEPTION     // exception
	INTERFACE     // interface
	METACLASS     // metaclass
	PROTOCOL      // protocol
	STEREOTYPE    // stereotype
	STRUCT        // struct
	classEnd

	// UML Diagram Elements
	diagramElementBeg
	ACTOR       // actor
	USECASE     // usecase
	COMPONENT   // component
	NODE        // node
	DATABASE    // database
	ARTIFACT    // artifact
	PARTICIPANT // participant
	OBJECT      // object
	FRAME       // frame
	CLOUD       // cloud
	STORAGE     // storage
	AGENT       // agent
	STACK       // stack
	BOUNDARY    // boundary
	CONTROL     // control
	CARD        // card
	FILE        // file
	QUEUE       // queue
	ARCHIMATE   // archimate
	DETACH      // detach
	diagramElementEnd

	// Diagram Directives and Control Tokens
	directiveAndControlBeg
	ALSO       // also
	AUTONUMBER // autonumber
	CAPTION    // caption
	TITLE      // title
	NEWPAGE    // newpage
	LOOP       // loop
	BREAK      // break
	CRITICAL   // critical
	NOTE       // note
	LEGEND     // legend
	GROUP      // group
	LEFT       // left
	RIGHT      // right
	LINK       // link
	OVER       // over
	ACTIVATE   // activate
	DEACTIVATE // deactivate
	DESTROY    // destroy
	CREATE     // create
	FOOTBOX    // footbox
	HIDE       // hide
	SHOW       // show
	SKINPARAM  // skinparam
	SKIN       // skin
	BOTTOM     // bottom
	NAMESPACE  // namespace
	PAGE       // page
	DOWN       // down
	ELSE       // else
	ENDIF      // endif
	PARTITION  // partition
	FOOTER     // footer
	HEADER     // header
	CENTER     // center
	ROTATE     // rotate
	RETURN     // return
	REPEAT     // repeat
	START      // start
	STOP       // stop
	WHILE      // while
	ENDWHILE   // endwhile
	FORK       // fork
	AGAIN      // again
	KILL       // kill
	ORDER      // order
	MAINFRAME  // mainframe
	ACROSS     // across
	SPLIT      // split
	STYLE      // style
	SPRITE     // sprite
	directiveAndControlEnd

	// Preprocessor and Styling Directives
	EXIT         // exit
	INCLUDE      // include
	PRAGMA       // pragma
	UNDEF        // undef
	IFDEF        // ifdef
	IFNDEF       // ifndef
	FUNCTION     // function
	PROCEDURE    // procedure
	ENDFUNCTION  // endfunction
	ENDPROCEDURE // endprocedure
	UNQUOTED     // unquoted
	STARTSUB     // startsub
	ENDSUB       // endsub
	ASSERT       // assert
	LOCAL        // local

	// Styling Elements
	stylingBeg
	DEFINE           // define
	ALIAS            // alias
	SHAPE            // shape
	LABEL            // label
	BACKGROUND_COLOR // BackgroundColor
	COLOR_CAPITAL    // Color
	COLOR_LOWER      // color
	ENTITY_CAPITAL   // Entity
	ENTITY_UPPER     // ENTITY
	COLOR_UPPER      // COLOR
	LARGE            // LARGE
	stylingEnd
	keywordEnd

	defineBeg
	DEFINE_LONG     // !definelong
	END_DEFINE_LONG // !enddefinelong
	DEFINE_SHORT    // !define
	defineEnd

	groupBeg
	GROUP_SOLID  // --
	GROUP_DOTTY  // ..
	GROUP_WEAK   // __
	GROUP_DOUBLE // ==
	groupEnd

	// Relations between classes
	relationBeg
	RELATION_ASSOCIATION      // ..
	RELATION_ASSOCIATION_WEAK // --

	LRELATION_EXTENSION      // <|--
	LRELATION_COMPOSITION    // *--
	LRELATION_AGGREGATION    // o--
	LRELATION_EXTENSION_WEAK // <|..
	LRELATION_ARROW          // <--
	LRELATION_ARROW_WEAK     // <..
	LRELATION_ARROW_STRONG   // <--*
	LRELATION_SQUARE         // #--
	LRELATION_CANCELLATION   // x--
	LRELATION_DEPARTURE      // }--
	LRELATION_ADDITION       // +--
	LRELATION_EXTENSION_V2   // ^--

	RRELATION_EXTENSION      // --|>
	RRELATION_COMPOSITION    // --*
	RRELATION_AGGREGATION    // --o
	RRELATION_EXTENSION_WEAK // ..|>
	RRELATION_ARROW          // -->
	RRELATION_ARROW_WEAK     // ..>
	RRELATION_ARROW_STRONG   // *-->
	RRELATION_SQUARE         // --#
	RRELATION_CANCELLATION   // --x
	RRELATION_DEPARTURE      // --{
	RRELATION_ADDITION       // --+
	RRELATION_EXTENSION_V2   // --^
	relationEnd
)

var tokens = [...]string{
	INVALID:   "invalid",
	START_UML: "startuml",
	END_UML:   "enduml",
	EMPTY:     "",

	PLUS:   "+-",
	MINUS:  "-",
	STAR:   "*",
	EQ:     "=",
	AT:     "@",
	QUO:    "/",
	REM:    "%",
	LSS:    "<",
	GTR:    ">",
	COMMA:  ",",
	DOT:    ".",
	SEMI:   ";",
	COLON:  ":",
	AND:    "&",
	OR:     "|",
	XOR:    "^",
	TILDE:  "~",
	POUND:  "#",
	DOLLAR: "$",

	LPAREN: "(",
	LBRACE: "{",
	LBRACK: "[",

	RPAREN: ")",
	RBRACE: "}",
	RBRACK: "]",

	CLASS: "class",

	// TODO: implement all list of tokens const where key is a const
	// name and value is a constant comment
}

// String returns the string corresponding to the token id
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokens)) {
		return tokens[tok]
	}

	// corresponding as invalid token
	return tokens[0] + "<" + strconv.Itoa(int(tok)) + ">"
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keywordEnd-(keywordBeg+1))
	// for i := keywordBeg + 1; i < keywordEnd; i++ {
	// 	keywords[tokens[i]] = i
	// }
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
func Lookup(ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}

	return IDENT
}

// IsKeyword reports whether name is a PlantUML keyword.
func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

// IsClassWord reports whether name that is acceptable in PlantUML class diagram.
func IsClassWord(name string) bool {
	tok, ok := keywords[name]
	return ok && classBeg < tok && tok < classEnd
}

func (tok Token) IsLiteral() bool {
	return literalBeg < tok && tok < literalEnd
}

// TokenSet is a result set of tokenization of any entity within a defined PlantUML AST.
type TokenSet struct {
	tokSeq []Token  // sequence of tokens, e.g. `IDENTIFIER``
	lexSeq []Lexeme // sequence of tokens lexical values, e.g. `MyFunc`
	index  int
}

func NewTokenSet() *TokenSet {
	return &TokenSet{lexSeq: make([]Lexeme, 0), tokSeq: make([]Token, 0)}
}

func (set *TokenSet) AddToken(tok Token) {
	set.tokSeq = append(set.tokSeq, tok)
}

func (set *TokenSet) AddLexeme(lex string) {
	set.tokSeq = append(set.tokSeq, IDENT)
	set.lexSeq = append(set.lexSeq, Lexeme(lex))
}

func (ts *TokenSet) Next() (string, bool) {
	// Проверяем, достигнут ли конец массива tokSeq.
	if ts.index >= len(ts.tokSeq) {
		return "", false
	}

	// Получаем текущий токен.
	currentToken := ts.tokSeq[ts.index]

	// Инкрементируем индекс после получения текущего токена.
	ts.index++

	// Проверяем, является ли текущий токен IDENT.
	if currentToken == IDENT {
		// Если это IDENT и индекс в пределах lexSeq, возвращаем Lexeme.
		if ts.index-1 < len(ts.lexSeq) {
			return string(ts.lexSeq[ts.index-1]), true
		}
	}

	// В противном случае возвращаем строковое представление токена.
	return currentToken.String(), true
}
