package main

type Params struct {
	Location   string
	Categories []string
	StartDate  int64
	EndDate    int64
	IsFree     bool
}

func New(
	location string,
	categories []string,
	startDate int64,
	endDate int64,
	isFree bool,
) Params {
	return Params{
		Location:   location,
		Categories: categories,
		StartDate:  startDate,
		EndDate:    endDate,
		IsFree:     isFree,
	}
}

type EventsResp struct {
	Count    *int         `json:"count"`
	Next     *string      `json:"next"`
	Previous *string      `json:"previous"`
	Results  []EventShort `json:"results"`
}

type EventShort struct {
	Id    *int    `json:"id"`
	Title *string `json:"title"`
	Slug  *string `json:"slug"`
}

type Event struct {
	ID              int          `json:"id"`
	PublicationDate int          `json:"publication_date"`
	//Dates           Dates        `json:"dates"`
	//Title           string       `json:"title"`
	//Slug            string       `json:"slug"`
	//Place           Place        `json:"place"`
	//Description     string       `json:"description"`
	//BodyText        string       `json:"body_text"`
	//Location        Location     `json:"location"`
	//Categories      []string     `json:"categories"`
	//Tagline         string       `json:"tagline"`
	//AgeRestriction  string       `json:"age_restriction"`
	//Price           string       `json:"price"`
	//IsFree          bool         `json:"is_free"`
	//Images          Images       `json:"images"`
	//FavoritesCount  int          `json:"favorites_count"`
	//CommentsCount   int          `json:"comments_count"`
	//SiteURL         string       `json:"site_url"`
	//ShortTitle      string       `json:"short_title"`
	//Tags            []string     `json:"tags"`
	//DisableComments bool         `json:"disable_comments"`
	//Participants    Participants `json:"participants"`
}

type Dates []Date

type Date struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Place struct {
	ID int `json:"id"`
}

type Location struct {
	Slug string `json:"slug"`
}

type Images []Image

type Image struct {
	Image  string `json:"image"`
	Source struct {
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"source"`
}

type Participants []Participant

type Participant struct {
	Role  Role  `json:"role"`
	Agent Agent `json:"agent"`
}

type Role struct {
	Slug string `json:"slug"`
}

type Agent struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Slug      string        `json:"slug"`
	AgentType string        `json:"agent_type"`
	Images    []interface{} `json:"images"` //думаю это все не надо вообще
	SiteURL   string        `json:"site_url"`
}
