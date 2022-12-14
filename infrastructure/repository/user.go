package repository

import (
	"github.com/scarlet0725/prism-api/model"
)

type User interface {
	CreateUser(*model.User) (*model.User, error)
	GetUser(id string) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	GetUserByAPIKey(string) (*model.User, error)
	DeleteUser(*model.User) error
	GetUserCalendarByUserID(int) (*model.ExternalCalendar, error)
	GetGoogleOAuthToken(int) (*model.GoogleOAuthToken, error)
	SaveExternalCalendar(*model.ExternalCalendar) (*model.ExternalCalendar, error)
	AddRegistrationEvent(*model.User, *model.Event) error
	RemoveRegistrationEvent(*model.User, *model.Event) error
	GetGoogleCalendarConfig(int) (*model.GoogleCalenderConfig, error)
	GetUserByUsername(string) (*model.User, error)
}
