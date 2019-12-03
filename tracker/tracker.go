package tracker

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetTrackerList() ([][]string, error) {
	trackerListUrl := "https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_all.txt"
	resp, err := http.Get(trackerListUrl)
	trackerList := make([][]string, 0)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rawTrackerList := strings.Split(string(body), "\n")

	for _, url := range rawTrackerList {
		if len(url) > 1 {
			tracker := []string{url}
			trackerList = append(trackerList, tracker)
		}
	}

	return trackerList, nil
}
