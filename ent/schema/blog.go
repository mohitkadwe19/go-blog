package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("content").NotEmpty(),
		field.String("Author").NotEmpty(),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		//Create an inverse-edge called "owner" of type `User`
		//this is called the backref(Inverse edge)
		// and reference it to the "blogs" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("owner", User.Type).Ref("blogs").
			//Here we are telling th ent framework that a Blog can only have one user who owned it
			Unique(),
	}
}
