package types

import "github.com/graphql-go/graphql"

func GqString() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
	}
}

func GqInt() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Int,
	}
}

// Строковое поле обязательное для заполенения
func GqNewNonNullString() *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	}
}

// Int поле обязательное для заполенения
func GqNewNonNullInt() *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}
}