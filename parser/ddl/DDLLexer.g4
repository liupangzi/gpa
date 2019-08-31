// antlr -Dlanguage=Go DDLLexer.g4 -o . -package ddlparser

lexer grammar DDLLexer;

channels {
    INLINE_COMMENT_CHANNEL
}

AUTO_INCREMENT
    : [aA][uU][tT][oO]'_'[iI][nN][cC][rR][eE][mM][eE][nN][tT]
    ;

BIGINT
    : [bB][iI][gG][iI][nN][tT]
    ;

BINARY
    : [bB][iI][nN][aA][rR][yY]
    ;

BIT
    : [bB][iI][tT]
    ;

BLOB
    : [bB][lL][oO][bB]
    ;

BTREE
    : [bB][tT][rR][eE][eE]
    ;

BOOLEAN
    : [bB][oO][oO][lL][eE][aA][nN]
    ;

BOOL
    : [bB][oO][oO][lL]
    ;

CHARACTER
    : [cC][hH][aA][rR][aA][cC][tT][eE][rR]
    ;

CHARSET
    : [cC][hH][aA][rR][sS][eE][tT]
    ;

CHAR
    : [cC][hH][aA][rR]
    ;

COLLATE
    : [cC][oO][lL][lL][aA][tT][eE]
    ;

COMMENT
    : [cC][oO][mM][mM][eE][nN][tT]
    ;

CONSTRAINT
    : [cC][oO][nN][sS][tT][rR][aA][iI][nN][tT]
    ;

CREATE
    : [cC][rR][eE][aA][tT][eE]
    ;

CURRENT_TIMESTAMP
    : [cC][uU][rR][rR][eE][nN][tT]'_'[tT][iI][mM][eE][sS][tT][aA][mM][pP]
    ;

DATETIME
    : [dD][aA][tT][eE][tT][iI][mM][eE]
    ;

DATE
    : [dD][aA][tT][eE]
    ;

DECIMAL
    : [dD][eE][cC][iI][mM][aA][lL]
    ;

DEC
    : [dD][eE][cC]
    ;

DEFAULT
    : [dD][eE][fF][aA][uU][lL][tT]
    ;

DOUBLE
    : [dD][oO][uU][bB][lL][eE]
    ;

ENGINE
    : [eE][nN][gG][iI][nN][eE]
    ;

ENUM
    : [eE][nN][uU][mM]
    ;

EXISTS
    : [eE][xX][iI][sS][tT][sS]
    ;

FALSE
    : [fF][aA][lL][sS][eE]
    ;

FIXED
    : [fF][iI][xX][eE][dD]
    ;

FLOAT
    : [fF][lL][oO][aA][tT]
    ;

HASH
    : [hH][aA][sS][hH]
    ;

INDEX
    : [iI][nN][dD][eE][xX]
    ;

IF
    : [iI][fF]
    ;

INTEGER
    : [iI][nN][tT][eE][gG][eE][rR]
    ;

INT
    : [iI][nN][tT]
    ;

JSON
    : [jJ][sS][oO][nN]
    ;

KEY_BLOCK_SIZE
    : [kK][eE][yY]'_'[bB][lL][oO][cC][kK]'_'[sS][iI][zZ][eE]
    ;

KEY
    : [kK][eE][yY]
    ;

LOCALTIMESTAMP
    : [lL][oO][cC][aA][lL][tT][iI][mM][eE][sS][tT][aA][mM][pP]
    ;

LOCALTIME
    : [lL][oO][cC][aA][lL][tT][iI][mM][eE]
    ;

LONGBLOB
    : [lL][oO][nN][gG][bB][lL][oO][bB]
    ;

LONGTEXT
    : [lL][oO][nN][gG][tT][eE][xX][tT]
    ;

MEDIUMBLOB
    : [mM][eE][dD][iI][uU][mM][bB][lL][oO][bB]
    ;

MEDIUMINT
    : [mM][eE][dD][iI][uU][mM][iI][nN][tT]
    ;

MEDIUMTEXT
    : [mM][eE][dD][iI][uU][mM][tT][eE][xX][tT]
    ;

NATIONAL
    : [nN][aA][tT][iI][oO][nN][aA][lL]
    ;

NCHAR
    : [nN][cC][hH][aA][rR]
    ;

NOT
    : [nN][oO][tT]
    ;

NOW
    : [nN][oO][wW]
    ;

NULL
    : [nN][uU][lL][lL]
    ;

