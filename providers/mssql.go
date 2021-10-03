package providers

import (
	"fmt"
	"regexp"
	"strings"
)

type Mssql struct {
}

const mssqlDefaultSchema = "dbo"

func (mssql Mssql) GenerateMigration(name string) (string, error) {
	template := `/*******************************************************************************
* Migration: %s
*******************************************************************************/
-- Write your migration here for example:
-- CREATE TABLE [dbo].[MyAwesomeTable] (
-- 	[id] INT IDENTITY(1,1) NOT NULL
-- 		CONSTRAINT [PK_MyAwesomeTable] PRIMARY KEY CLUSTERED,
-- 	[name] NVARCHAR(100) NOT NULL
-- )
`
	return fmt.Sprintf(template, name), nil
}

func (mssql Mssql) GeneratePostDeploy(name string) (string, error) {
	template := `/*******************************************************************************
* Post-Deployment: %s
*******************************************************************************/

-- Write your post-deployment script here for example:
`
	return fmt.Sprintf(template, name), nil
}

func (mssql Mssql) GeneratePreDeploy(name string) (string, error) {
	template := `/*******************************************************************************
* Pre-Deployment: %s
*******************************************************************************/

-- Write your pre-deployment script here for example:
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

func (mssql Mssql) GetObjectSchemaAndName(content []byte) (schema, name string) {
	r := regexp.MustCompile(`(?i)CREATE\s+(?:FUNCTION|PROCEDURE|VIEW)\s+(\[?[\w\.\[\]]+)`)
	matches := r.FindSubmatch(content)
	if len(matches) == 2 {
		fullName := strings.ReplaceAll(strings.ReplaceAll(string(matches[1]), "[", ""), "]", "")
		parts := strings.Split(fullName, ".")
		if len(parts) == 1 {
			schema = mssqlDefaultSchema
			name = parts[0]
		} else {
			schema = parts[0]
			name = parts[1]
		}
	}
	return
}

func (mssql Mssql) ResolveSchema(schema string) (string, error) {
	if len(schema) == 0 {
		return mssqlDefaultSchema, nil
	} else {
		return schema, nil
	}
}
