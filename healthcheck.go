package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Request struct {
	start      time.Time
	end        time.Time
	totalTime  time.Duration
	respBody   string
	statusCode int
}

type Requests []Request

func (r Requests) GetAvgTotalTime() time.Duration {
	var totalDuration time.Duration
	for _, request := range r {
		totalDuration += request.totalTime
	}
	return time.Duration(int64(totalDuration) / int64(len(r)))
}

const requestLink1 = "http://localhost:3002/api/hrpmp/v2/ping?_1722545325981_______________"
const requestLink2 = "http://localhost:3002/api/hrpmp/v2/ping?125445335332588______________"
const requestLink3 = "http://localhost:3002/api/hrpmp/v2/ping?1353322222221388_____________"
const requestLink4 = "http://localhost:3002/api/hrpmp/v2/ping?23322222222211246_____________"
const requestLink5 = "http://localhost:3002/api/hrpmp/v2/ping?233222222222221111222223499____"
const requestLink6 = "http://localhost:3002/api/hrpmp/v2/ping?12222222222222233555555532508___"
const requestLink7 = "http://localhost:3002/api/hrpmp/v2/ping?122222222222222222333332332188__"
const requestLink8 = "http://localhost:3002/api/hrpmp/v2/ping?3122222222222222222222222217288_"
const requestLink9 = "http://localhost:3002/api/hrpmp/v2/ping?911222222222222222222222221__485"
const requestLink10 = "http://localhost:3002/api/hrpmp/v2/ping?831122222222222222222222227__388"
const requestLink11 = "http://localhost:3002/api/hrpmp/v2/ping?58212222222222222222222211___088"
const requestLink12 = "http://localhost:3002/api/hrpmp/v2/ping?_80172222222222222222227____888_"
const requestLink13 = "http://localhost:3002/api/hrpmp/v2/ping?__867222222222222221______0888__"
const requestLink14 = "http://localhost:3002/api/hrpmp/v2/ping?__18512222222211_______488886___"
const requestLink15 = "http://localhost:3002/api/hrpmp/v2/ping?___887777__________68888887___"
const requestLink16 = "http://localhost:3002/api/hrpmp/v2/ping?____88________508888888______"
const requestLink17 = "http://localhost:3002/api/hrpmp/v2/ping?_____85488888888885_________"

var sentRequests Requests
var not200Codes int
var resChan chan Request
var wg sync.WaitGroup

func stackRequests() {
	for r := range resChan {
		wg.Done()
		sentRequests = append(sentRequests, r)
	}
}

func doRequest() {
	start := time.Now()
	// resp, err := http.Get(requestLink1)
	// resp, err = http.Get(requestLink2)
	// resp, err = http.Get(requestLink3)
	// resp, err = http.Get(requestLink4)
	// resp, err = http.Get(requestLink5)
	// resp, err = http.Get(requestLink6)
	// resp, err = http.Get(requestLink7)
	// resp, err = http.Get(requestLink8)
	// resp, err = http.Get(requestLink9)
	// resp, err = http.Get(requestLink10)
	// resp, err = http.Get(requestLink11)
	// resp, err = http.Get(requestLink12)
	// resp, err = http.Get(requestLink13)
	// resp, err = http.Get(requestLink14)
	// resp, err = http.Get(requestLink15)
	// resp, err = http.Get(requestLink16)
	// resp, err = http.Get(requestLink17)
	resp, err := http.Get("https://acting-dev-globalunity.safeguardglobal.com/api/hrpmp/v2/ping")
	totalTime := time.Since(start)
	end := time.Now()
	if err != nil {
		fmt.Print("Error requesting:", err)
	}
	defer resp.Body.Close()
	var body string
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	body = string(bodyBytes)
	if resp.StatusCode != 200 {
		not200Codes++
		fmt.Println("Request with error", body)
	}

	r := Request{
		start:      start,
		end:        end,
		totalTime:  totalTime,
		statusCode: resp.StatusCode,
		respBody:   body,
	}
	resChan <- r
}

func ExecuteBulkREquests(reqNumber int, delay int64) {
	resChan = make(chan Request)
	fmt.Printf("\n\nStarting\n\n")

	go stackRequests()

	for i := 0; i < reqNumber; i++ {
		wg.Add(1)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		go doRequest()
	}
	wg.Wait()
	fmt.Printf("Total Requests: %d\n\n", len(sentRequests))
	fmt.Printf("Not 200 Codes: %d\n", not200Codes)
	fmt.Printf("avg time: %d ms\n", sentRequests.GetAvgTotalTime()/time.Millisecond)
}
