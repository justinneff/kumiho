package providers

import (
	"fmt"
)

type Provider interface {
	GenerateMigration(name string) (string, error)
}

func GetProvider(providerType string) (Provider, error) {
	switch providerType {
	case "mssql":
		return Mssql{}, nil
	default:
		return nil, fmt.Errorf("unknown provider type \"%s\"", providerType)
	}
}
