package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func AVG(nums []int) {
	var sum int
	for _, value := range nums {
		sum += value
	}
	fmt.Println("Среднее значение:", float64(sum)/float64(len(nums)))
}
func SUM(nums []int) {
	var sum int
	for _, value := range nums {
		sum += value
	}
	fmt.Println("Сумма всех элементов:", sum)
}
func MED(nums []int) {
	slices.Sort(nums)
	if len(nums)%2 != 0 {
		fmt.Println("Медиана:", float64(nums[len(nums)/2]))
	} else {
		fmt.Println("Медиана:", float64((nums[len(nums)/2]+nums[len(nums)/2-1]))/2)
	}
}
func main() {
	for {
		var str, operation string
		nums := []int{}
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Введите операцию (AVG/SUM/MED):")
		if !scanner.Scan() {
			return
		}
		operation = strings.ToUpper(strings.TrimSpace(scanner.Text()))
		fmt.Println("Введите числа через запятую:")
		if !scanner.Scan() {
			return
		}
		str = scanner.Text()
		parts := strings.FieldsFunc(str, func(r rune) bool {
			return r == ',' || r == ' '
		})
		for _, value := range parts {
			n, err := strconv.Atoi(value)
			if err == nil {
				nums = append(nums, n)
			}
		}
		if len(nums) == 0 {
			fmt.Println("Нет элементов в массиве")
			continue
		}
		switch {
		case operation == "AVG":
			AVG(nums)
		case operation == "SUM":
			SUM(nums)
		case operation == "MED":
			MED(nums)
		}
	}
}
