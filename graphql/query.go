package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql/language/ast"

	"github.com/graphql-go/graphql"
)

/**
GetSelectedFields Read the requested fields in Resolve
Ref: https://github.com/graphql-go/graphql/issues/125, https://github.com/graphql-go/graphql/issues/157
*/
func GetSelectedFields(params graphql.ResolveParams) (map[string]interface{}, error) {
	fieldASTs := params.Info.FieldASTs
	if len(fieldASTs) == 0 {
		return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
	}
	return selectedFieldsFromSelections(params, fieldASTs[0].SelectionSet.Selections)
}

func selectedFieldsFromSelections(params graphql.ResolveParams, selections []ast.Selection) (selected map[string]interface{}, err error) {
	selected = map[string]interface{}{}

	for _, s := range selections {
		switch s := s.(type) {
		case *ast.Field:
			if s.SelectionSet == nil {
				selected[s.Name.Value] = true
			} else {
				selected[s.Name.Value], err = selectedFieldsFromSelections(params, s.SelectionSet.Selections)
				if err != nil {
					return
				}
			}
		case *ast.FragmentSpread:
			n := s.Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				err = fmt.Errorf("getSelectedFields: no fragment found with name %v", n)

				return
			}
			selected[s.Name.Value], err = selectedFieldsFromSelections(params, frag.GetSelectionSet().Selections)
			if err != nil {
				return
			}
		default:
			err = fmt.Errorf("getSelectedFields: found unexpected selection type %v", s)

			return
		}
	}

	return
}
