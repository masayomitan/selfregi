// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"selfregi/ent/cart"
	"selfregi/ent/visitor"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Cart is the model entity for the Cart schema.
type Cart struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VisitorID holds the value of the "visitor_id" field.
	VisitorID int `json:"visitor_id,omitempty"`
	// Status holds the value of the "status" field.
	Status *uint `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CartQuery when eager-loading is set.
	Edges         CartEdges `json:"edges"`
	visitor_carts *int
}

// CartEdges holds the relations/edges for other nodes in the graph.
type CartEdges struct {
	// CartDetails holds the value of the cart_details edge.
	CartDetails []*CartDetail `json:"cart_details,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *Visitor `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CartDetailsOrErr returns the CartDetails value or an error if the edge
// was not loaded in eager-loading.
func (e CartEdges) CartDetailsOrErr() ([]*CartDetail, error) {
	if e.loadedTypes[0] {
		return e.CartDetails, nil
	}
	return nil, &NotLoadedError{edge: "cart_details"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartEdges) OwnerOrErr() (*Visitor, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: visitor.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cart) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cart.FieldID, cart.FieldVisitorID, cart.FieldStatus:
			values[i] = new(sql.NullInt64)
		case cart.FieldCreatedAt, cart.FieldUpdatedAt, cart.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case cart.ForeignKeys[0]: // visitor_carts
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cart", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cart fields.
func (c *Cart) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cart.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cart.FieldVisitorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field visitor_id", values[i])
			} else if value.Valid {
				c.VisitorID = int(value.Int64)
			}
		case cart.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = new(uint)
				*c.Status = uint(value.Int64)
			}
		case cart.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cart.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case cart.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = new(time.Time)
				*c.DeletedAt = value.Time
			}
		case cart.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field visitor_carts", value)
			} else if value.Valid {
				c.visitor_carts = new(int)
				*c.visitor_carts = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCartDetails queries the "cart_details" edge of the Cart entity.
func (c *Cart) QueryCartDetails() *CartDetailQuery {
	return (&CartClient{config: c.config}).QueryCartDetails(c)
}

// QueryOwner queries the "owner" edge of the Cart entity.
func (c *Cart) QueryOwner() *VisitorQuery {
	return (&CartClient{config: c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Cart.
// Note that you need to call Cart.Unwrap() before calling this method if this Cart
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cart) Update() *CartUpdateOne {
	return (&CartClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cart) Unwrap() *Cart {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cart is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cart) String() string {
	var builder strings.Builder
	builder.WriteString("Cart(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("visitor_id=")
	builder.WriteString(fmt.Sprintf("%v", c.VisitorID))
	builder.WriteString(", ")
	if v := c.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Carts is a parsable slice of Cart.
type Carts []*Cart

func (c Carts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}