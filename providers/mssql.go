package providers

import "fmt"

type Mssql struct {
}

func (mssql Mssql) GenerateMigration(name string) (string, error) {
	template := `/*******************************************************************************
* Migration: %s
*******************************************************************************/
/*
	Write your migration here for example:
	CREATE TABLE [dbo].[MyAwesomeTable] (
		[id] INT IDENTITY(1,1) NOT NULL,
		[name] NVARCHAR(100) NOT NULL,
		CONSTRAINT [PK_MyAwesomeTable] PRIMARY KEY CLUSTERED ([id])
	)
*/
`
	return fmt.Sprintf(template, name), nil
}

func (mssql Mssql) GenerateProcedure(schema, name string) (string, error) {
	template := `CREATE PROCEDURE [%s].[%s]
(
	@param1 INT,
	@param2 INT
)
AS
BEGIN

	SET NOCOUNT ON;

END
`
	return fmt.Sprintf(template, schema, name), nil
}

func (mssql Mssql) GenerateScalarFunction(schema, name string) (string, error) {
	template := `CREATE FUNCTION [%s].[%s]
(
	@param1 INT
)
RETURNS INT
AS
BEGIN

	DECLARE @m_Value INT;

	SET @m_Value = @param1;

	return @m_Value;

END
`
	return fmt.Sprintf(template, schema, name), nil
}

func (mssql Mssql) GenerateTableFunction(schema, name string) (string, error) {
	template := `CREATE FUNCTION [%s].[%s]
(
	@param1 INT
)
RETURNS TABLE
AS
RETURN (
	SELECT
		*
	FROM
		[dbo].[SomeTableOrView]
);
`
	return fmt.Sprintf(template, schema, name), nil
}

func (mssql Mssql) GenerateView(schema, name string) (string, error) {
	template := `CREATE VIEW [%s].[%s]
AS
SELECT
	*
FROM
	[dbo].[SomeTableOrView];
`
	return fmt.Sprintf(template, schema, name), nil
}

func (mssql Mssql) ResolveSchema(schema string) (string, error) {
	if len(schema) == 0 {
		return "dbo", nil
	} else {
		return schema, nil
	}
}
