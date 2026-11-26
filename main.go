package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func AVG(nums []int) {
	var sum int
	for _, value := range nums {
		sum += value
	}
	fmt.Println("Среднее значение:", sum/len(nums))
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
		fmt.Println("Медиана:", nums[len(nums)/2])
	} else {
		fmt.Println("Медиана:", (nums[len(nums)/2]+nums[len(nums)/2-1])/2)
	}
}
func main() {
	for {
		var str, operation string
		nums := []int{}
		fmt.Println("Введите операцию(AVG/SUM/MED):")
		fmt.Scan(&operation)
		fmt.Println("Введите числа через запятую:")
		fmt.Scan(&str)
		parts := strings.Split(str, ",")
		for _, value := range parts {
			n, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Ошибка парсинга")
				continue
			}
			nums = append(nums, n)
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
