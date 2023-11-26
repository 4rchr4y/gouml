package gouml

type Token int

const (
	// Special Tokens
	INVALID   Token = iota
	START_UML       // startuml
	END_UML         // enduml

	// Operators and delimiters
	operatorBeg
	AT        // @
	ADD       // +
	SUB       // -
	MUL       // *
	QUO       // /
	REM       // %
	LSS       // <
	GTR       // >
	ASSIGN    // =
	COMMA     // ,
	PERIOD    // .
	SEMICOLON // ;
	COLON     // :
	AND       // &
	OR        // |
	XOR       // ^

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
	classDiagramEntityBeg
	ABSTRACT   // abstract
	ANNOTATION // annotation
	CLASS      // class
	CIRCLE     // circle
	DIAMOND    // diamond
	ENTITY     // entity
	ENUM       // enum
	EXCEPTION  // exception
	INTERFACE  // interface
	METACLASS  // metaclass
	PROTOCOL   // protocol
	STRUCT     // struct
	classDiagramEntityEnd

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
	STEREOTYPE // stereotype
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
	INVALID: "<invalid>",
}

func (tok Token) String() string {
	return ""
}
