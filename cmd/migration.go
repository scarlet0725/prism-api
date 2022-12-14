package cmd

import (
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/schema"
	"gorm.io/gorm"
)

func MigrationDB(db *gorm.DB) {

	db.AutoMigrate(
		&model.User{},
		&schema.Artist{},
		&schema.Event{},
		&schema.Venue{},
		&schema.RyzmEvent{},
		&schema.UnStructuredEventInformation{},
		&schema.Role{},
		&model.GoogleOAuthState{},
		&model.GoogleOAuthToken{},
		&model.ExternalCalendar{},
		&model.RyzmCrawlConfig{},
	)

	db.Exec("ALTER TABLE `events` CHANGE `name` `name` longtext COLLATE 'utf8mb4_ja_0900_as_cs_ks' NOT NULL AFTER `event_id`, CHANGE `description` `description` longtext COLLATE 'utf8mb4_bin' NULL AFTER `end_time`;")
	db.Exec("ALTER TABLE `artists` CHANGE `name` `name` longtext COLLATE 'utf8mb4_ja_0900_as_cs_ks' NOT NULL AFTER `artist_id`;")
	db.Exec("ALTER TABLE `venues`CHANGE `name` `name` longtext COLLATE 'utf8mb4_ja_0900_as_cs_ks' NOT NULL AFTER `venue_id`,CHANGE `description` `description` longtext COLLATE 'utf8mb4_bin' NULL AFTER `name`,CHANGE `prefecture` `prefecture` longtext COLLATE 'utf8mb4_ja_0900_as_cs_ks' NULL AFTER `postcode`,CHANGE `city` `city` longtext COLLATE 'utf8mb4_ja_0900_as_cs_ks' NULL AFTER `prefecture`,CHANGE `street` `street` longtext COLLATE 'utf8mb4_ja_0900_as_cs_ks' NULL AFTER `city`;")

	adminRole := &model.Role{
		RoleID: "Y5KIY8GI4PJ1AT7G",
		Name:   "Administrator",
	}

	memberRole := &model.Role{
		RoleID: "HTTFN3AHRSIZ8KKT",
		Name:   "Member",
	}

	db.Create(adminRole)
	db.Create(memberRole)

	user := &schema.User{
		User: model.User{
			Username:        "Admin",
			UserID:          "7CAMWFMHF4YXS23P",
			Email:           "example@example.com",
			Password:        []byte("$2a$10$h2qZT.YdMDWYRZV36LkUQO8A6sB.coL8mzkMl25VA2eOAGfJTcGZ2"),
			APIKey:          "776f9e9c12eeffad240a488d12d5c8276c947ec3d67dcce5520be08580755f8edff66e5b502f27e7c400f5b96927e478426f44ee823b484951ba789e4ed1e070",
			IsAdminVerified: true,
			DeleteProtected: true,
			Roles:           []*model.Role{adminRole},
		},
	}

	db.Create(&user)

	artists := []schema.Artist{
		{
			Artist: model.Artist{
				ArtistID: "2MERWD724422E6D8",
				Name:     "prsmin",
				URL:      "https://prsmin.com",
				RyzmCrawlConfig: &model.RyzmCrawlConfig{
					RyzmHost:       "prsmin.com",
					CrawlTargetURL: "https://api.ryzm.jp/public/lives",
					CrawlSiteType:  "ryzm",
				},
			},
		},
		{
			Artist: model.Artist{
				ArtistID: "7MHK8G565KEFQERZ",
				Name:     "On the treat Super Season",
				URL:      "https://onthetreatsuperseason.com",
				RyzmCrawlConfig: &model.RyzmCrawlConfig{
					RyzmHost:       "onthetreatsuperseason.com",
					CrawlTargetURL: "https://api.ryzm.jp/public/lives",
					CrawlSiteType:  "ryzm",
				},
			},
		},
		{
			Artist: model.Artist{
				ArtistID: "Y5KIY8GI4PJ1AT7G",
				Name:     "yosugala",
				URL:      "https://yosugala2022.ryzm.jp",
				RyzmCrawlConfig: &model.RyzmCrawlConfig{
					RyzmHost:       "yosugala2022.ryzm.jp",
					CrawlTargetURL: "https://api.ryzm.jp/public/lives",
					CrawlSiteType:  "ryzm",
				},
			},
		},
	}

	db.Create(&artists)

}
