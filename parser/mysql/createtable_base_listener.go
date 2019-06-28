// Code generated from CreateTable.g4 by ANTLR 4.7.2. DO NOT EDIT.

package mysql // CreateTable
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCreateTableListener is a complete listener for a parse tree produced by CreateTableParser.
type BaseCreateTableListener struct{}

var _ CreateTableListener = &BaseCreateTableListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCreateTableListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCreateTableListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCreateTableListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCreateTableListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCreateTableListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCreateTableListener) ExitRoot(ctx *RootContext) {}

// EnterCreateTableDDL is called when production createTableDDL is entered.
func (s *BaseCreateTableListener) EnterCreateTableDDL(ctx *CreateTableDDLContext) {}

// ExitCreateTableDDL is called when production createTableDDL is exited.
func (s *BaseCreateTableListener) ExitCreateTableDDL(ctx *CreateTableDDLContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseCreateTableListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseCreateTableListener) ExitTableName(ctx *TableNameContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseCreateTableListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseCreateTableListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterCreateDefinitions is called when production createDefinitions is entered.
func (s *BaseCreateTableListener) EnterCreateDefinitions(ctx *CreateDefinitionsContext) {}

// ExitCreateDefinitions is called when production createDefinitions is exited.
func (s *BaseCreateTableListener) ExitCreateDefinitions(ctx *CreateDefinitionsContext) {}

// EnterCreateDefinition is called when production createDefinition is entered.
func (s *BaseCreateTableListener) EnterCreateDefinition(ctx *CreateDefinitionContext) {}

// ExitCreateDefinition is called when production createDefinition is exited.
func (s *BaseCreateTableListener) ExitCreateDefinition(ctx *CreateDefinitionContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseCreateTableListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseCreateTableListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseCreateTableListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseCreateTableListener) ExitDataType(ctx *DataTypeContext) {}

// EnterLengthOneDimension is called when production lengthOneDimension is entered.
func (s *BaseCreateTableListener) EnterLengthOneDimension(ctx *LengthOneDimensionContext) {}

// ExitLengthOneDimension is called when production lengthOneDimension is exited.
func (s *BaseCreateTableListener) ExitLengthOneDimension(ctx *LengthOneDimensionContext) {}

// EnterLengthTwoDimension is called when production lengthTwoDimension is entered.
func (s *BaseCreateTableListener) EnterLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// ExitLengthTwoDimension is called when production lengthTwoDimension is exited.
func (s *BaseCreateTableListener) ExitLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// EnterLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is entered.
func (s *BaseCreateTableListener) EnterLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// ExitLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is exited.
func (s *BaseCreateTableListener) ExitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// EnterCollectionOptions is called when production collectionOptions is entered.
func (s *BaseCreateTableListener) EnterCollectionOptions(ctx *CollectionOptionsContext) {}

// ExitCollectionOptions is called when production collectionOptions is exited.
func (s *BaseCreateTableListener) ExitCollectionOptions(ctx *CollectionOptionsContext) {}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *BaseCreateTableListener) EnterColumnConstraint(ctx *ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *BaseCreateTableListener) ExitColumnConstraint(ctx *ColumnConstraintContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseCreateTableListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseCreateTableListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseCreateTableListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseCreateTableListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCreateTableListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCreateTableListener) ExitConstant(ctx *ConstantContext) {}

// EnterCurrentTimestamp is called when production currentTimestamp is entered.
func (s *BaseCreateTableListener) EnterCurrentTimestamp(ctx *CurrentTimestampContext) {}

// ExitCurrentTimestamp is called when production currentTimestamp is exited.
func (s *BaseCreateTableListener) ExitCurrentTimestamp(ctx *CurrentTimestampContext) {}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseCreateTableListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseCreateTableListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseCreateTableListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseCreateTableListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseCreateTableListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseCreateTableListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexColumnNames is called when production indexColumnNames is entered.
func (s *BaseCreateTableListener) EnterIndexColumnNames(ctx *IndexColumnNamesContext) {}

// ExitIndexColumnNames is called when production indexColumnNames is exited.
func (s *BaseCreateTableListener) ExitIndexColumnNames(ctx *IndexColumnNamesContext) {}

// EnterTableOption is called when production tableOption is entered.
func (s *BaseCreateTableListener) EnterTableOption(ctx *TableOptionContext) {}

// ExitTableOption is called when production tableOption is exited.
func (s *BaseCreateTableListener) ExitTableOption(ctx *TableOptionContext) {}
