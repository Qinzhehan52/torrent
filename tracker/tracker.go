package tracker

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetTrackerList() ([][]string, error) {
	var body []byte
	var err error
	//read online tracker list
	trackerListUrl := "https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_all.txt"
	resp, err := http.Get(trackerListUrl)
	trackerList := make([][]string, 0)

	//if failed, read local trackers
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("在线获取tracker失败，读取本地缓存")
		body, err = GetTrackListLocal()
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}

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

func GetTrackListLocal() ([]byte, error) {
	file, err := os.Open("trackers_cache")
	if err != nil {
		return nil, errors.New("Failed to open local tracker: " + err.Error())
	}
	defer file.Close()

	body, err := ioutil.ReadAll(file)

	return body, err
}
