package model

import "time"

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}

type CreateArtist struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CreateEvent struct {
	Name        string     `json:"name"`
	DateStr     string     `json:"date"`
	Date        *time.Time `json:"-"`
	Description string     `json:"description"`
	OpenTime    *time.Time `json:"open_time"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	UUID        string     `json:"uuid"`
	VenueName   string     `json:"venue_name"`
	Url         string     `json:"url"`
	TicketURL   string     `json:"ticket_url"`
	ArtistName  string     `json:"artist_name"`
	ArtistIDs   []string   `json:"artist_ids"`
}

type GetEvent struct {
	EventID string `json:"event_id" form:"event_id"`
}