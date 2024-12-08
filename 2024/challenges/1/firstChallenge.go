package challenges

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func removeElementByIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func FirstChallenge() float64 {
	var leftIndex int
	var rightIndex int
	var distance float64
	var result float64

	left, right := getLeftRight()
	listSize := len(left)

	for i := 0; i < listSize; i++ {
		distance = findSmallestNumber(left, right, &leftIndex, &rightIndex)

		left = removeElementByIndex(left, leftIndex)
		right = removeElementByIndex(right, rightIndex)

		result += distance
	}

	return result
}

func getLeftRight() ([]int, []int) {
	var numbers []string
	var left []int
	var right []int

	file, err := os.Open("challenges/1/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		numbers = strings.Split(line, "   ")

		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[1])

		left = append(left, leftNumber)
		right = append(right, rightNumber)

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
			os.Exit(-1)
		}
	}

	return left, right
}

func findSmallestNumber(left []int, right []int, leftIndex *int, rightIndex *int) float64 {
	leftMinValue := left[0]
	rightMinValue := right[0]

	for j := 0; j < len(left); j++ {
		if left[j] <= leftMinValue {
			leftMinValue = left[j]
			*leftIndex = j
		}

		if right[j] <= rightMinValue {
			rightMinValue = right[j]
			*rightIndex = j
		}
	}

	return math.Abs(float64(leftMinValue) - float64(rightMinValue))
}
