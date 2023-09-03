// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/IshinShibata/chatApp/ent/chat"
	"github.com/IshinShibata/chatApp/ent/schema"
	"github.com/IshinShibata/chatApp/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatFields := schema.Chat{}.Fields()
	_ = chatFields
	// chatDescTitle is the schema descriptor for title field.
	chatDescTitle := chatFields[0].Descriptor()
	// chat.DefaultTitle holds the default value on creation for the title field.
	chat.DefaultTitle = chatDescTitle.Default.(string)
	// chatDescImage is the schema descriptor for image field.
	chatDescImage := chatFields[2].Descriptor()
	// chat.DefaultImage holds the default value on creation for the image field.
	chat.DefaultImage = chatDescImage.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
