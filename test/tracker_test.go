package test

import (
	"os"
	"testing"
	"torrent/tracker"
)

func Test_GetTrackerList(t *testing.T) {
	_ = os.Chdir("../")
	trackers, err := tracker.GetTrackerList()
	if err != nil {
		t.Error("获取tracker失败")
	} else {
		t.Log(trackers)
		t.Log("获取tracker成功")
	}
}
