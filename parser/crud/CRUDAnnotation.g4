// antlr -Dlanguage=Go CRUDAnnotation.g4 -o . -package crudannotation

grammar CRUDAnnotation;

crud
    : insertClause EOF
    | selectClause EOF
    | updateClause EOF
    | deleteClause EOF
    ;

insertClause
    : At Insert LeftParenthesis RightParenthesis
    ;

selectClause
    : At Select LeftParenthesis sql=Literal RightParenthesis
    ;

updateClause
    : At Update LeftParenthesis sql=Literal RightParenthesis
    ;

deleteClause
    : At Delete LeftParenthesis sql=Literal RightParenthesis
    ;

At
    : '@'
    ;

LeftParenthesis
    : '('
    ;

RightParenthesis
    : ')'
    ;

Insert
    : [iI][nN][sS][eE][rR][tT]
    ;

Select
    : [sS][eE][lL][eE][cC][tT]
    ;

Update
    : [uU][pP][dD][aA][tT][eE]
    ;

Delete
    : [dD][eE][lL][eE][tT][eE]
    ;

Literal
    : '"' ( '\\"' | UnicodeCharacter )*? '"'
    | '\'' ( '\\\'' | UnicodeCharacter )*? '\''
    ;

fragment UnicodeCharacter
    : ~[\u000A-\u001F]
    | ~[\u007F-\u009F]
    ;

WS
    : [ \t\r\n]+ -> skip
    ;
