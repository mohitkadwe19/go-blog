// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Blog/ent/blog"
	"Blog/ent/schema"
	"Blog/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	blogFields := schema.Blog{}.Fields()
	_ = blogFields
	// blogDescTitle is the schema descriptor for title field.
	blogDescTitle := blogFields[0].Descriptor()
	// blog.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	blog.TitleValidator = blogDescTitle.Validators[0].(func(string) error)
	// blogDescContent is the schema descriptor for content field.
	blogDescContent := blogFields[1].Descriptor()
	// blog.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	blog.ContentValidator = blogDescContent.Validators[0].(func(string) error)
	// blogDescAuthor is the schema descriptor for Author field.
	blogDescAuthor := blogFields[2].Descriptor()
	// blog.AuthorValidator is a validator for the "Author" field. It is called by the builders before save.
	blog.AuthorValidator = blogDescAuthor.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}