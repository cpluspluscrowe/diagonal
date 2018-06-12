package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func strToNumbers(rowStr []string) []int {
	var numbers []int
	for _, item := range rowStr {
		var number, _ = strconv.Atoi(item)
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	var reader = bufio.NewReaderSize(os.Stdin, 1024*1024)
	var firstInput, _, _ = reader.ReadLine()
	var rowCount, _ = strconv.Atoi(string(firstInput))
	var primary int = 0
	var secondary int = 0
	for i := 0; i < rowCount; i++ {
		var input, _, _ = reader.ReadLine()
		var rowStr = strings.Split(string(input), " ")
		var numbers = strToNumbers(rowStr)
		primary += numbers[i]
		secondary += numbers[rowCount-i-1]
	}
	fmt.Println(math.Abs(float64(primary - secondary)))
}
