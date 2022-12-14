// Code generated by ent, DO NOT EDIT.

package item

import (
	"time"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldIsDisplay holds the string denoting the is_display field in the database.
	FieldIsDisplay = "is_display"
	// FieldTax holds the string denoting the tax field in the database.
	FieldTax = "tax"
	// FieldTaxRate holds the string denoting the tax_rate field in the database.
	FieldTaxRate = "tax_rate"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldTemporaryStock holds the string denoting the temporary_stock field in the database.
	FieldTemporaryStock = "temporary_stock"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// Table holds the table name of the item in the database.
	Table = "items"
	// ImagesTable is the table that holds the images relation/edge. The primary key declared below.
	ImagesTable = "item_images"
	// ImagesInverseTable is the table name for the Images entity.
	// It exists in this package in order to avoid circular dependency with the "images" package.
	ImagesInverseTable = "images"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "items"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_id"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldCategoryID,
	FieldIsDisplay,
	FieldTax,
	FieldTaxRate,
	FieldPrice,
	FieldTemporaryStock,
}

var (
	// ImagesPrimaryKey and ImagesColumn2 are the table columns denoting the
	// primary key for the images relation (M2M).
	ImagesPrimaryKey = []string{"item_id", "images_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultIsDisplay holds the default value on creation for the "is_display" field.
	DefaultIsDisplay bool
)
