package main

type coords struct {
	row uint
	col uint
}

//Keywords for query
type keyword string

const (
	selectKeyword keyword = "select"
	fromKeyword   keyword = "from"
	asKeyword     keyword = "as"
	tableKeyword  keyword = "table"
	createKeyword keyword = "create"
	insertkeyword keyword = "insert"
	intokeyword   keyword = "into"
	valuesKeyword keyword = "values"
	intKeyword    keyword = "int"
	textKeyword   keyword = "text"
	boolKeyword   keyword = "bool"
)

//Symbols for query
type symbol string

const (
	semicolonSymbol  symbol = ";"
	starSymbol       symbol = "*"
	commaSymbol      symbol = ","
	leftParenSymbol  symbol = "("
	rightParenSymbol symbol = ")"
)

type tokenKind uint

const (
	keywordKind tokenKind = iota
	symbolKind
	identifierKind
	stringkind
	numericKind
)

type token struct {
	value string
	kind  tokenKind
	loc   coords
}

type cursor struct {
	pointer uint
	loc     coords
}

func (t *token) equals(other *token) bool {
	return t.value == other.value && t.kind == other.kind
}
