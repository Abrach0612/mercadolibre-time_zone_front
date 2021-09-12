package util

import (
	"strconv"
	"strings"
	"time"
)

func GetTimeGMT(timeNew time.Time, gmt string) time.Time {
	t := timeNew.Add(time.Duration(-GetAtoI(gmt)) * time.Hour)
	return t
}

func GetAtoI(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

func GetSplitDatenadHourETS(jsonDate string) (hourSplit []string, dateSplit []string, gmt string) {
	hourSplit = strings.Split(jsonDate[11:16], ":")
	dateSplit = strings.Split(jsonDate[:10], "-")
	if len(jsonDate) == 21 {
		gmt = jsonDate[len(jsonDate)-2:]

	} else if len(jsonDate) == 22 {
		gmt = jsonDate[len(jsonDate)-3:]

	}
	return hourSplit, dateSplit, gmt
}
