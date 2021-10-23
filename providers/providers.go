package providers

import (
	"fmt"

	"github.com/justinneff/kumiho/entities"
)

func GetProvider(providerType string) (entities.Provider, error) {
	switch providerType {
	case "mssql":
		return Mssql{}, nil
	default:
		return nil, fmt.Errorf("unknown provider type \"%s\"", providerType)
	}
}
