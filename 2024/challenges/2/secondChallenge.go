package secondchallenges

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func getAndCheckReports() int {
	var reportString []string
	var report []int
	var amountSafeReports int

	file, err := os.Open("challenges/2/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		reportString = strings.Split(line, " ")

		for i := 0; i < len(reportString); i++ {
			level, _ := strconv.Atoi(reportString[i])
			report = append(report, level)
		}

		if checkReportIsSafe(report) {
			amountSafeReports++
		}

		report = nil

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
			os.Exit(-1)
		}
	}

	return amountSafeReports
}

func checkReportIsSafe(report []int) bool {
	increasing := false
	decreasing := false

	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			decreasing = true

			if math.Abs(float64(report[i])-float64(report[i+1])) > 3 {
				return false
			}

			continue
		}

		if report[i] > report[i+1] {
			increasing = true

			if math.Abs(float64(report[i])-float64(report[i+1])) > 3 {
				return false
			}

			continue
		}

		if report[i] == report[i+1] {
			return false
		}
	}

	return increasing != decreasing
}

func SecondChallenge() int {
	return getAndCheckReports()
}
