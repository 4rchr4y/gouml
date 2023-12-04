package token

import (
	"strconv"
)

type (
	// Token token is an identifier for various syntactic features supported by PlantUML.
	Token int

	// TokenSet is a result set of tokenization of any entity within a defined PlantUML AST.
	TokenSet struct {
		TokSeq []Token  // sequence of tokens, e.g. `IDENTIFIER``
		LexSeq []string // sequence of tokens lexical values, e.g. `MyFunc`
	}
)

const (
	// Special Tokens
	INVALID   Token = iota
	START_UML       // startuml
	END_UML         // enduml

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

	// UML Diagram Elements 🤓☝ вапщето это deployment diagram elements
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
	ARCHIMATE   // archimate // а это Archimate Diagram
	DETACH      // detach // а это, по идее, вообще Activity Diagram
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

	IDENT: "IDENT", // хз, что тут должно быть
	DUMMY: "DUMMY", // хз, что тут должно быть

	AT:    "@",
	PLUS:  "+",
	MINUS: "-",
	STAR:  "*", // это не умножение???
	QUO:   "/",
	REM:   "%",
	LSS:   "<",
	GTR:   ">",
	EQ:    "=", // equal, а не присваивание??
	COMMA: ",",
	DOT:   ".",
	SEMI:  ";",
	COLON: ":",
	AND:   "&",
	OR:    "|",
	XOR:   "^",

	LPAREN: "(",
	LBRACE: "{",
	LBRACK: "[",

	RPAREN: ")",
	RBRACE: "}",
	RBRACK: "]",

	PACKAGE: "package",
	MODULE:  "module",
	FOLDER:  "folder",

	ABSTRACT:      "abstract",
	ANNOTATION:    "annotation",
	CLASS:         "class",
	CIRCLE:        "circle",
	CIRCLE_SHORT:  "()",
	DIAMOND:       "diamond",
	DIAMOND_SHORT: "<>",
	ENTITY:        "entity",
	ENUM:          "enum",
	EXCEPTION:     "exception",
	INTERFACE:     "interface",
	METACLASS:     "metaclass",
	PROTOCOL:      "protocol",
	STEREOTYPE:    "stereotype",
	STRUCT:        "struct",

	ACTOR:       "actor",
	USECASE:     "usecase",
	COMPONENT:   "component",
	NODE:        "node",
	DATABASE:    "database",
	ARTIFACT:    "artifact",
	PARTICIPANT: "participant",
	OBJECT:      "object",
	FRAME:       "frame",
	CLOUD:       "cloud",
	STORAGE:     "storage",
	AGENT:       "agent",
	STACK:       "stack",
	BOUNDARY:    "boundary",
	CONTROL:     "control",
	CARD:        "card",
	FILE:        "file",
	QUEUE:       "queue",
	ARCHIMATE:   "archimate",
	DETACH:      "detach",

	ALSO:       "also",
	AUTONUMBER: "autonumber",
	CAPTION:    "caption",
	TITLE:      "title",
	NEWPAGE:    "newpage",
	LOOP:       "loop",
	BREAK:      "break",
	CRITICAL:   "critical",
	NOTE:       "note",
	LEGEND:     "legend",
	GROUP:      "group",
	LEFT:       "left",
	RIGHT:      "right",
	LINK:       "link",
	OVER:       "over",
	ACTIVATE:   "activate",
	DEACTIVATE: "deactivate",
	DESTROY:    "destroy",
	CREATE:     "create",
	FOOTBOX:    "footbox",
	HIDE:       "hide",
	SHOW:       "show",
	SKINPARAM:  "skinparam",
	SKIN:       "skin",
	BOTTOM:     "bottom",
	NAMESPACE:  "namespace",
	PAGE:       "page",
	DOWN:       "down",
	ELSE:       "else",
	ENDIF:      "endif",
	PARTITION:  "partition",
	FOOTER:     "footer",
	HEADER:     "header",
	CENTER:     "center",
	ROTATE:     "rotate",
	RETURN:     "return",
	REPEAT:     "repeat",
	START:      "start",
	STOP:       "stop",
	WHILE:      "while",
	ENDWHILE:   "endwhile",
	FORK:       "fork",
	AGAIN:      "again",
	KILL:       "kill",
	ORDER:      "order",
	MAINFRAME:  "mainframe",
	ACROSS:     "across",
	SPLIT:      "split",
	STYLE:      "style",
	SPRITE:     "sprite",

	EXIT:         "exit",
	INCLUDE:      "include",
	PRAGMA:       "pragma",
	UNDEF:        "undef",
	IFDEF:        "ifdef",
	IFNDEF:       "ifndef",
	FUNCTION:     "function",
	PROCEDURE:    "procedure",
	ENDFUNCTION:  "endfunction",
	ENDPROCEDURE: "endprocedure",
	UNQUOTED:     "unquoted",
	STARTSUB:     "startsub",
	ENDSUB:       "endsub",
	ASSERT:       "assert",
	LOCAL:        "local",

	DEFINE:           "define",
	ALIAS:            "alias",
	SHAPE:            "shape",
	LABEL:            "label",
	BACKGROUND_COLOR: "BackgroundColor",
	COLOR_CAPITAL:    "Color",
	COLOR_LOWER:      "color",
	ENTITY_CAPITAL:   "Entity",
	ENTITY_UPPER:     "ENTITY",
	COLOR_UPPER:      "COLOR",
	LARGE:            "LARGE",

	DEFINE_LONG:     "!definelong",
	END_DEFINE_LONG: "!enddefinelong",
	DEFINE_SHORT:    "!define",

	GROUP_SOLID:  "--",
	GROUP_DOTTY:  "..",
	GROUP_WEAK:   "__",
	GROUP_DOUBLE: "==",

	RELATION_ASSOCIATION:      "..",
	RELATION_ASSOCIATION_WEAK: "--",
	LRELATION_EXTENSION:       "<|--",
	LRELATION_COMPOSITION:     "*--",
	LRELATION_AGGREGATION:     "o--",
	LRELATION_EXTENSION_WEAK:  "<|..",
	LRELATION_ARROW:           "<--",
	LRELATION_ARROW_WEAK:      "<..",
	LRELATION_ARROW_STRONG:    "<--*",
	LRELATION_SQUARE:          "#--",
	LRELATION_CANCELLATION:    "x--",
	LRELATION_DEPARTURE:       "}--",
	LRELATION_ADDITION:        "+--",
	LRELATION_EXTENSION_V2:    "^--",

	RRELATION_EXTENSION:      "--|>",
	RRELATION_COMPOSITION:    "--*",
	RRELATION_AGGREGATION:    "--o",
	RRELATION_EXTENSION_WEAK: "..|>",
	RRELATION_ARROW:          "-->",
	RRELATION_ARROW_WEAK:     "..>",
	RRELATION_ARROW_STRONG:   "*-->",
	RRELATION_SQUARE:         "--#",
	RRELATION_CANCELLATION:   "--x",
	RRELATION_DEPARTURE:      "--{",
	RRELATION_ADDITION:       "--+",
	RRELATION_EXTENSION_V2:   "--^",
}

// String returns the string corresponding to the token id
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokens)) {
		return tokens[tok]
	}

	return tokens[0] + "<" + strconv.Itoa(int(tok)) + ">"
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keywordEnd-(keywordBeg+1))
	for i := keywordBeg + 1; i < keywordEnd; i++ {
		keywords[tokens[i]] = i
	}
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
