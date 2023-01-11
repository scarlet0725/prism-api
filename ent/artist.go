// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/scarlet0725/prism-api/ent/artist"
)

// Artist is the model entity for the Artist schema.
type Artist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ArtistID holds the value of the "artist_id" field.
	ArtistID string `json:"artist_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtistQuery when eager-loading is set.
	Edges ArtistEdges `json:"edges"`
}

// ArtistEdges holds the relations/edges for other nodes in the graph.
type ArtistEdges struct {
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e ArtistEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[0] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Artist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case artist.FieldID:
			values[i] = new(sql.NullInt64)
		case artist.FieldArtistID, artist.FieldName, artist.FieldURL:
			values[i] = new(sql.NullString)
		case artist.FieldCreatedAt, artist.FieldUpdatedAt, artist.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Artist", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Artist fields.
func (a *Artist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case artist.FieldArtistID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field artist_id", values[i])
			} else if value.Valid {
				a.ArtistID = value.String
			}
		case artist.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case artist.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case artist.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case artist.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case artist.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryEvents queries the "events" edge of the Artist entity.
func (a *Artist) QueryEvents() *EventQuery {
	return (&ArtistClient{config: a.config}).QueryEvents(a)
}

// Update returns a builder for updating this Artist.
// Note that you need to call Artist.Unwrap() before calling this method if this Artist
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Artist) Update() *ArtistUpdateOne {
	return (&ArtistClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Artist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Artist) Unwrap() *Artist {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Artist is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Artist) String() string {
	var builder strings.Builder
	builder.WriteString("Artist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("artist_id=")
	builder.WriteString(a.ArtistID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(a.URL)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := a.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Artists is a parsable slice of Artist.
type Artists []*Artist

func (a Artists) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}