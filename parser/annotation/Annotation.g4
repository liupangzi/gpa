// antlr -Dlanguage=Go Annotation.g4 -o . -package annotation

grammar Annotation;

root
    : esql
    | es EOF
    ;

es
    : At ES LeftParenthesis URLS Equal LeftBracket Literal (Comma Literal)* RightBracket RightParenthesis
    ;

esql
    : At ESql LeftParenthesis SQL Equal Literal RightParenthesis
    ;

At
    : '@'
    ;

Equal
    : '='
    ;

Comma
    : ','
    ;

Literal
    : '"' ( '\\"' | UnicodeCharacter )*? '"'
    | '\'' ( '\\\'' | UnicodeCharacter )*? '\''
    ;

Integer
    : ('+' | '-')? [1-9] DecimalDigit*
    | ('+' | '-')? DecimalDigit
    ;

Number
    : ('+' | '-')? [1-9] DecimalDigit* ( '.' DecimalDigit+ )?
    | ('+' | '-')? DecimalDigit ( '.' DecimalDigit+ )?
    ;

BackQuotedString
    : '`' RawString '`'
    ;

PlaceholderString
    : '{' RawString '}'
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
