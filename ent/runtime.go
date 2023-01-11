// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/scarlet0725/prism-api/ent/artist"
	"github.com/scarlet0725/prism-api/ent/event"
	"github.com/scarlet0725/prism-api/ent/externalcalendar"
	"github.com/scarlet0725/prism-api/ent/googleoauthstate"
	"github.com/scarlet0725/prism-api/ent/schema"
	"github.com/scarlet0725/prism-api/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	artistFields := schema.Artist{}.Fields()
	_ = artistFields
	// artistDescArtistID is the schema descriptor for artist_id field.
	artistDescArtistID := artistFields[0].Descriptor()
	// artist.ArtistIDValidator is a validator for the "artist_id" field. It is called by the builders before save.
	artist.ArtistIDValidator = artistDescArtistID.Validators[0].(func(string) error)
	// artistDescName is the schema descriptor for name field.
	artistDescName := artistFields[1].Descriptor()
	// artist.NameValidator is a validator for the "name" field. It is called by the builders before save.
	artist.NameValidator = artistDescName.Validators[0].(func(string) error)
	// artistDescCreatedAt is the schema descriptor for created_at field.
	artistDescCreatedAt := artistFields[3].Descriptor()
	// artist.DefaultCreatedAt holds the default value on creation for the created_at field.
	artist.DefaultCreatedAt = artistDescCreatedAt.Default.(func() time.Time)
	// artistDescUpdatedAt is the schema descriptor for updated_at field.
	artistDescUpdatedAt := artistFields[4].Descriptor()
	// artist.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	artist.DefaultUpdatedAt = artistDescUpdatedAt.Default.(func() time.Time)
	// artist.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	artist.UpdateDefaultUpdatedAt = artistDescUpdatedAt.UpdateDefault.(func() time.Time)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescEventID is the schema descriptor for event_id field.
	eventDescEventID := eventFields[0].Descriptor()
	// event.EventIDValidator is a validator for the "event_id" field. It is called by the builders before save.
	event.EventIDValidator = eventDescEventID.Validators[0].(func(string) error)
	// eventDescName is the schema descriptor for name field.
	eventDescName := eventFields[1].Descriptor()
	// event.NameValidator is a validator for the "name" field. It is called by the builders before save.
	event.NameValidator = eventDescName.Validators[0].(func(string) error)
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventFields[9].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescUpdatedAt is the schema descriptor for updated_at field.
	eventDescUpdatedAt := eventFields[10].Descriptor()
	// event.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	event.DefaultUpdatedAt = eventDescUpdatedAt.Default.(func() time.Time)
	// event.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	event.UpdateDefaultUpdatedAt = eventDescUpdatedAt.UpdateDefault.(func() time.Time)
	externalcalendarFields := schema.ExternalCalendar{}.Fields()
	_ = externalcalendarFields
	// externalcalendarDescName is the schema descriptor for name field.
	externalcalendarDescName := externalcalendarFields[0].Descriptor()
	// externalcalendar.NameValidator is a validator for the "name" field. It is called by the builders before save.
	externalcalendar.NameValidator = externalcalendarDescName.Validators[0].(func(string) error)
	// externalcalendarDescCalendarID is the schema descriptor for calendar_id field.
	externalcalendarDescCalendarID := externalcalendarFields[2].Descriptor()
	// externalcalendar.CalendarIDValidator is a validator for the "calendar_id" field. It is called by the builders before save.
	externalcalendar.CalendarIDValidator = externalcalendarDescCalendarID.Validators[0].(func(string) error)
	// externalcalendarDescType is the schema descriptor for type field.
	externalcalendarDescType := externalcalendarFields[3].Descriptor()
	// externalcalendar.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	externalcalendar.TypeValidator = externalcalendarDescType.Validators[0].(func(string) error)
	// externalcalendarDescCreatedAt is the schema descriptor for created_at field.
	externalcalendarDescCreatedAt := externalcalendarFields[4].Descriptor()
	// externalcalendar.DefaultCreatedAt holds the default value on creation for the created_at field.
	externalcalendar.DefaultCreatedAt = externalcalendarDescCreatedAt.Default.(func() time.Time)
	// externalcalendarDescUpdatedAt is the schema descriptor for updated_at field.
	externalcalendarDescUpdatedAt := externalcalendarFields[5].Descriptor()
	// externalcalendar.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	externalcalendar.DefaultUpdatedAt = externalcalendarDescUpdatedAt.Default.(func() time.Time)
	// externalcalendar.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	externalcalendar.UpdateDefaultUpdatedAt = externalcalendarDescUpdatedAt.UpdateDefault.(func() time.Time)
	googleoauthstateFields := schema.GoogleOauthState{}.Fields()
	_ = googleoauthstateFields
	// googleoauthstateDescState is the schema descriptor for state field.
	googleoauthstateDescState := googleoauthstateFields[0].Descriptor()
	// googleoauthstate.StateValidator is a validator for the "state" field. It is called by the builders before save.
	googleoauthstate.StateValidator = googleoauthstateDescState.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserID is the schema descriptor for user_id field.
	userDescUserID := userFields[0].Descriptor()
	// user.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	user.UserIDValidator = userDescUserID.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func([]byte) error)
	// userDescIsAdminVerified is the schema descriptor for is_admin_verified field.
	userDescIsAdminVerified := userFields[6].Descriptor()
	// user.DefaultIsAdminVerified holds the default value on creation for the is_admin_verified field.
	user.DefaultIsAdminVerified = userDescIsAdminVerified.Default.(bool)
	// userDescDeleteProtected is the schema descriptor for delete_protected field.
	userDescDeleteProtected := userFields[7].Descriptor()
	// user.DefaultDeleteProtected holds the default value on creation for the delete_protected field.
	user.DefaultDeleteProtected = userDescDeleteProtected.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[9].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[10].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
