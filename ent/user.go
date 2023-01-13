// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/scarlet0725/prism-api/ent/externalcalendar"
	"github.com/scarlet0725/prism-api/ent/googleoauthstate"
	"github.com/scarlet0725/prism-api/ent/googleoauthtoken"
	"github.com/scarlet0725/prism-api/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password []byte `json:"password,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// IsAdminVerified holds the value of the "is_admin_verified" field.
	IsAdminVerified bool `json:"is_admin_verified,omitempty"`
	// DeleteProtected holds the value of the "delete_protected" field.
	DeleteProtected bool `json:"delete_protected,omitempty"`
	// APIKey holds the value of the "api_key" field.
	APIKey string `json:"api_key,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// GoogleOauthTokens holds the value of the google_oauth_tokens edge.
	GoogleOauthTokens *GoogleOauthToken `json:"google_oauth_tokens,omitempty"`
	// GoogleOauthStates holds the value of the google_oauth_states edge.
	GoogleOauthStates *GoogleOauthState `json:"google_oauth_states,omitempty"`
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// ExternalCalendars holds the value of the external_calendars edge.
	ExternalCalendars *ExternalCalendar `json:"external_calendars,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// GoogleOauthTokensOrErr returns the GoogleOauthTokens value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) GoogleOauthTokensOrErr() (*GoogleOauthToken, error) {
	if e.loadedTypes[0] {
		if e.GoogleOauthTokens == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: googleoauthtoken.Label}
		}
		return e.GoogleOauthTokens, nil
	}
	return nil, &NotLoadedError{edge: "google_oauth_tokens"}
}

// GoogleOauthStatesOrErr returns the GoogleOauthStates value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) GoogleOauthStatesOrErr() (*GoogleOauthState, error) {
	if e.loadedTypes[1] {
		if e.GoogleOauthStates == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: googleoauthstate.Label}
		}
		return e.GoogleOauthStates, nil
	}
	return nil, &NotLoadedError{edge: "google_oauth_states"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[2] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// ExternalCalendarsOrErr returns the ExternalCalendars value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ExternalCalendarsOrErr() (*ExternalCalendar, error) {
	if e.loadedTypes[3] {
		if e.ExternalCalendars == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: externalcalendar.Label}
		}
		return e.ExternalCalendars, nil
	}
	return nil, &NotLoadedError{edge: "external_calendars"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldPassword:
			values[i] = new([]byte)
		case user.FieldIsAdminVerified, user.FieldDeleteProtected:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUserID, user.FieldUsername, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldAPIKey:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				u.UserID = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value != nil {
				u.Password = *value
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldIsAdminVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_admin_verified", values[i])
			} else if value.Valid {
				u.IsAdminVerified = value.Bool
			}
		case user.FieldDeleteProtected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field delete_protected", values[i])
			} else if value.Valid {
				u.DeleteProtected = value.Bool
			}
		case user.FieldAPIKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_key", values[i])
			} else if value.Valid {
				u.APIKey = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = new(time.Time)
				*u.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryGoogleOauthTokens queries the "google_oauth_tokens" edge of the User entity.
func (u *User) QueryGoogleOauthTokens() *GoogleOauthTokenQuery {
	return (&UserClient{config: u.config}).QueryGoogleOauthTokens(u)
}

// QueryGoogleOauthStates queries the "google_oauth_states" edge of the User entity.
func (u *User) QueryGoogleOauthStates() *GoogleOauthStateQuery {
	return (&UserClient{config: u.config}).QueryGoogleOauthStates(u)
}

// QueryEvents queries the "events" edge of the User entity.
func (u *User) QueryEvents() *EventQuery {
	return (&UserClient{config: u.config}).QueryEvents(u)
}

// QueryExternalCalendars queries the "external_calendars" edge of the User entity.
func (u *User) QueryExternalCalendars() *ExternalCalendarQuery {
	return (&UserClient{config: u.config}).QueryExternalCalendars(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("user_id=")
	builder.WriteString(u.UserID)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(fmt.Sprintf("%v", u.Password))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("is_admin_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.IsAdminVerified))
	builder.WriteString(", ")
	builder.WriteString("delete_protected=")
	builder.WriteString(fmt.Sprintf("%v", u.DeleteProtected))
	builder.WriteString(", ")
	builder.WriteString("api_key=")
	builder.WriteString(u.APIKey)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := u.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
