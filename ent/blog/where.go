// Code generated by ent, DO NOT EDIT.

package blog

import (
	"Blog/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldContent, v))
}

// Author applies equality check predicate on the "Author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldAuthor, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldContent, v))
}

// AuthorEQ applies the EQ predicate on the "Author" field.
func AuthorEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "Author" field.
func AuthorNEQ(v string) predicate.Blog {
	return predicate.Blog(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "Author" field.
func AuthorIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "Author" field.
func AuthorNotIn(vs ...string) predicate.Blog {
	return predicate.Blog(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "Author" field.
func AuthorGT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "Author" field.
func AuthorGTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "Author" field.
func AuthorLT(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "Author" field.
func AuthorLTE(v string) predicate.Blog {
	return predicate.Blog(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "Author" field.
func AuthorContains(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "Author" field.
func AuthorHasPrefix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "Author" field.
func AuthorHasSuffix(v string) predicate.Blog {
	return predicate.Blog(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorEqualFold applies the EqualFold predicate on the "Author" field.
func AuthorEqualFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "Author" field.
func AuthorContainsFold(v string) predicate.Blog {
	return predicate.Blog(sql.FieldContainsFold(FieldAuthor, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		p(s.Not())
	})
}