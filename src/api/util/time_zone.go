package util

import (
	"bufio"
	"fmt"
	"os"
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
	dateSplit = strings.Split(jsonDate[:10], "-")
	hourSplit = strings.Split(jsonDate[11:16], ":")
	gmt = jsonDate[19:22]
	return hourSplit, dateSplit, gmt
}

func PrintPressEnter() {
	fmt.Println("\nPress Enter...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func CreateCPTString(data string, gmt string) string {
	return data[:10] + "  " + data[11:16] + "   " + "[(" + gmt + "])"
}
