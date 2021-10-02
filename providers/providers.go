package providers

import (
	"fmt"
)

type Provider interface {
	GenerateMigration(name string) (string, error)
	GenerateProcedure(schema, name string) (string, error)
	GenerateScalarFunction(schema, name string) (string, error)
	GenerateTableFunction(schema, name string) (string, error)
	GenerateView(schema, name string) (string, error)
	ResolveSchema(schema string) (string, error)
}

func GetProvider(providerType string) (Provider, error) {
	switch providerType {
	case "mssql":
		return Mssql{}, nil
	default:
		return nil, fmt.Errorf("unknown provider type \"%s\"", providerType)
	}
}
