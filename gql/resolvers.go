package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/simple-gql-api/postgres"
)

type Resolver struct {
	db *postgres.Db
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)
	if ok {
		if users, err := r.db.GetUsersByName(name); err != nil {
			return nil, err
		}
		return users, nil
	}

	return nil, fmt.Errorf("expected name to be of type string, was %T", name)
}
