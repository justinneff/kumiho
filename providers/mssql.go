package providers

import "fmt"

type Mssql struct {
}

func (mssql Mssql) GenerateMigration(name string) (string, error) {
	template := `
/*******************************************************************************
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
