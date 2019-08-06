// antlr -Dlanguage=Go SelectStatementParser.g4 -o . -package selectparser

grammar SelectStatementParser;

statement
    : selectStatement? Semicolon? EOF
    ;

selectStatement
    : Select (Asterisk | selectField) fromClause whereClause? groupByClause? orderByClause? limitClause?
    ;

selectField
    : ((selectAsPrefix)? (RawString | BackQuotedString)) (Comma (((selectAsPrefix)? (RawString | BackQuotedString))))*
    ;

selectAsPrefix
    : groupFunction LeftParenthesis (RawString | BackQuotedString) RightParenthesis As
    ;

fromClause
    : From table=(BackQuotedString | RawString)
    ;

whereClause
    : Where orExpression
    ;

groupByClause
	: Group By groupByField (Comma groupByField)* (Having atomExpression ((And | Or) atomExpression)*)?
	;

groupByField
    : field=(BackQuotedString | RawString)
    ;

orderByClause
    : Order By orderByField (Comma orderByField)*
    ;

orderByField
    : field=(BackQuotedString | RawString) order=(Asc | Desc)?
    ;

limitClause
    : Limit
    (
      (offsetValue Comma)? limitValue
      | limitValue Offset offsetValue
    )
    ;

offsetValue
    : Number | QuestionMark
    ;

limitValue
    : Number | QuestionMark
    ;

orExpression
    : (subExpression | andExpression) (Or (subExpression | andExpression))*
    ;

andExpression
    : (subExpression | atomExpression) (And (subExpression | atomExpression))*
    ;

subExpression
    : LeftParenthesis (andExpression | orExpression) RightParenthesis
    ;

atomExpression
    : field=(BackQuotedString | RawString) Like val=(Literal | QuestionMark)
    | field=(BackQuotedString | RawString) Is Not? Null
    | field=(BackQuotedString | RawString) Not? In LeftParenthesis inExpression RightParenthesis
    | field=(BackQuotedString | RawString) op=(Equal |  NotEqual) val=(Number | Literal | QuestionMark)
    | field=(BackQuotedString | RawString) op=(GT | LT | GTE | LTE) val=(Number | QuestionMark)
    ;

inExpression
    : (Number | QuestionMark | Literal) (Comma (Number | QuestionMark | Literal))*
    ;

groupFunction
    : Count
    | Sum
    | Max
    | Min
    | Avg
    ;

Is
	: [iI][sS]
	;

In
    : [iI][nN]
    ;

Not
	: [nN][oO][tT]
	;

Null
	: [nN][uU][lL][lL]
	;

LeftParenthesis
    : '('
    ;

RightParenthesis
    : ')'
    ;

Semicolon
    : ';'
    ;

Comma
    : ','
    ;

Asterisk
    : '*'
    ;

Equal
    : '='
    ;

NotEqual
    : '!='
    ;

QuestionMark
    : '?'
    ;

GT
    : '>'
    ;

GTE
    : '>='
    ;

LT
    : '<'
    ;

LTE
    : '<='
    ;

Select
    : [sS][eE][lL][eE][cC][tT]
    ;

Delete
	: [Dd][Ee][Ll][Ee][Tt][Ee]
	;

From
    : [fF][rR][oO][mM]
    ;

Where
    : [wW][hH][eE][rR][eE]
    ;

Having
    : [hH][aA][vV][iI][nN][gG]
    ;

Like
    : [lL][iI][kK][eE]
    ;

And
    : [aA][nN][dD]
    ;

Or
    : [oO][rR]
    ;

Order
    : [oO][rR][dD][eE][rR]
    ;

Group
	: [Gg][Rr][Oo][Uu][Pp]
	;

By
    : [bB][yY]
    ;

Asc
    : [aA][sS][cC]
    ;

Desc
    : [dD][eE][sS][cC]
    ;

Offset
    : [oO][fF][fF][sS][eE][tT]
    ;

Limit
    : [lL][iI][mM][iI][tT]
    ;

As
	: [Aa][Ss]
	;

Count
    : [cC][oO][uU][nN][tT]
    ;

Sum
    : [sS][uU][mM]
    ;

Max
    : [mM][aA][xX]
    ;

Min
    : [mM][iI][nN]
    ;

Avg
    : [aA][vV][gG]
    ;

Number
    : ('+' | '-')? [1-9] DecimalDigit* ( '.' DecimalDigit+ )?
    | ('+' | '-')? DecimalDigit ( '.' DecimalDigit+ )?
    ;

Literal
    : '"' ( '\\"' | UnicodeCharacter )*? '"'
    | '\'' ( '\\\'' | UnicodeCharacter )*? '\''
    ;

BackQuotedString
    : '`' RawString '`'
    ;

RawString
    : (Character | DecimalDigit)+
    ;

fragment DecimalDigit
    : [0-9]
    ;

fragment Character
    : [a-zA-Z]
    | '_'
    | '-'
    ;

fragment UnicodeCharacter
    : ~[\u000A-\u001F]
    | ~[\u007F-\u009F]
    ;

WS
    : [ \t\r\n]+ -> skip
    ;
