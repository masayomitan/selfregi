// Code generated by ent, DO NOT EDIT.

package ent

import (
	"selfregi/ent/account"
	"selfregi/ent/accountdetail"
	"selfregi/ent/admin"
	"selfregi/ent/cart"
	"selfregi/ent/cartdetail"
	"selfregi/ent/categories"
	"selfregi/ent/images"
	"selfregi/ent/item"
	"selfregi/ent/journals"
	"selfregi/ent/schema"
	"selfregi/ent/visitor"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountMixin := schema.Account{}.Mixin()
	accountMixinFields0 := accountMixin[0].Fields()
	_ = accountMixinFields0
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountMixinFields0[0].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountMixinFields0[1].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescVisitorID is the schema descriptor for visitor_id field.
	accountDescVisitorID := accountFields[0].Descriptor()
	// account.VisitorIDValidator is a validator for the "visitor_id" field. It is called by the builders before save.
	account.VisitorIDValidator = accountDescVisitorID.Validators[0].(func(int) error)
	accountdetailMixin := schema.AccountDetail{}.Mixin()
	accountdetailMixinFields0 := accountdetailMixin[0].Fields()
	_ = accountdetailMixinFields0
	accountdetailFields := schema.AccountDetail{}.Fields()
	_ = accountdetailFields
	// accountdetailDescCreatedAt is the schema descriptor for created_at field.
	accountdetailDescCreatedAt := accountdetailMixinFields0[0].Descriptor()
	// accountdetail.DefaultCreatedAt holds the default value on creation for the created_at field.
	accountdetail.DefaultCreatedAt = accountdetailDescCreatedAt.Default.(func() time.Time)
	// accountdetailDescUpdatedAt is the schema descriptor for updated_at field.
	accountdetailDescUpdatedAt := accountdetailMixinFields0[1].Descriptor()
	// accountdetail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	accountdetail.DefaultUpdatedAt = accountdetailDescUpdatedAt.Default.(func() time.Time)
	// accountdetail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	accountdetail.UpdateDefaultUpdatedAt = accountdetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountdetailDescAccountID is the schema descriptor for account_id field.
	accountdetailDescAccountID := accountdetailFields[0].Descriptor()
	// accountdetail.AccountIDValidator is a validator for the "account_id" field. It is called by the builders before save.
	accountdetail.AccountIDValidator = accountdetailDescAccountID.Validators[0].(func(int) error)
	// accountdetailDescVisitorID is the schema descriptor for visitor_id field.
	accountdetailDescVisitorID := accountdetailFields[1].Descriptor()
	// accountdetail.VisitorIDValidator is a validator for the "visitor_id" field. It is called by the builders before save.
	accountdetail.VisitorIDValidator = accountdetailDescVisitorID.Validators[0].(func(int) error)
	// accountdetailDescProductID is the schema descriptor for product_id field.
	accountdetailDescProductID := accountdetailFields[2].Descriptor()
	// accountdetail.ProductIDValidator is a validator for the "product_id" field. It is called by the builders before save.
	accountdetail.ProductIDValidator = accountdetailDescProductID.Validators[0].(func(int) error)
	// accountdetailDescCategoryID is the schema descriptor for category_id field.
	accountdetailDescCategoryID := accountdetailFields[3].Descriptor()
	// accountdetail.CategoryIDValidator is a validator for the "category_id" field. It is called by the builders before save.
	accountdetail.CategoryIDValidator = accountdetailDescCategoryID.Validators[0].(func(int) error)
	// accountdetailDescQuantity is the schema descriptor for quantity field.
	accountdetailDescQuantity := accountdetailFields[6].Descriptor()
	// accountdetail.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	accountdetail.QuantityValidator = accountdetailDescQuantity.Validators[0].(func(int) error)
	// accountdetailDescDiscountID is the schema descriptor for discount_id field.
	accountdetailDescDiscountID := accountdetailFields[10].Descriptor()
	// accountdetail.DiscountIDValidator is a validator for the "discount_id" field. It is called by the builders before save.
	accountdetail.DiscountIDValidator = accountdetailDescDiscountID.Validators[0].(func(int) error)
	adminMixin := schema.Admin{}.Mixin()
	adminMixinFields0 := adminMixin[0].Fields()
	_ = adminMixinFields0
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminMixinFields0[0].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminMixinFields0[1].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() time.Time)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() time.Time)
	// adminDescName is the schema descriptor for name field.
	adminDescName := adminFields[0].Descriptor()
	// admin.DefaultName holds the default value on creation for the name field.
	admin.DefaultName = adminDescName.Default.(string)
	cartMixin := schema.Cart{}.Mixin()
	cartMixinFields0 := cartMixin[0].Fields()
	_ = cartMixinFields0
	cartFields := schema.Cart{}.Fields()
	_ = cartFields
	// cartDescCreatedAt is the schema descriptor for created_at field.
	cartDescCreatedAt := cartMixinFields0[0].Descriptor()
	// cart.DefaultCreatedAt holds the default value on creation for the created_at field.
	cart.DefaultCreatedAt = cartDescCreatedAt.Default.(func() time.Time)
	// cartDescUpdatedAt is the schema descriptor for updated_at field.
	cartDescUpdatedAt := cartMixinFields0[1].Descriptor()
	// cart.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cart.DefaultUpdatedAt = cartDescUpdatedAt.Default.(func() time.Time)
	// cart.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cart.UpdateDefaultUpdatedAt = cartDescUpdatedAt.UpdateDefault.(func() time.Time)
	// cartDescVisitorID is the schema descriptor for visitor_id field.
	cartDescVisitorID := cartFields[0].Descriptor()
	// cart.VisitorIDValidator is a validator for the "visitor_id" field. It is called by the builders before save.
	cart.VisitorIDValidator = cartDescVisitorID.Validators[0].(func(int) error)
	cartdetailMixin := schema.CartDetail{}.Mixin()
	cartdetailMixinFields0 := cartdetailMixin[0].Fields()
	_ = cartdetailMixinFields0
	cartdetailFields := schema.CartDetail{}.Fields()
	_ = cartdetailFields
	// cartdetailDescCreatedAt is the schema descriptor for created_at field.
	cartdetailDescCreatedAt := cartdetailMixinFields0[0].Descriptor()
	// cartdetail.DefaultCreatedAt holds the default value on creation for the created_at field.
	cartdetail.DefaultCreatedAt = cartdetailDescCreatedAt.Default.(func() time.Time)
	// cartdetailDescUpdatedAt is the schema descriptor for updated_at field.
	cartdetailDescUpdatedAt := cartdetailMixinFields0[1].Descriptor()
	// cartdetail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cartdetail.DefaultUpdatedAt = cartdetailDescUpdatedAt.Default.(func() time.Time)
	// cartdetail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cartdetail.UpdateDefaultUpdatedAt = cartdetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// cartdetailDescAccountID is the schema descriptor for account_id field.
	cartdetailDescAccountID := cartdetailFields[0].Descriptor()
	// cartdetail.AccountIDValidator is a validator for the "account_id" field. It is called by the builders before save.
	cartdetail.AccountIDValidator = cartdetailDescAccountID.Validators[0].(func(int) error)
	// cartdetailDescVisitorID is the schema descriptor for visitor_id field.
	cartdetailDescVisitorID := cartdetailFields[1].Descriptor()
	// cartdetail.VisitorIDValidator is a validator for the "visitor_id" field. It is called by the builders before save.
	cartdetail.VisitorIDValidator = cartdetailDescVisitorID.Validators[0].(func(int) error)
	// cartdetailDescProductID is the schema descriptor for product_id field.
	cartdetailDescProductID := cartdetailFields[2].Descriptor()
	// cartdetail.ProductIDValidator is a validator for the "product_id" field. It is called by the builders before save.
	cartdetail.ProductIDValidator = cartdetailDescProductID.Validators[0].(func(int) error)
	// cartdetailDescCategoryID is the schema descriptor for category_id field.
	cartdetailDescCategoryID := cartdetailFields[3].Descriptor()
	// cartdetail.CategoryIDValidator is a validator for the "category_id" field. It is called by the builders before save.
	cartdetail.CategoryIDValidator = cartdetailDescCategoryID.Validators[0].(func(int) error)
	// cartdetailDescQuantity is the schema descriptor for quantity field.
	cartdetailDescQuantity := cartdetailFields[6].Descriptor()
	// cartdetail.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	cartdetail.QuantityValidator = cartdetailDescQuantity.Validators[0].(func(int) error)
	// cartdetailDescDiscountID is the schema descriptor for discount_id field.
	cartdetailDescDiscountID := cartdetailFields[10].Descriptor()
	// cartdetail.DiscountIDValidator is a validator for the "discount_id" field. It is called by the builders before save.
	cartdetail.DiscountIDValidator = cartdetailDescDiscountID.Validators[0].(func(int) error)
	categoriesMixin := schema.Categories{}.Mixin()
	categoriesMixinFields0 := categoriesMixin[0].Fields()
	_ = categoriesMixinFields0
	categoriesFields := schema.Categories{}.Fields()
	_ = categoriesFields
	// categoriesDescCreatedAt is the schema descriptor for created_at field.
	categoriesDescCreatedAt := categoriesMixinFields0[0].Descriptor()
	// categories.DefaultCreatedAt holds the default value on creation for the created_at field.
	categories.DefaultCreatedAt = categoriesDescCreatedAt.Default.(func() time.Time)
	// categoriesDescUpdatedAt is the schema descriptor for updated_at field.
	categoriesDescUpdatedAt := categoriesMixinFields0[1].Descriptor()
	// categories.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	categories.DefaultUpdatedAt = categoriesDescUpdatedAt.Default.(func() time.Time)
	// categories.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	categories.UpdateDefaultUpdatedAt = categoriesDescUpdatedAt.UpdateDefault.(func() time.Time)
	imagesMixin := schema.Images{}.Mixin()
	imagesMixinFields0 := imagesMixin[0].Fields()
	_ = imagesMixinFields0
	imagesFields := schema.Images{}.Fields()
	_ = imagesFields
	// imagesDescCreatedAt is the schema descriptor for created_at field.
	imagesDescCreatedAt := imagesMixinFields0[0].Descriptor()
	// images.DefaultCreatedAt holds the default value on creation for the created_at field.
	images.DefaultCreatedAt = imagesDescCreatedAt.Default.(func() time.Time)
	// imagesDescUpdatedAt is the schema descriptor for updated_at field.
	imagesDescUpdatedAt := imagesMixinFields0[1].Descriptor()
	// images.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	images.DefaultUpdatedAt = imagesDescUpdatedAt.Default.(func() time.Time)
	// images.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	images.UpdateDefaultUpdatedAt = imagesDescUpdatedAt.UpdateDefault.(func() time.Time)
	itemMixin := schema.Item{}.Mixin()
	itemMixinFields0 := itemMixin[0].Fields()
	_ = itemMixinFields0
	itemFields := schema.Item{}.Fields()
	_ = itemFields
	// itemDescCreatedAt is the schema descriptor for created_at field.
	itemDescCreatedAt := itemMixinFields0[0].Descriptor()
	// item.DefaultCreatedAt holds the default value on creation for the created_at field.
	item.DefaultCreatedAt = itemDescCreatedAt.Default.(func() time.Time)
	// itemDescUpdatedAt is the schema descriptor for updated_at field.
	itemDescUpdatedAt := itemMixinFields0[1].Descriptor()
	// item.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	item.DefaultUpdatedAt = itemDescUpdatedAt.Default.(func() time.Time)
	// item.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	item.UpdateDefaultUpdatedAt = itemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// itemDescName is the schema descriptor for name field.
	itemDescName := itemFields[0].Descriptor()
	// item.DefaultName holds the default value on creation for the name field.
	item.DefaultName = itemDescName.Default.(string)
	journalsMixin := schema.Journals{}.Mixin()
	journalsMixinFields0 := journalsMixin[0].Fields()
	_ = journalsMixinFields0
	journalsFields := schema.Journals{}.Fields()
	_ = journalsFields
	// journalsDescCreatedAt is the schema descriptor for created_at field.
	journalsDescCreatedAt := journalsMixinFields0[0].Descriptor()
	// journals.DefaultCreatedAt holds the default value on creation for the created_at field.
	journals.DefaultCreatedAt = journalsDescCreatedAt.Default.(func() time.Time)
	// journalsDescUpdatedAt is the schema descriptor for updated_at field.
	journalsDescUpdatedAt := journalsMixinFields0[1].Descriptor()
	// journals.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	journals.DefaultUpdatedAt = journalsDescUpdatedAt.Default.(func() time.Time)
	// journals.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	journals.UpdateDefaultUpdatedAt = journalsDescUpdatedAt.UpdateDefault.(func() time.Time)
	visitorMixin := schema.Visitor{}.Mixin()
	visitorMixinFields0 := visitorMixin[0].Fields()
	_ = visitorMixinFields0
	visitorFields := schema.Visitor{}.Fields()
	_ = visitorFields
	// visitorDescCreatedAt is the schema descriptor for created_at field.
	visitorDescCreatedAt := visitorMixinFields0[0].Descriptor()
	// visitor.DefaultCreatedAt holds the default value on creation for the created_at field.
	visitor.DefaultCreatedAt = visitorDescCreatedAt.Default.(func() time.Time)
	// visitorDescUpdatedAt is the schema descriptor for updated_at field.
	visitorDescUpdatedAt := visitorMixinFields0[1].Descriptor()
	// visitor.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	visitor.DefaultUpdatedAt = visitorDescUpdatedAt.Default.(func() time.Time)
	// visitor.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	visitor.UpdateDefaultUpdatedAt = visitorDescUpdatedAt.UpdateDefault.(func() time.Time)
	// visitorDescName is the schema descriptor for name field.
	visitorDescName := visitorFields[0].Descriptor()
	// visitor.DefaultName holds the default value on creation for the name field.
	visitor.DefaultName = visitorDescName.Default.(string)
}
