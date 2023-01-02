// Code generated by ent, DO NOT EDIT.

package accountdetail

import (
	"time"
)

const (
	// Label holds the string label denoting the accountdetail type in the database.
	Label = "account_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldVisitorID holds the string denoting the visitor_id field in the database.
	FieldVisitorID = "visitor_id"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldOptions holds the string denoting the options field in the database.
	FieldOptions = "options"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldTax holds the string denoting the tax field in the database.
	FieldTax = "tax"
	// FieldTaxRate holds the string denoting the tax_rate field in the database.
	FieldTaxRate = "tax_rate"
	// FieldDiscountID holds the string denoting the discount_id field in the database.
	FieldDiscountID = "discount_id"
	// FieldDiscountName holds the string denoting the discount_name field in the database.
	FieldDiscountName = "discount_name"
	// FieldDiscountClass holds the string denoting the discount_class field in the database.
	FieldDiscountClass = "discount_class"
	// FieldDiscountRate holds the string denoting the discount_rate field in the database.
	FieldDiscountRate = "discount_rate"
	// FieldDiscountPrice holds the string denoting the discount_price field in the database.
	FieldDiscountPrice = "discount_price"
	// EdgeManagedAccount holds the string denoting the managed_account edge name in mutations.
	EdgeManagedAccount = "managed_account"
	// Table holds the table name of the accountdetail in the database.
	Table = "account_details"
	// ManagedAccountTable is the table that holds the managed_account relation/edge.
	ManagedAccountTable = "account_details"
	// ManagedAccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	ManagedAccountInverseTable = "accounts"
	// ManagedAccountColumn is the table column denoting the managed_account relation/edge.
	ManagedAccountColumn = "account_managed_account_details"
)

// Columns holds all SQL columns for accountdetail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAccountID,
	FieldVisitorID,
	FieldProductID,
	FieldCategoryID,
	FieldData,
	FieldOptions,
	FieldQuantity,
	FieldPrice,
	FieldTax,
	FieldTaxRate,
	FieldDiscountID,
	FieldDiscountName,
	FieldDiscountClass,
	FieldDiscountRate,
	FieldDiscountPrice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "account_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_managed_account_details",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// AccountIDValidator is a validator for the "account_id" field. It is called by the builders before save.
	AccountIDValidator func(int) error
	// VisitorIDValidator is a validator for the "visitor_id" field. It is called by the builders before save.
	VisitorIDValidator func(int) error
	// ProductIDValidator is a validator for the "product_id" field. It is called by the builders before save.
	ProductIDValidator func(int) error
	// CategoryIDValidator is a validator for the "category_id" field. It is called by the builders before save.
	CategoryIDValidator func(int) error
	// QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	QuantityValidator func(int) error
	// DiscountIDValidator is a validator for the "discount_id" field. It is called by the builders before save.
	DiscountIDValidator func(int) error
)