NUMERIC
    : [nN][uU][mM][eE][rR][iI][cC]
    ;

NVARCHAR
    : [nN][vV][aA][rR][cC][hH][aA][rR]
    ;

ON
    : [oO][nN]
    ;

PARSER
    : [pP][aA][rR][sS][eE][rR]
    ;

PRECISION
    : [pP][rR][eE][cC][iI][sS][iI][oO][nN]
    ;

PRIMARY
    : [pP][rR][iI][mM][aA][rR][yY]
    ;

REAL
    : [rR][eE][aA][lL]
    ;

SERIAL
    : [sS][eE][rR][iI][aA][lL]
    ;

SET
    : [sS][eE][tT]
    ;

SIGNED
    : [sS][iI][gG][nN][eE][dD]
    ;

SMALLINT
    : [sS][mM][aA][lL][lL][iI][nN][tT]
    ;

TABLE
    : [tT][aA][bB][lL][eE]
    ;

TEMPORARY
    : [tT][eE][mM][pP][oO][rR][aA][rR][yY]
    ;

TEXT
    : [tT][eE][xX][tT]
    ;

TIMESTAMP
    : [tT][iI][mM][eE][sS][tT][aA][mM][pP]
    ;

TIME
    : [tT][iI][mM][eE]
    ;

TINYBLOB
    : [tT][iI][nN][yY][bB][lL][oO][bB]
    ;

TINYINT
    : [tT][iI][nN][yY][iI][nN][tT]
    ;

TINYTEXT
    : [tT][iI][nN][yY][tT][eE][xX][tT]
    ;

TRUE
    : [tT][rR][uU][eE]
    ;

UNIQUE
    : [uU][nN][iI][qQ][uU][eE]
    ;

UNSIGNED
    : [uU][nN][sS][iI][gG][nN][eE][dD]
    ;

UPDATE
    : [uU][pP][dD][aA][tT][eE]
    ;

USING
    : [uU][sS][iI][nN][gG]
    ;

VALUE
    : [vV][aA][lL][uU][eE]
    ;

VARBINARY
    : [vV][aA][rR][bB][iI][nN][aA][rR][yY]
    ;

VARCHAR
    : [vV][aA][rR][cC][hH][aA][rR]
    ;

VARYING
    : [vV][aA][rR][yY][iI][nN][gG]
    ;

WITH
    : [wW][iI][tT][hH]
    ;

YEAR
    : [yY][eE][aA][rR]
    ;

ZEROFILL
    : [zZ][eE][rR][oO][fF][iI][lL][lL]
    ;

Equal
    : '='
    ;

Semicolon
    : ';'
    ;

Comma
    : ','
    ;

LeftParenthesis
    : '('
    ;

RightParenthesis
    : ')'
    ;

Tilde
    : '~'
    ;

Exclamation
    : '!'
    ;

Plus
    : '+'
    ;

Minus
    : '-'
    ;

Integer
    : (Plus | Minus)? [1-9] DECIMAL_DIGIT*
    | (Plus | Minus)? DECIMAL_DIGIT
    ;

Number
    : (Plus | Minus)? [1-9] DECIMAL_DIGIT* ( '.' DECIMAL_DIGIT+ )?
    | (Plus | Minus)? DECIMAL_DIGIT ( '.' DECIMAL_DIGIT+ )?
    ;

FilesizeLiteral
    : DECIMAL_DIGIT+ ('K'|'M'|'G'|'T')
    ;

Literal
    : DQUOTA_STRING
    | SQUOTA_STRING
    ;

PlaceholderString
    : '{' RawString '}'
    ;

BackQuotedString
    : '`' ( '\\'. | ~('`'|'\\'))* '`'
    ;

RawString
    : (SINGLE_CHAR | DECIMAL_DIGIT)+
    ;

fragment DECIMAL_DIGIT
    : [0-9]
    ;

fragment SINGLE_CHAR
    : [a-zA-Z]
    | '_'
    | '-'
    ;

fragment DQUOTA_STRING
    : '"' ( '\\'. | '""' | ~('"'| '\\') )* '"'
    ;

fragment SQUOTA_STRING
    : '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\''
    ;

SPACE
    : [ \t\r\n]+ -> channel(HIDDEN)
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
    ;

LINE_COMMENT
    : (
        ('-- ' | '#') ~[\r\n]* ('\r'? '\n' | EOF)
        | '--' ('\r'? '\n' | EOF)
      ) -> channel(INLINE_COMMENT_CHANNEL);
