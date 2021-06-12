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
