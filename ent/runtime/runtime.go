// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/modelname"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/passwordtoken"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/schema"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	modelnameFields := schema.ModelName{}.Fields()
	_ = modelnameFields
	// modelnameDescFieldName is the schema descriptor for field_name field.
	modelnameDescFieldName := modelnameFields[0].Descriptor()
	// modelname.FieldNameValidator is a validator for the "field_name" field. It is called by the builders before save.
	modelname.FieldNameValidator = modelnameDescFieldName.Validators[0].(func(string) error)
	// modelnameDescCreatedAt is the schema descriptor for created_at field.
	modelnameDescCreatedAt := modelnameFields[1].Descriptor()
	// modelname.DefaultCreatedAt holds the default value on creation for the created_at field.
	modelname.DefaultCreatedAt = modelnameDescCreatedAt.Default.(func() time.Time)
	passwordtokenFields := schema.PasswordToken{}.Fields()
	_ = passwordtokenFields
	// passwordtokenDescHash is the schema descriptor for hash field.
	passwordtokenDescHash := passwordtokenFields[0].Descriptor()
	// passwordtoken.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	passwordtoken.HashValidator = passwordtokenDescHash.Validators[0].(func(string) error)
	// passwordtokenDescCreatedAt is the schema descriptor for created_at field.
	passwordtokenDescCreatedAt := passwordtokenFields[1].Descriptor()
	// passwordtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	passwordtoken.DefaultCreatedAt = passwordtokenDescCreatedAt.Default.(func() time.Time)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[3].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}

const (
	Version = "v0.14.4"                                         // Version of ent codegen.
	Sum     = "h1:/DhDraSLXIkBhyiVoJeSshr4ZYi7femzhj6/TckzZuI=" // Sum of ent codegen.
)
