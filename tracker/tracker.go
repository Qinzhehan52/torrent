package tracker

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetTrackerList() ([]string, error) {
	trackerListUrl := "https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_all.txt"
	resp, err := http.Get(trackerListUrl)
	trackerList := make([]string, 100)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rawTrackerList := strings.Split(string(body), "\n")
	log.Println(string(body))

	for _, url := range rawTrackerList {
		if len(url) > 1 {
			trackerList = append(trackerList, url)
		}
	}

	return trackerList, nil
}
