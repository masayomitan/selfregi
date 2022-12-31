package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"entgo.io/ent/schema/edge"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
// paid_price change
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("visitor_id").
      Positive().
			Optional(),
		field.Uint("status").
			Nillable(),
		field.Int("subtotal").
			Nillable(),
		field.Int("total").
			Nillable(),
		field.Int("tax").
			Nillable(),
		field.Int("tax_rate").
			Nillable(),
		field.Uint("discount_class").
			Nillable(),
		field.Float("discount_rate").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
					dialect.MySQL:    "decimal(10,2)",
			}).
			Nillable(),
		field.Int("discount_price").
			Nillable(),
		field.Int("paid_price").
			Nillable(),
		field.Int("change").
			Nillable(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("managed_account_details", AccountDetail.Type),
		edge.From("owner", Visitor.Type).
			Ref("managed_accounts").
			Unique().
			Required(),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}