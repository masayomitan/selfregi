package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"entgo.io/ent/schema/edge"
	"time"
)

// AccountDetail holds the schema definition for the AccountDetail entity.
type AccountDetail struct {
	ent.Schema
}

// Fields of the AccountDetail.
func (AccountDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id").
			Positive().
			Optional(),
		field.Int("visitor_id").
      Positive().
			Optional(),
		field.Int("product_id").
      Positive().
			Optional(),
		field.Int("category_id").
      Positive().
			Optional(),
		field.JSON("data", []string{}).
			Optional(),
		field.JSON("options", []string{}).
			Optional(),
		field.Int("quantity").
			Positive(),
		field.Int("price").
			Nillable(),
		field.Int("tax").
			Nillable(),
		field.Int("tax_rate").
			Nillable(),
		field.Int("discount_id").
			Positive().
			Optional().
			Nillable(),
		field.String("discount_name").
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

		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Nillable(),
	}
}

// Edges of the AccountDetail.
func (AccountDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
				Ref("account_details").
				Unique(),
		edge.From("item", Item.Type).
				Ref("account_details").
				Unique(),
	}
}
