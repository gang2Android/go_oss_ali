package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func ListAll(path string, curHier int) string {
	readerInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	readerInfos1 := sortByTime(readerInfos)
	for _, info := range readerInfos1 {
		if info.IsDir() {
			ListAll(path+"/"+info.Name(), curHier+1)
		} else {
			return path + "/" + info.Name()
		}
	}
	return ""
}

func sortByTime(pl []os.FileInfo) []os.FileInfo {
	sort.Slice(pl, func(i, j int) bool {
		flag := false
		if pl[i].ModTime().After(pl[j].ModTime()) {
			flag = true
		} else if pl[i].ModTime().Equal(pl[j].ModTime()) {
			if pl[i].Name() < pl[j].Name() {
				flag = true
			}
		}
		return flag
	})
	return pl
}
