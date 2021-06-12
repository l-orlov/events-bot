package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func main() {
	loc := "msk"
	cat := []string{"concert", "cinema", "exhibition"}
	timeNow := time.Now()
	timeNowUnix := timeNow.Unix()
	timeAfterUnix := timeNow.Add(time.Hour * 24 * 7).Unix()
	params := New(loc, cat, timeNowUnix, timeAfterUnix, false)

	GetEvents(params)
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

	body := resp.Body()
	fmt.Printf("Response body is: %s", body)

	var events EventsResp

	if err := json.Unmarshal(resp.Body(), &events); err != nil {
		fmt.Printf("Unmarshal err: %s\n", err.Error())
		return nil, err
	}

	//Debug
	js, _ := json.MarshalIndent(events, "", "    ")
	fmt.Println(string(js))
	return &events, nil
}

func GetEventDescription(id int64) {

}
