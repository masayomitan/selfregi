// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"selfregi/ent/images"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Images is the model entity for the Images schema.
type Images struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImagesQuery when eager-loading is set.
	Edges ImagesEdges `json:"edges"`
}

// ImagesEdges holds the relations/edges for other nodes in the graph.
type ImagesEdges struct {
	// Items holds the value of the items edge.
	Items []*Item `json:"items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e ImagesEdges) ItemsOrErr() ([]*Item, error) {
	if e.loadedTypes[0] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Images) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case images.FieldID:
			values[i] = new(sql.NullInt64)
		case images.FieldName, images.FieldPath:
			values[i] = new(sql.NullString)
		case images.FieldCreatedAt, images.FieldUpdatedAt, images.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Images", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Images fields.
func (i *Images) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case images.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case images.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case images.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case images.FieldDeletedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[j])
			} else if value.Valid {
				i.DeletedAt = new(time.Time)
				*i.DeletedAt = value.Time
			}
		case images.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case images.FieldPath:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[j])
			} else if value.Valid {
				i.Path = value.String
			}
		}
	}
	return nil
}

// QueryItems queries the "items" edge of the Images entity.
func (i *Images) QueryItems() *ItemQuery {
	return (&ImagesClient{config: i.config}).QueryItems(i)
}

// Update returns a builder for updating this Images.
// Note that you need to call Images.Unwrap() before calling this method if this Images
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Images) Update() *ImagesUpdateOne {
	return (&ImagesClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Images entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Images) Unwrap() *Images {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Images is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Images) String() string {
	var builder strings.Builder
	builder.WriteString("Images(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := i.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(i.Path)
	builder.WriteByte(')')
	return builder.String()
}

// ImagesSlice is a parsable slice of Images.
type ImagesSlice []*Images

func (i ImagesSlice) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
