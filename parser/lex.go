package parser

import (
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher"
	vMatcher "github.com/viant/velty/parser/matcher"
)

const (
	specialSignToken = iota
	whiteSpaceToken

	selectorStartToken
	selectorBlockToken
	selectorToken

	ifToken
	expressionBlockToken
	expressionEndToken

	stringToken
	booleanToken
	numberMatcher

	greaterToken
	greaterEqualToken
	lessToken
	lessEqualToken

	equalToken
	notEqualToken
	negationToken

	andToken
	orToken

	addToken
	subToken
	mulToken
	quoToken

	binaryExpressionStartToken
)

var WhiteSpace = parsly.NewToken(whiteSpaceToken, "Whitespace", matcher.NewWhiteSpace())
var SpecialSign = parsly.NewToken(specialSignToken, "Special sign", matcher.NewRunes([]rune{'#', '$'}))

var SelectorBlock = parsly.NewToken(selectorBlockToken, "Selector block", matcher.NewBlock('{', '}', '\\'))
var Selector = parsly.NewToken(selectorToken, "Selector", vMatcher.NewIdentity())
var SelectorStart = parsly.NewToken(selectorStartToken, "Selector start", matcher.NewRunes([]rune{'$'}))

var If = parsly.NewToken(ifToken, "If", matcher.NewFragment("if"))
var ExpressionBlock = parsly.NewToken(expressionBlockToken, "Expression block", matcher.NewBlock('(', ')', '\\'))
var ExpressionEnd = parsly.NewToken(expressionEndToken, "Expression end", vMatcher.NewExpressionEnd())

var Equal = parsly.NewToken(equalToken, "Equal", matcher.NewFragment("=="))
var NotEqual = parsly.NewToken(notEqualToken, "Not equal", matcher.NewFragment("!="))
var Negation = parsly.NewToken(negationToken, "Negation", matcher.NewByte('!'))

var Greater = parsly.NewToken(greaterToken, "Greater", matcher.NewByte('>'))
var GreaterEqual = parsly.NewToken(greaterEqualToken, "Greater or equal", matcher.NewFragment(">="))
var Less = parsly.NewToken(lessToken, "Less", matcher.NewByte('<'))
var LessEqual = parsly.NewToken(lessEqualToken, "Less or equal", matcher.NewFragment("<="))

var And = parsly.NewToken(andToken, "And", matcher.NewFragment("&&"))
var Or = parsly.NewToken(orToken, "Or", matcher.NewFragment("||"))

var String = parsly.NewToken(stringToken, "String", vMatcher.NewStringMatcher('"'))
var Boolean = parsly.NewToken(booleanToken, "Boolean", matcher.NewFragments([]byte("true"), []byte("false")))
var Number = parsly.NewToken(numberMatcher, "Number", matcher.NewNumber())

var Add = parsly.NewToken(addToken, "Add", matcher.NewByte('+'))
var Sub = parsly.NewToken(subToken, "Subtract", matcher.NewByte('-'))
var Multiply = parsly.NewToken(mulToken, "Multiply", matcher.NewByte('*'))
var Quo = parsly.NewToken(quoToken, "Quo", matcher.NewByte('/'))
