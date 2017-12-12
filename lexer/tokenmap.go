package lexer

var OPERATORS_AND_PUNCTUATORS = map[string]TokenKind{
	"[":  OpenBracketToken,
	"]":  CloseBracketToken,
	"(":  OpenParenToken,
	")":  CloseParenToken,
	"{":  OpenBraceToken,
	"}":  CloseBraceToken,
	".":  DotToken,
	"->": ArrowToken,
	//	":":         DoubleArrowToken,
	"++":        PlusPlusToken,
	"--":        MinusMinusToken,
	"**":        AsteriskAsteriskToken,
	"*":         AsteriskToken,
	"+":         PlusToken,
	"-":         MinusToken,
	"~":         TildeToken,
	"!":         ExclamationToken,
	"$":         DollarToken,
	"/":         SlashToken,
	"%":         PercentToken,
	"<<":        LessThanLessThanToken,
	">>":        GreaterThanGreaterThanToken,
	"<":         LessThanToken,
	">":         GreaterThanToken,
	"<=":        LessThanEqualsToken,
	"<=>":       LessThanEqualsGreaterThanToken,
	">=":        GreaterThanEqualsToken,
	"==":        EqualsEqualsToken,
	"===":       EqualsEqualsEqualsToken,
	"!=":        ExclamationEqualsToken,
	"!==":       ExclamationEqualsEqualsToken,
	"^":         CaretToken,
	"|":         BarToken,
	"&":         AmpersandToken,
	"&&":        AmpersandAmpersandToken,
	"||":        BarBarToken,
	"?":         QuestionToken,
	":":         ColonToken,
	"::":        ColonColonToken,
	";":         SemicolonToken,
	"=":         EqualsToken,
	"**=":       AsteriskAsteriskEqualsToken,
	"*=":        AsteriskEqualsToken,
	"/=":        SlashEqualsToken,
	"%=":        PercentEqualsToken,
	"+=":        PlusEqualsToken,
	"-=":        MinusEqualsToken,
	".=":        DotEqualsToken,
	"<<=":       LessThanLessThanEqualsToken,
	">>=":       GreaterThanGreaterThanEqualsToken,
	"&=":        AmpersandEqualsToken,
	"^=":        CaretEqualsToken,
	"|=":        BarEqualsToken,
	",":         CommaToken,
	"??":        QuestionQuestionToken,
	"<:":        LessThanEqualsGreaterThanToken,
	"<>":        LessThanGreaterThanToken,
	"...":       DotDotDotToken,
	"\\":        BackslashToken,
	"<?=":       ScriptSectionStartTag, // TODO, technically not an operator
	"<?php ":    ScriptSectionStartTag, // TODO, technically not an operator
	"<?php\t":   ScriptSectionStartTag, // TODO add tests
	"<?php\n":   ScriptSectionStartTag,
	"<?php\r":   ScriptSectionStartTag,
	"<?php\r\n": ScriptSectionStartTag,
	"?>":        ScriptSectionEndTag, // TODO, technically not an operator
	"?>\n":      ScriptSectionEndTag, // TODO, technically not an operator
	"?>\r\n":    ScriptSectionEndTag, // TODO, technically not an operator
	"?>\r":      ScriptSectionEndTag, // TODO, technically not an operator
	"@":         AtSymbolToken,       // TODO not in spec
	"`":         BacktickToken,
}
var RESERVED_WORDS = map[string]TokenKind{
	// http://php.net/manual/en/reserved.constants.php
	// TRUE, FALSE, NULL are special predefined constants
	// TODO - also consider adding other constants
	"true":  TrueReservedWord,
	"false": FalseReservedWord,
	"null":  NullReservedWord,

	// RESERVED WORDS:
	// http://php.net/manual/en/reserved.other-reserved-words.php
	"int":     IntReservedWord,
	"float":   FloatReservedWord,
	"bool":    BoolReservedWord,
	"string":  StringReservedWord,
	"binary":  BinaryReservedWord,
	"boolean": BooleanReservedWord,
	"double":  DoubleReservedWord,
	"integer": IntegerReservedWord,
	"object":  ObjectReservedWord,
	"real":    RealReservedWord,
	"void":    VoidReservedWord,
}
var KEYWORDS = map[string]TokenKind{
	"abstract":     AbstractKeyword,
	"and":          AndKeyword,
	"array":        ArrayKeyword,
	"as":           AsKeyword,
	"break":        BreakKeyword,
	"callable":     CallableKeyword,
	"case":         CaseKeyword,
	"catch":        CatchKeyword,
	"class":        ClassKeyword,
	"clone":        CloneKeyword,
	"const":        ConstKeyword,
	"continue":     ContinueKeyword,
	"declare":      DeclareKeyword,
	"default":      DefaultKeyword,
	"die":          DieKeyword,
	"do":           DoKeyword,
	"echo":         EchoKeyword,
	"else":         ElseKeyword,
	"elseif":       ElseIfKeyword,
	"empty":        EmptyKeyword,
	"enddeclare":   EndDeclareKeyword,
	"endfor":       EndForKeyword,
	"endforeach":   EndForEachKeyword,
	"endif":        EndIfKeyword,
	"endswitch":    EndSwitchKeyword,
	"endwhile":     EndWhileKeyword,
	"eval":         EvalKeyword,
	"exit":         ExitKeyword,
	"extends":      ExtendsKeyword,
	"final":        FinalKeyword,
	"finally":      FinallyKeyword,
	"for":          ForKeyword,
	"foreach":      ForeachKeyword,
	"function":     FunctionKeyword,
	"global":       GlobalKeyword,
	"goto":         GotoKeyword,
	"if":           IfKeyword,
	"implements":   ImplementsKeyword,
	"include":      IncludeKeyword,
	"include_once": IncludeOnceKeyword,
	"instanceof":   InstanceOfKeyword,
	"insteadof":    InsteadOfKeyword,
	"interface":    InterfaceKeyword,
	"isset":        IsSetKeyword,
	"list":         ListKeyword,
	"namespace":    NamespaceKeyword,
	"new":          NewKeyword,
	"or":           OrKeyword,
	"print":        PrintKeyword,
	"private":      PrivateKeyword,
	"protected":    ProtectedKeyword,
	"public":       PublicKeyword,
	"require":      RequireKeyword,
	"require_once": RequireOnceKeyword,
	"return":       ReturnKeyword,
	"static":       StaticKeyword,
	"switch":       SwitchKeyword,
	"throw":        ThrowKeyword,
	"trait":        TraitKeyword,
	"try":          TryKeyword,
	"unset":        UnsetKeyword,
	"use":          UseKeyword,
	"var":          VarKeyword,
	"while":        WhileKeyword,
	"xor":          XorKeyword,
	"yield":        YieldKeyword,
	"yield from":   YieldFromKeyword,
}

func valueInMap(v TokenKind, m map[string]TokenKind) bool {
	for _, value := range m {
		if v == value {
			return true
		}
	}
	return false
}

func IsReserverdWordToken(v TokenKind) bool {
	return valueInMap(v, RESERVED_WORDS)
}
func IsKeywordToken(v TokenKind) bool {
	return valueInMap(v, KEYWORDS)
}

func IsKeywordOrReserverdWordToken(v TokenKind) bool {
	return IsKeywordToken(v) || IsReserverdWordToken(v)
}

func ReserverTokens() []TokenKind {
	var tokens []TokenKind
	for _, value := range RESERVED_WORDS {
		tokens = append(tokens, value)
	}
	return tokens
}
