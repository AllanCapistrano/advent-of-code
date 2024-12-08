package secondchallenge

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
		if checkValidLevelsDiffer(report[i], report[i+1]) || report[i] == report[i+1] {
			return false
		}

		if report[i] < report[i+1] {
			decreasing = true

			continue
		}

		if report[i] > report[i+1] {
			increasing = true

			continue
		}
	}

	return increasing != decreasing
}

func checkValidLevelsDiffer(currentLevel int, nextLevel int) bool {
	return math.Abs(float64(currentLevel)-float64(nextLevel)) > 3
}

func SecondChallenge() {
	result := getAndCheckReports()

	fmt.Printf("Second challenge answer: %.d\n", result)
}
