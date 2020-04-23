package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	uuid "github.com/satori/go.uuid"
)

func coerceUUID(value interface{}) interface{} {
	switch value := value.(type) {
	case string:
		return value
	case *string:
		return *value
	}
	return nil
}

var UUID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "UUID",
	Description: "The UUID scalar to check string is UUID format",
	Serialize:   coerceUUID,
	ParseValue:  coerceUUID,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			if uuid, err := uuid.FromString(valueAST.Value); err == nil {
				return uuid.String()
			}
		}
		return nil
	},
})
