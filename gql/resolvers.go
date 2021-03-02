package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/thospol/go-graphql/repository"
	"gorm.io/gorm"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	database       *gorm.DB
	userRepository repository.UserRepository
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string

	name, ok := p.Args["name"].(string)
	if ok {
		return r.userRepository.FindByName(r.database, name)
	}

	return nil, nil
}
