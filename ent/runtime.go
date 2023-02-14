// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aodebiao/ent/article"
	"aodebiao/ent/schema"
	"aodebiao/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleMixin := schema.Article{}.Mixin()
	articleMixinFields0 := articleMixin[0].Fields()
	_ = articleMixinFields0
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleMixinFields0[0].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(func() time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleMixinFields0[1].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(func() time.Time)
	// article.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	article.UpdateDefaultUpdatedAt = articleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// articleDescDelFlag is the schema descriptor for del_flag field.
	articleDescDelFlag := articleMixinFields0[3].Descriptor()
	// article.DefaultDelFlag holds the default value on creation for the del_flag field.
	article.DefaultDelFlag = articleDescDelFlag.Default.(int)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescDelFlag is the schema descriptor for del_flag field.
	userDescDelFlag := userMixinFields0[3].Descriptor()
	// user.DefaultDelFlag holds the default value on creation for the del_flag field.
	user.DefaultDelFlag = userDescDelFlag.Default.(int)
}
