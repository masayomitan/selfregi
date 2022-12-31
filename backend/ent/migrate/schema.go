// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "visitor_id", Type: field.TypeInt, Nullable: true},
		{Name: "status", Type: field.TypeUint},
		{Name: "subtotal", Type: field.TypeInt},
		{Name: "total", Type: field.TypeInt},
		{Name: "tax", Type: field.TypeInt},
		{Name: "tax_rate", Type: field.TypeInt},
		{Name: "discount_class", Type: field.TypeUint},
		{Name: "discount_rate", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(10,2)"}},
		{Name: "discount_price", Type: field.TypeInt},
		{Name: "paid_price", Type: field.TypeInt},
		{Name: "change", Type: field.TypeInt},
		{Name: "visitor_managed_accounts", Type: field.TypeInt},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_visitors_managed_accounts",
				Columns:    []*schema.Column{AccountsColumns[15]},
				RefColumns: []*schema.Column{VisitorsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AccountDetailsColumns holds the columns for the "account_details" table.
	AccountDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "account_id", Type: field.TypeInt, Nullable: true},
		{Name: "visitor_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
		{Name: "category_id", Type: field.TypeInt, Nullable: true},
		{Name: "data", Type: field.TypeJSON, Nullable: true},
		{Name: "options", Type: field.TypeJSON, Nullable: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "price", Type: field.TypeInt},
		{Name: "tax", Type: field.TypeInt},
		{Name: "tax_rate", Type: field.TypeInt},
		{Name: "discount_id", Type: field.TypeInt, Nullable: true},
		{Name: "discount_name", Type: field.TypeString},
		{Name: "discount_class", Type: field.TypeUint},
		{Name: "discount_rate", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(10,2)"}},
		{Name: "discount_price", Type: field.TypeInt},
		{Name: "account_managed_account_details", Type: field.TypeInt, Nullable: true},
	}
	// AccountDetailsTable holds the schema information for the "account_details" table.
	AccountDetailsTable = &schema.Table{
		Name:       "account_details",
		Columns:    AccountDetailsColumns,
		PrimaryKey: []*schema.Column{AccountDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_details_accounts_managed_account_details",
				Columns:    []*schema.Column{AccountDetailsColumns[19]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "password", Type: field.TypeString},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:       "admins",
		Columns:    AdminsColumns,
		PrimaryKey: []*schema.Column{AdminsColumns[0]},
	}
	// CartsColumns holds the columns for the "carts" table.
	CartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "visitor_id", Type: field.TypeInt, Nullable: true},
		{Name: "status", Type: field.TypeUint},
		{Name: "visitor_carts", Type: field.TypeInt},
	}
	// CartsTable holds the schema information for the "carts" table.
	CartsTable = &schema.Table{
		Name:       "carts",
		Columns:    CartsColumns,
		PrimaryKey: []*schema.Column{CartsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "carts_visitors_carts",
				Columns:    []*schema.Column{CartsColumns[6]},
				RefColumns: []*schema.Column{VisitorsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CartDetailsColumns holds the columns for the "cart_details" table.
	CartDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "account_id", Type: field.TypeInt, Nullable: true},
		{Name: "visitor_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
		{Name: "category_id", Type: field.TypeInt, Nullable: true},
		{Name: "data", Type: field.TypeJSON, Nullable: true},
		{Name: "options", Type: field.TypeJSON, Nullable: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "price", Type: field.TypeInt},
		{Name: "tax", Type: field.TypeInt},
		{Name: "tax_rate", Type: field.TypeInt},
		{Name: "discount_id", Type: field.TypeInt, Nullable: true},
		{Name: "discount_name", Type: field.TypeString},
		{Name: "discount_class", Type: field.TypeUint},
		{Name: "discount_rate", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(10,2)"}},
		{Name: "discount_price", Type: field.TypeInt},
		{Name: "cart_cart_details", Type: field.TypeInt, Nullable: true},
	}
	// CartDetailsTable holds the schema information for the "cart_details" table.
	CartDetailsTable = &schema.Table{
		Name:       "cart_details",
		Columns:    CartDetailsColumns,
		PrimaryKey: []*schema.Column{CartDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cart_details_carts_cart_details",
				Columns:    []*schema.Column{CartDetailsColumns[19]},
				RefColumns: []*schema.Column{CartsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "is_display", Type: field.TypeInt},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "path", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(255)"}},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "category_id", Type: field.TypeInt},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "items_categories_items",
				Columns:    []*schema.Column{ItemsColumns[5]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// JournalsColumns holds the columns for the "journals" table.
	JournalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// JournalsTable holds the schema information for the "journals" table.
	JournalsTable = &schema.Table{
		Name:       "journals",
		Columns:    JournalsColumns,
		PrimaryKey: []*schema.Column{JournalsColumns[0]},
	}
	// VisitorsColumns holds the columns for the "visitors" table.
	VisitorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "sex", Type: field.TypeInt},
	}
	// VisitorsTable holds the schema information for the "visitors" table.
	VisitorsTable = &schema.Table{
		Name:       "visitors",
		Columns:    VisitorsColumns,
		PrimaryKey: []*schema.Column{VisitorsColumns[0]},
	}
	// ItemImagesColumns holds the columns for the "item_images" table.
	ItemImagesColumns = []*schema.Column{
		{Name: "item_id", Type: field.TypeInt},
		{Name: "images_id", Type: field.TypeInt},
	}
	// ItemImagesTable holds the schema information for the "item_images" table.
	ItemImagesTable = &schema.Table{
		Name:       "item_images",
		Columns:    ItemImagesColumns,
		PrimaryKey: []*schema.Column{ItemImagesColumns[0], ItemImagesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "item_images_item_id",
				Columns:    []*schema.Column{ItemImagesColumns[0]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "item_images_images_id",
				Columns:    []*schema.Column{ItemImagesColumns[1]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AccountDetailsTable,
		AdminsTable,
		CartsTable,
		CartDetailsTable,
		CategoriesTable,
		ImagesTable,
		ItemsTable,
		JournalsTable,
		VisitorsTable,
		ItemImagesTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = VisitorsTable
	AccountDetailsTable.ForeignKeys[0].RefTable = AccountsTable
	CartsTable.ForeignKeys[0].RefTable = VisitorsTable
	CartDetailsTable.ForeignKeys[0].RefTable = CartsTable
	ItemsTable.ForeignKeys[0].RefTable = CategoriesTable
	ItemImagesTable.ForeignKeys[0].RefTable = ItemsTable
	ItemImagesTable.ForeignKeys[1].RefTable = ImagesTable
}
