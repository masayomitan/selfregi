// Code generated by ent, DO NOT EDIT.

package visitor

import (
	"selfregi/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Sex applies equality check predicate on the "sex" field. It's identical to SexEQ.
func Sex(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSex), v))
	})
}

// BillStatus applies equality check predicate on the "bill_status" field. It's identical to BillStatusEQ.
func BillStatus(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBillStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SexEQ applies the EQ predicate on the "sex" field.
func SexEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSex), v))
	})
}

// SexNEQ applies the NEQ predicate on the "sex" field.
func SexNEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSex), v))
	})
}

// SexIn applies the In predicate on the "sex" field.
func SexIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSex), v...))
	})
}

// SexNotIn applies the NotIn predicate on the "sex" field.
func SexNotIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSex), v...))
	})
}

// SexGT applies the GT predicate on the "sex" field.
func SexGT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSex), v))
	})
}

// SexGTE applies the GTE predicate on the "sex" field.
func SexGTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSex), v))
	})
}

// SexLT applies the LT predicate on the "sex" field.
func SexLT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSex), v))
	})
}

// SexLTE applies the LTE predicate on the "sex" field.
func SexLTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSex), v))
	})
}

// BillStatusEQ applies the EQ predicate on the "bill_status" field.
func BillStatusEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBillStatus), v))
	})
}

// BillStatusNEQ applies the NEQ predicate on the "bill_status" field.
func BillStatusNEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBillStatus), v))
	})
}

// BillStatusIn applies the In predicate on the "bill_status" field.
func BillStatusIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBillStatus), v...))
	})
}

// BillStatusNotIn applies the NotIn predicate on the "bill_status" field.
func BillStatusNotIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBillStatus), v...))
	})
}

// BillStatusGT applies the GT predicate on the "bill_status" field.
func BillStatusGT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBillStatus), v))
	})
}

// BillStatusGTE applies the GTE predicate on the "bill_status" field.
func BillStatusGTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBillStatus), v))
	})
}

// BillStatusLT applies the LT predicate on the "bill_status" field.
func BillStatusLT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBillStatus), v))
	})
}

// BillStatusLTE applies the LTE predicate on the "bill_status" field.
func BillStatusLTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBillStatus), v))
	})
}

// HasManagedAccounts applies the HasEdge predicate on the "managed_accounts" edge.
func HasManagedAccounts() predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ManagedAccountsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ManagedAccountsTable, ManagedAccountsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasManagedAccountsWith applies the HasEdge predicate on the "managed_accounts" edge with a given conditions (other predicates).
func HasManagedAccountsWith(preds ...predicate.Account) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ManagedAccountsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ManagedAccountsTable, ManagedAccountsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCarts applies the HasEdge predicate on the "carts" edge.
func HasCarts() predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CartsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CartsTable, CartsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCartsWith applies the HasEdge predicate on the "carts" edge with a given conditions (other predicates).
func HasCartsWith(preds ...predicate.Cart) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CartsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CartsTable, CartsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Visitor) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Visitor) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
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
func Not(p predicate.Visitor) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		p(s.Not())
	})
}
