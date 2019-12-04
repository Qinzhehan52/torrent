package test

import (
	"testing"
	"torrent/tracker"
)

func Test_GetTrackerList(t *testing.T) {
	_, err := tracker.GetTrackerList()
	if err != nil {
		t.Error("获取tracker失败")
	} else {
		t.Log("获取tracker成功")
	}
}
