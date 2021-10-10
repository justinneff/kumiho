package entities

type Provider interface {
	GenerateMigration(name string) (string, error)
	GeneratePostDeploy(name string) (string, error)
	GeneratePreDeploy(name string) (string, error)
	GenerateProcedure(schema, name string) (string, error)
	GenerateScalarFunction(schema, name string) (string, error)
	GenerateTableFunction(schema, name string) (string, error)
	GenerateView(schema, name string) (string, error)
	GetObjectSchemaAndName(content []byte) (string, string)
	IsDependency(content []byte, schema, name string) (bool, error)
	ResolveSchema(schema string) (string, error)
}
