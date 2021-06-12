package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

type MutexEvents struct {
	sync.Mutex
	Events []Event
}



func main() {
	loc := "kzn"
	cat := []string{"concert"}
	timeNow := time.Now()
	timeNowUnix := timeNow.Unix()
	timeAfterUnix := timeNow.Add(time.Hour * 24 * 20).Unix()
	params := New(loc, cat, timeNowUnix, timeAfterUnix, false)

	mtx := MutexEvents{}
	events, err := GetEvents(params)
	if err == nil && len(events.Results) != 0 {
		fmt.Println("найдено эвентов:", len(events.Results))
		wg := sync.WaitGroup{}

		for _, event := range events.Results {
			if event.Id == nil {
				continue
			}

			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				event, err := GetEventDescription(id)
				if err == nil && event != nil {
					mtx.Lock()
					defer mtx.Unlock()

					mtx.Events = append(mtx.Events, *event)
				}
			}(*event.Id)
		}
		wg.Wait()
	}

	fmt.Println("events finded:", len(mtx.Events))
}

const (
	kudaGoUrl       = "https://kudago.com/public-api/v1.4/events/"
	locationParam   = "&location="
	languageParam   = "?lang="
	dataStartParam  = "&actual_since="
	dataEndParam    = "&actual_until="
	categoriesParam = "&categories="
	isFreeParam     = "&is_free="
)

func GetEvents(params Params) (*EventsResp, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	url := buildUrl(params)

	//debug
	fmt.Println("url:", url)

	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		return nil, err
	}

	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		fmt.Printf("Expected content type application/json but got %s\n", contentType)
		return nil, err
	}

	//body := resp.Body()
	//fmt.Printf("Response body is: %s", body)

	var events EventsResp

	if err := json.Unmarshal(resp.Body(), &events); err != nil {
		fmt.Printf("Unmarshal err: %s\n", err.Error())
		return nil, err
	}

	//Debug
	//js, _ := json.MarshalIndent(events, "", "    ")
	//fmt.Println(string(js))
	return &events, nil
}

func GetEventDescription(id int) (*Event, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	url := fmt.Sprintf("%s%d/", kudaGoUrl, id)

	//debug
	fmt.Println("url:", url)

	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		return nil, err
	}

	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		fmt.Printf("Expected content type application/json but got %s\n", contentType)
		return nil, err
	}

	var event Event

	if err := json.Unmarshal(resp.Body(), &event); err != nil {
		fmt.Printf("Unmarshal err: %s\n", err.Error())
		return nil, err
	}

	//Debug
	//js, _ := json.MarshalIndent(event, "", "    ")
	//fmt.Println(string(js))
	return &event, nil
}
