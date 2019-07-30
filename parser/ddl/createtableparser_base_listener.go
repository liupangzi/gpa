// Code generated from CreateTableParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package ddlparser // CreateTableParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCreateTableParserListener is a complete listener for a parse tree produced by CreateTableParser.
type BaseCreateTableParserListener struct{}

var _ CreateTableParserListener = &BaseCreateTableParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCreateTableParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCreateTableParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCreateTableParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCreateTableParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCreateTableParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCreateTableParserListener) ExitRoot(ctx *RootContext) {}

// EnterCreateTableDDL is called when production createTableDDL is entered.
func (s *BaseCreateTableParserListener) EnterCreateTableDDL(ctx *CreateTableDDLContext) {}

// ExitCreateTableDDL is called when production createTableDDL is exited.
func (s *BaseCreateTableParserListener) ExitCreateTableDDL(ctx *CreateTableDDLContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseCreateTableParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseCreateTableParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseCreateTableParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseCreateTableParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterCreateDefinitions is called when production createDefinitions is entered.
func (s *BaseCreateTableParserListener) EnterCreateDefinitions(ctx *CreateDefinitionsContext) {}

// ExitCreateDefinitions is called when production createDefinitions is exited.
func (s *BaseCreateTableParserListener) ExitCreateDefinitions(ctx *CreateDefinitionsContext) {}

// EnterCreateDefinition is called when production createDefinition is entered.
func (s *BaseCreateTableParserListener) EnterCreateDefinition(ctx *CreateDefinitionContext) {}

// ExitCreateDefinition is called when production createDefinition is exited.
func (s *BaseCreateTableParserListener) ExitCreateDefinition(ctx *CreateDefinitionContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseCreateTableParserListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseCreateTableParserListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseCreateTableParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseCreateTableParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterLengthOneDimension is called when production lengthOneDimension is entered.
func (s *BaseCreateTableParserListener) EnterLengthOneDimension(ctx *LengthOneDimensionContext) {}

// ExitLengthOneDimension is called when production lengthOneDimension is exited.
func (s *BaseCreateTableParserListener) ExitLengthOneDimension(ctx *LengthOneDimensionContext) {}

// EnterLengthTwoDimension is called when production lengthTwoDimension is entered.
func (s *BaseCreateTableParserListener) EnterLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// ExitLengthTwoDimension is called when production lengthTwoDimension is exited.
func (s *BaseCreateTableParserListener) ExitLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// EnterLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is entered.
func (s *BaseCreateTableParserListener) EnterLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// ExitLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is exited.
func (s *BaseCreateTableParserListener) ExitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// EnterCollectionOptions is called when production collectionOptions is entered.
func (s *BaseCreateTableParserListener) EnterCollectionOptions(ctx *CollectionOptionsContext) {}

// ExitCollectionOptions is called when production collectionOptions is exited.
func (s *BaseCreateTableParserListener) ExitCollectionOptions(ctx *CollectionOptionsContext) {}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *BaseCreateTableParserListener) EnterColumnConstraint(ctx *ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *BaseCreateTableParserListener) ExitColumnConstraint(ctx *ColumnConstraintContext) {}

// EnterNullNotNull is called when production nullNotNull is entered.
func (s *BaseCreateTableParserListener) EnterNullNotNull(ctx *NullNotNullContext) {}

// ExitNullNotNull is called when production nullNotNull is exited.
func (s *BaseCreateTableParserListener) ExitNullNotNull(ctx *NullNotNullContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseCreateTableParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseCreateTableParserListener) ExitComment(ctx *CommentContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseCreateTableParserListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseCreateTableParserListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterPrimaryKey is called when production primaryKey is entered.
func (s *BaseCreateTableParserListener) EnterPrimaryKey(ctx *PrimaryKeyContext) {}

// ExitPrimaryKey is called when production primaryKey is exited.
func (s *BaseCreateTableParserListener) ExitPrimaryKey(ctx *PrimaryKeyContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseCreateTableParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseCreateTableParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCreateTableParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCreateTableParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterCurrentTimestamp is called when production currentTimestamp is entered.
func (s *BaseCreateTableParserListener) EnterCurrentTimestamp(ctx *CurrentTimestampContext) {}

// ExitCurrentTimestamp is called when production currentTimestamp is exited.
func (s *BaseCreateTableParserListener) ExitCurrentTimestamp(ctx *CurrentTimestampContext) {}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseCreateTableParserListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseCreateTableParserListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseCreateTableParserListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseCreateTableParserListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseCreateTableParserListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseCreateTableParserListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexColumnNames is called when production indexColumnNames is entered.
func (s *BaseCreateTableParserListener) EnterIndexColumnNames(ctx *IndexColumnNamesContext) {}

// ExitIndexColumnNames is called when production indexColumnNames is exited.
func (s *BaseCreateTableParserListener) ExitIndexColumnNames(ctx *IndexColumnNamesContext) {}

// EnterTableOption is called when production tableOption is entered.
func (s *BaseCreateTableParserListener) EnterTableOption(ctx *TableOptionContext) {}

// ExitTableOption is called when production tableOption is exited.
func (s *BaseCreateTableParserListener) ExitTableOption(ctx *TableOptionContext) {}
