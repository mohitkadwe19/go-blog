package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// ent frame work has validator build in
		// the code below indicate that we define a field with type of string
		// named "email" and we validate it as required/ not empty
		field.String("email").
			NotEmpty(),

		// Giving the field name of type string to have default value of unkown if no value is supplied
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//To() function will take two parameter (name or the relation, to where the relation is pointed)
		// defining one to many(one = Current entity which is "User")
		// many will be the entity pass in to the Second Parameter("Todo").
		edge.To("blogs", Blog.Type),
	}
}
