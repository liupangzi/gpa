// antlr -Dlanguage=Go CreateTableParser.g4 -o . -package ddlparser

parser grammar CreateTableParser;

options { tokenVocab=DDLLexer; }

root
    : createTableDDL (Semicolon createTableDDL)* Semicolon? EOF
    ;

createTableDDL
    : CREATE TEMPORARY? TABLE ifNotExists? tableName
      createDefinitions
      ( tableOption (Comma? tableOption)* )?
    ;

tableName
    : BackQuotedString
    | RawString
    ;

ifNotExists
    : IF NOT EXISTS
    ;

createDefinitions
    : LeftParenthesis createDefinition (Comma createDefinition)* RightParenthesis
    ;

createDefinition
    : field=(BackQuotedString | RawString) columnDefinition
    | tableConstraint
    ;

columnDefinition
    : dataType columnConstraint*
    ;

dataType
    : typeName=(CHAR | VARCHAR | TINYTEXT | TEXT | MEDIUMTEXT | LONGTEXT | NCHAR | NVARCHAR)
      lengthOneDimension? BINARY? ((CHARACTER SET | CHARSET) (BINARY | RawString | BackQuotedString))? (COLLATE (BackQuotedString | RawString))?
    | NATIONAL typeName=(VARCHAR | CHARACTER) lengthOneDimension? BINARY?
    | NCHAR typeName=VARCHAR lengthOneDimension? BINARY?
    | NATIONAL typeName=(CHAR | CHARACTER) VARYING lengthOneDimension? BINARY?
    | typeName=(TINYINT | SMALLINT | MEDIUMINT | INT | INTEGER | BIGINT) lengthOneDimension? sign=(SIGNED | UNSIGNED)? ZEROFILL?
    | typeName=REAL lengthTwoDimension? sign=(SIGNED | UNSIGNED)? ZEROFILL?
    | typeName=DOUBLE PRECISION? lengthTwoDimension? sign=(SIGNED | UNSIGNED)? ZEROFILL?
    | typeName=(DECIMAL | DEC | FIXED | NUMERIC | FLOAT) lengthTwoOptionalDimension? sign=(SIGNED | UNSIGNED)? ZEROFILL?
    | typeName=(DATE | TINYBLOB | BLOB | MEDIUMBLOB | LONGBLOB | BOOL | BOOLEAN | SERIAL | JSON)
    | typeName=(BIT | TIME | TIMESTAMP | DATETIME | BINARY | VARBINARY | YEAR) lengthOneDimension?
    | typeName=(ENUM | SET) collectionOptions BINARY? ((CHARACTER SET | CHARSET) (BINARY | RawString | BackQuotedString))?
    ;

lengthOneDimension
    : LeftParenthesis Integer RightParenthesis
    ;

lengthTwoDimension
    : LeftParenthesis Integer Comma Integer RightParenthesis
    ;

lengthTwoOptionalDimension
    : LeftParenthesis Integer (Comma Integer)? RightParenthesis
    ;

collectionOptions
    : LeftParenthesis Literal (Comma Literal)* RightParenthesis
    ;

columnConstraint
    : nullNotNull
    | DEFAULT defaultValue
    | (AUTO_INCREMENT | ON UPDATE currentTimestamp)
    | primaryKey
    | KEY
    | UNIQUE KEY?
    | comment
    | COLLATE BackQuotedString
    | SERIAL DEFAULT VALUE
    ;

nullNotNull
    : NOT? NULL
    ;

comment
    : COMMENT content=Literal
    ;

defaultValue
    : NULL
    | unaryOperator? constant
    | currentTimestamp (ON UPDATE currentTimestamp)?
    ;

primaryKey
    : PRIMARY KEY
    ;

unaryOperator
    : Exclamation | Tilde | Plus | Minus | NOT
    ;

constant
    : Literal | Integer
    | Minus Integer
    | (TRUE | FALSE)
    | Number
    | NOT? NULL
    ;

currentTimestamp
    :
    (
      (CURRENT_TIMESTAMP | LOCALTIME | LOCALTIMESTAMP) (LeftParenthesis Integer? RightParenthesis)?
      | NOW LeftParenthesis Integer? RightParenthesis
    )
    ;

tableConstraint
    : (CONSTRAINT name=(Literal | RawString)?)?
      pk=primaryKey index=(Literal | RawString)? (BackQuotedString | RawString)?
      indexColumnNames indexOption*
    | (CONSTRAINT name=(Literal | RawString)?)?
      UNIQUE? indexFormat=(INDEX | KEY)? index=(BackQuotedString | RawString)?
      indexType? indexColumnNames indexOption*
    ;

indexOption
    : KEY_BLOCK_SIZE Equal? FilesizeLiteral
    | WITH PARSER (Literal | RawString)
    | COMMENT Literal
    ;

indexType
    : USING (BTREE | HASH)
    ;

indexColumnNames
    : LeftParenthesis
        (BackQuotedString | RawString) (LeftParenthesis Integer RightParenthesis)?
        (Comma (BackQuotedString | RawString) (LeftParenthesis Integer RightParenthesis)?)*
     RightParenthesis
    ;

tableOption
    : ENGINE Equal? (BackQuotedString | RawString)
    | AUTO_INCREMENT Equal? Integer
    | DEFAULT? (CHARACTER SET | CHARSET) Equal? (BackQuotedString | RawString)
    | DEFAULT? COLLATE Equal? (BackQuotedString | RawString)
    | COMMENT Equal? tableComment=Literal
    ;
