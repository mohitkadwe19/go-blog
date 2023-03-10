// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlogsColumns holds the columns for the "blogs" table.
	BlogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "user_blogs", Type: field.TypeInt, Nullable: true},
	}
	// BlogsTable holds the schema information for the "blogs" table.
	BlogsTable = &schema.Table{
		Name:       "blogs",
		Columns:    BlogsColumns,
		PrimaryKey: []*schema.Column{BlogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blogs_users_blogs",
				Columns:    []*schema.Column{BlogsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlogsTable,
		UsersTable,
	}
)

func init() {
	BlogsTable.ForeignKeys[0].RefTable = UsersTable
}
