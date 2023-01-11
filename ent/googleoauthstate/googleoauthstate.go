// Code generated by ent, DO NOT EDIT.

package googleoauthstate

const (
	// Label holds the string label denoting the googleoauthstate type in the database.
	Label = "google_oauth_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the googleoauthstate in the database.
	Table = "google_oauth_states"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "google_oauth_states"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_google_oauth_states"
)

// Columns holds all SQL columns for googleoauthstate fields.
var Columns = []string{
	FieldID,
	FieldState,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "google_oauth_states"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_google_oauth_states",
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
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(string) error
)
