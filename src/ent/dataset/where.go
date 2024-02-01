// Code generated by ent, DO NOT EDIT.

package dataset

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KenshiTech/unchained/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DataSet {
	return predicate.DataSet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DataSet {
	return predicate.DataSet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DataSet {
	return predicate.DataSet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DataSet {
	return predicate.DataSet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DataSet {
	return predicate.DataSet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DataSet {
	return predicate.DataSet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DataSet {
	return predicate.DataSet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DataSet {
	return predicate.DataSet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DataSet {
	return predicate.DataSet(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.DataSet {
	return predicate.DataSet(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.DataSet {
	return predicate.DataSet(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.DataSet {
	return predicate.DataSet(sql.FieldContainsFold(FieldName, v))
}

// HasAssetPrice applies the HasEdge predicate on the "assetPrice" edge.
func HasAssetPrice() predicate.DataSet {
	return predicate.DataSet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AssetPriceTable, AssetPricePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetPriceWith applies the HasEdge predicate on the "assetPrice" edge with a given conditions (other predicates).
func HasAssetPriceWith(preds ...predicate.AssetPrice) predicate.DataSet {
	return predicate.DataSet(func(s *sql.Selector) {
		step := newAssetPriceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DataSet) predicate.DataSet {
	return predicate.DataSet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DataSet) predicate.DataSet {
	return predicate.DataSet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DataSet) predicate.DataSet {
	return predicate.DataSet(sql.NotPredicates(p))
}