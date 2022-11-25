package usecase

import (
	"fmt"
	"time"

	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/repository"
	"github.com/scarlet0725/prism-api/selializer"
)

type Event interface {
	CreateEvent(*model.CreateEvent) (*model.Event, error)
	DeleteEvent(*model.Event) error
	//GetEvent(string) (*model.Event, error)
	GetEventsByArtistName(string) ([]*model.Event, *model.AppError)
	CreateArtistEventsFromCrawlData(id string) ([]*model.Event, error)
	GetEventByID(string) (*model.Event, error)
}

type eventUsecase struct {
	db         repository.DB
	fetch      controller.FetchController
	parser     parser.DocParser
	selializer selializer.ResponseSerializer
	json       parser.JsonParser
}

func NewEventApplication(db repository.DB, fetch controller.FetchController, parser parser.DocParser, selializer selializer.ResponseSerializer, json parser.JsonParser) Event {
	return &eventUsecase{
		db:         db,
		fetch:      fetch,
		parser:     parser,
		selializer: selializer,
		json:       json,
	}
}

func (a *eventUsecase) CreateEvent(e *model.CreateEvent) (*model.Event, error) {
	id := cmd.MakeRamdomID(eventIDLength)

	artists, _ := a.db.GetArtistsByIDs(e.ArtistIDs)
	venue, _ := a.db.GetVenueByID(e.VenueID)
	jst, _ := time.LoadLocation(Locale)

	date, err := time.ParseInLocation("2006-01-02", e.Date, jst)

	if err != nil {
		return nil, err
	}

	fmt.Println(artists)

	event := &model.Event{
		EventID:     id,
		Name:        e.Name,
		Date:        &date,
		Description: e.Description,
		OpenTime:    e.OpenTime,
		StartTime:   e.StartTime,
		EndTime:     e.EndTime,
		Url:         e.Url,
		TicketURL:   e.TicketURL,
		Artists:     artists,
		Venue:       venue,
	}
	fmt.Println(event)

	return a.db.CreateEvent(event)
}

func (a *eventUsecase) GetEventsByName(name string) ([]model.Event, error) {
	return nil, nil
}

func (a *eventUsecase) GetEventsByArtistName(name string) ([]*model.Event, *model.AppError) {
	artist, err := a.db.GetArtistByName(name)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "artist_not_found",
		}
	}

	req := &model.ScrapingRequest{
		Host: artist.CrawlTargetURL,
		URL:  artist.CrawlTargetURL,
		Type: artist.CrawlSiteType,
		Option: model.FetchOptions{
			IsUseCache: true,
		},
	}

	if req.Type == "ryzm" {
		req.Option = model.FetchOptions{
			HTTPHeader: map[string]string{
				"X-RYZM-HOST": artist.RyzmHost,
			},
			HTTPParams: map[string]string{
				"archived": "0",
			},
		}
	}

	res, err := a.fetch.Fetch(req)

	if err != nil {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "fetch_error",
		}
	}

	switch artist.CrawlSiteType {
	case "ryzm":
		json, err := a.json.Ryzm(res.Data)
		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}
		result, err := a.selializer.SelializeRyzmData(json)

		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}

		return result, nil
	default:
		return nil, &model.AppError{
			Code: 500,
			Msg:  "error!",
		}

	}
}

func (a *eventUsecase) CreateArtistEventsFromCrawlData(id string) ([]*model.Event, error) {
	artist, err := a.db.GetArtistByID(id)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "artist_not_found",
		}
	}

	if artist.CrawlTargetURL == "" || artist.CrawlSiteType == "" || artist.RyzmHost == "" {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "This artist is not supported auto update",
		}
	}

	req := &model.ScrapingRequest{
		Host: artist.CrawlTargetURL,
		URL:  artist.CrawlTargetURL,
		Type: artist.CrawlSiteType,
		Option: model.FetchOptions{
			IsUseCache: true,
		},
	}

	if req.Type == "ryzm" {
		req.Option = model.FetchOptions{
			HTTPHeader: map[string]string{
				"X-RYZM-HOST": artist.RyzmHost,
			},
			HTTPParams: map[string]string{
				"archived": "0",
			},
		}
	}

	res, err := a.fetch.Fetch(req)

	if err != nil {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "fetch_error",
		}
	}

	switch artist.CrawlSiteType {
	case "ryzm":

		json, err := a.json.Ryzm(res.Data)
		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}
		result, err := a.selializer.SelializeRyzmData(json)

		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}

		uuids := []string{}
		eventAlreadyRegistered := map[string]*model.Event{}

		for _, event := range result {
			uuids = append(uuids, event.UUID)
		}

		registeredEvents, err := a.db.GetEventsByUUIDs(uuids)

		for _, event := range registeredEvents {
			eventAlreadyRegistered[event.UUID] = event
		}

		//TODO: 既に登録されているイベントは更新する
		//FIXME: データベースから更新できない場合の処理を考える
		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_get_events",
			}
		}

		registrationExpectedEvents := []*model.Event{}

		for _, event := range result {
			_, ok := eventAlreadyRegistered[event.UUID]
			if ok {
				continue
			}
			event.EventID = cmd.MakeRamdomID(eventIDLength)
			event.Artists = append(event.Artists, artist)
			registrationExpectedEvents = append(registrationExpectedEvents, event)
		}

		if len(registrationExpectedEvents) <= 0 {
			return []*model.Event{}, nil
		}

		return a.db.CreateEvents(registrationExpectedEvents)

	default:
		return nil, &model.AppError{
			Code: 500,
			Msg:  "error!",
		}

	}

}

func (a *eventUsecase) GetEventByID(ID string) (*model.Event, error) {
	return a.db.GetEventByID(ID)
}

func (a *eventUsecase) DeleteEvent(event *model.Event) error {
	return a.db.DeleteEvent(event)
}