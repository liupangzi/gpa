// Code generated from CreateTableParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package ddlparser // CreateTableParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CreateTableParserListener is a complete listener for a parse tree produced by CreateTableParser.
type CreateTableParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterCreateTableDDL is called when entering the createTableDDL production.
	EnterCreateTableDDL(c *CreateTableDDLContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterIfNotExists is called when entering the ifNotExists production.
	EnterIfNotExists(c *IfNotExistsContext)

	// EnterCreateDefinitions is called when entering the createDefinitions production.
	EnterCreateDefinitions(c *CreateDefinitionsContext)

	// EnterCreateDefinition is called when entering the createDefinition production.
	EnterCreateDefinition(c *CreateDefinitionContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterLengthOneDimension is called when entering the lengthOneDimension production.
	EnterLengthOneDimension(c *LengthOneDimensionContext)

	// EnterLengthTwoDimension is called when entering the lengthTwoDimension production.
	EnterLengthTwoDimension(c *LengthTwoDimensionContext)

	// EnterLengthTwoOptionalDimension is called when entering the lengthTwoOptionalDimension production.
	EnterLengthTwoOptionalDimension(c *LengthTwoOptionalDimensionContext)

	// EnterCollectionOptions is called when entering the collectionOptions production.
	EnterCollectionOptions(c *CollectionOptionsContext)

	// EnterColumnConstraint is called when entering the columnConstraint production.
	EnterColumnConstraint(c *ColumnConstraintContext)

	// EnterNullNotNull is called when entering the nullNotNull production.
	EnterNullNotNull(c *NullNotNullContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterPrimaryKey is called when entering the primaryKey production.
	EnterPrimaryKey(c *PrimaryKeyContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterCurrentTimestamp is called when entering the currentTimestamp production.
	EnterCurrentTimestamp(c *CurrentTimestampContext)

	// EnterTableConstraint is called when entering the tableConstraint production.
	EnterTableConstraint(c *TableConstraintContext)

	// EnterIndexOption is called when entering the indexOption production.
	EnterIndexOption(c *IndexOptionContext)

	// EnterIndexType is called when entering the indexType production.
	EnterIndexType(c *IndexTypeContext)

	// EnterIndexColumnNames is called when entering the indexColumnNames production.
	EnterIndexColumnNames(c *IndexColumnNamesContext)

	// EnterTableOption is called when entering the tableOption production.
	EnterTableOption(c *TableOptionContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitCreateTableDDL is called when exiting the createTableDDL production.
	ExitCreateTableDDL(c *CreateTableDDLContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitIfNotExists is called when exiting the ifNotExists production.
	ExitIfNotExists(c *IfNotExistsContext)

	// ExitCreateDefinitions is called when exiting the createDefinitions production.
	ExitCreateDefinitions(c *CreateDefinitionsContext)

	// ExitCreateDefinition is called when exiting the createDefinition production.
	ExitCreateDefinition(c *CreateDefinitionContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitLengthOneDimension is called when exiting the lengthOneDimension production.
	ExitLengthOneDimension(c *LengthOneDimensionContext)

	// ExitLengthTwoDimension is called when exiting the lengthTwoDimension production.
	ExitLengthTwoDimension(c *LengthTwoDimensionContext)

	// ExitLengthTwoOptionalDimension is called when exiting the lengthTwoOptionalDimension production.
	ExitLengthTwoOptionalDimension(c *LengthTwoOptionalDimensionContext)

	// ExitCollectionOptions is called when exiting the collectionOptions production.
	ExitCollectionOptions(c *CollectionOptionsContext)

	// ExitColumnConstraint is called when exiting the columnConstraint production.
	ExitColumnConstraint(c *ColumnConstraintContext)

	// ExitNullNotNull is called when exiting the nullNotNull production.
	ExitNullNotNull(c *NullNotNullContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitPrimaryKey is called when exiting the primaryKey production.
	ExitPrimaryKey(c *PrimaryKeyContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitCurrentTimestamp is called when exiting the currentTimestamp production.
	ExitCurrentTimestamp(c *CurrentTimestampContext)

	// ExitTableConstraint is called when exiting the tableConstraint production.
	ExitTableConstraint(c *TableConstraintContext)

	// ExitIndexOption is called when exiting the indexOption production.
	ExitIndexOption(c *IndexOptionContext)

	// ExitIndexType is called when exiting the indexType production.
	ExitIndexType(c *IndexTypeContext)

	// ExitIndexColumnNames is called when exiting the indexColumnNames production.
	ExitIndexColumnNames(c *IndexColumnNamesContext)

	// ExitTableOption is called when exiting the tableOption production.
	ExitTableOption(c *TableOptionContext)
}
