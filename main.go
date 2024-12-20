package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// findNo([]int{1, 23, 4, 5}, 4)
	// sumOfNumbers([]int{1, 2, 3, 4, 5})
	// reverseString("emoh")

	// FizzBuzz(20)
	// DecimalToBase(14, 2)
	// BubbleSort([]int{1, 3, 2, 5, 9, 6})
	// LinearSearch(5, []int{1, 3, 2, 5, 9, 6})
	// BinarySearch(18, []int{1, 3, 5, 6, 8, 12, 18, 19})
	// HasDuplicateValue([]int{1, 3, 5, 6, 8, 12, 18, 19})
	// GreatestNum([]int{1, 3, 5, 6, 8, 12, 18, 19})
	// SelectionSort([]int{1, 3, 2, 5, 9, 6})
	// InsertionSort([]int{1, 3, 2, 5, 9, 6})
	// AverageOfEvenNumbers([]int{1, 3, 2, 5, 9, 6})
	// WordBuilder([]string{"a", "b", "c", "d"})
	// MarkInventory([]string{"Purple shirt", "Black shirt"})
	// CountTheOnes([][]int{{0, 1, 1, 1, 0}, {0, 1, 0, 1, 0, 1}, {1, 0}})
	// IsPalindrome("racecar")
	// TwoNumberProducts([]int{1, 2, 3, 4, 5})
	// CircularSentences("hey yoly ylh")
	// PartitionEqualSubsetSum([]int{1, 5, 11, 5})
	// NoOfChars([]string{"a", "a", "b", "b", "b", "c", "c", "d"})
	// NoOfCharStrings("AABBBCCDDDD")
	// EncodeStringChars("AABBBCCDDDD")
	// DecodeStringChars("2A3B2C4D")
	// TwoSumProblemBrute([]int{1, 2, 3, 4, 5, 6}, 9)
	TwoSumProblemOp([]int{1, 2, 3, 4, 5, 6}, 9)
}

func log(data any) {
	fmt.Println("-> ", data)
}

func FindNo(listOfNums []int, num int) bool {
	for i := range listOfNums {
		if listOfNums[i] == num {
			println("Number exists here bro")
			return true
		}
	}
	println("Number does not exist bro")
	return false
}

func SumOfNumbers(nos []int) int {
	total := 0
	for _, i := range nos {
		total = total + i
	}
	println(total)
	return total
}

func ReverseString(str string) string {
	var word string
	for _, i := range str {
		word = string(i) + word
	}
	return word
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}

func DecimalToBase(num, base int) string {
	const charset = "0123456789ABCDEF"
	var sb strings.Builder
	for num > 0 {
		rem := num % base
		sb.WriteByte(charset[rem])
		num = num / base
	}
	fmt.Println("after sb", sb.String())
	return sb.String()
}

func BubbleSort(list []int) {
	n := len(list)
	swapped := true

	for swapped {
		swapped = false
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				// Swap elements
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		n-- // Optimize by reducing the array size after each iteration
	}
	fmt.Println("Bubble sort:", list)
}

func LinearSearch(num int, list []int) {
	for _, i := range list {
		if i == num {
			fmt.Println(i)
		} else if i > num {
			fmt.Println(i)
			return
		}
	}
}

func BinarySearch(num int, list []int) (int, bool) {
	lowerBound := 0
	higherBound := len(list)

	for lowerBound <= higherBound {
		midPoint := (higherBound + lowerBound) / 2

		valueAtMidPoint := list[midPoint]

		if num == valueAtMidPoint {
			return num, true
		} else if num < valueAtMidPoint {
			higherBound = midPoint - 1
		} else if num > valueAtMidPoint {
			lowerBound = midPoint + 1
		}
	}
	return 0, false
}

func HasDuplicateValue(list []int) bool {
	for _, i := range list {
		for _, j := range list {
			if i != j && list[i] == list[j] {
				fmt.Print("has duplicate value", true)
				return true
			}
		}
	}
	return false
}

func GreatestNum(list []int) int {
	greatestNumSoFar := list[0]
	for _, i := range list {
		if i > greatestNumSoFar {
			greatestNumSoFar = i
		}
	}
	fmt.Print("greatest num so far", greatestNumSoFar)
	return greatestNumSoFar
}

func SelectionSort(list []int) {
	l := len(list)
	for i := 0; i < l-1; i++ {
		lowest := i
		for j := i + 1; j < l; j++ {
			if list[j] < list[lowest] {
				lowest = j
			}
		}
		if lowest != i {
			list[i], list[lowest] = list[lowest], list[i]
		}
	}
	fmt.Print("Selection sort ", list)
}

func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		temp := list[i]
		pos := i - 1
		for pos >= 0 && list[pos] > temp {
			list[pos+1] = list[pos]
			pos = pos - 1
		}
		list[pos+1] = temp
	}
	fmt.Println("InsertionSort ", list)
	return list
}

func AverageOfEvenNumbers(list []int) int {
	sum := 0
	countOfAvgNum := 0

	for _, i := range list {
		if i%2 == 0 {
			sum += i
			countOfAvgNum += 1
		}
	}

	if countOfAvgNum == 0 {
		return 0
	}
	fmt.Print("AverageOfEvenNumbers ", sum)
	return sum
}

func WordBuilder(list []string) []string {
	l := []string{}
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list); j++ {
			l = append(l, list[i]+list[j])
		}
	}
	fmt.Print("WordBuilder ", l)
	return l
}

func AverageCelcius(farenheitReadings []int) float64 {
	if len(farenheitReadings) == 0 {
		return 0
	}
	sum := 0.0
	for _, fahrenheit := range farenheitReadings {
		celsiusConversion := float64(fahrenheit-32) * 5 / 9
		sum += celsiusConversion
	}
	return sum / float64(len(farenheitReadings))
}

func MarkInventory(list []string) {
	clothings := []string{}

	for i := 0; i < len(list); i++ {
		num := []int{1, 2, 3, 4, 5}
		for _, r := range num {
			clothings = append(clothings, list[i]+" Sum "+strconv.Itoa(r))
		}
	}

	fmt.Print("MarkInventory, ", clothings)
}

func CountTheOnes(list [][]int) int {
	count := 0

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			if list[i][j] == 1 {
				count += 1
			}
		}
	}
	fmt.Print("CountTheOnes ", count)
	return count
}

func IsPalindrome(str string) bool {
	isPalindrome := false
	if str == ReverseString(str) {
		isPalindrome = true
	}
	fmt.Print("IsPalindrome ", isPalindrome)
	return isPalindrome
}

func TwoNumberProducts(list []int) {
	products := []int{}

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			mul := list[i] * list[j]
			products = append(products, mul)
		}
	}

	fmt.Print("TwoNumberProducts ", products)
}

func CircularSentences(str string) bool {
	first := str[0]
	last := str[len(str)-1]

	words := strings.Split(str, " ")
	if first != last {
		return false
	}
	for i := 0; i < len(words)-1; i++ {
		currWord := words[i]
		nextWord := words[i+1]
		if currWord[len(currWord)-1] != nextWord[0] {
			log(false)
			return false
		}
	}
	log(true)
	return true
}

func PartitionEqualSubsetSum(list []int) bool {
	sorted := InsertionSort(list)
	greatest_num := sorted[len(sorted)-1]
	total_count := 0
	// l := []int{}
	for i := 0; i < len(list); i++ {
		if list[i] != greatest_num {
			// l = append(l, list[i])
			total_count = total_count + list[i]
		}
	}
	// m := []int{greatest_num}
	if total_count != greatest_num {
		log("false")
		return false
	}
	log("PartitionEqualSubsetSum")
	log(total_count)
	return true
}

func NoOfChars(list []string) {
	l := []string{}
	count := 1
	current := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] == current {
			count++
		} else {
			l = append(l, current, strconv.Itoa(count))
			current = list[i]
			count = 1
		}
	}
	l = append(l, current, strconv.Itoa(count))
	fmt.Println(l)
}

func NoOfCharStrings(input string) string {
	if len(input) == 0 {
		return ""
	}
	var result string
	current := input[0]
	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] == current {
			count++
		} else {
			result += fmt.Sprintf("%d%c", count, current)
			current = input[i]
			count = 1
		}
	}
	result += fmt.Sprintf("%d%c", count, current)
	log(result)
	return result
}

func EncodeStringChars(str string) string {
	var encoded strings.Builder
	count := 1

	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i] == str[i+1] {
			count++
		} else {
			encoded.WriteString(strconv.Itoa(count))
			encoded.WriteByte(str[i])
			count = 1
		}
	}
	log(encoded.String())
	return encoded.String()
}

func DecodeStringChars(str string) string {
	var decoded strings.Builder
	var count int

	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			count = count*10 + int(str[i]-'0')
		} else {
			decoded.WriteString(strings.Repeat(string(str[i]), count))
			count = 0
		}
	}
	log(decoded.String())
	return decoded.String()
}

// two sum problem [1, 2, 3] - 5 = [2, 3]

func TwoSumProblemBrute(list []int, target int) []int {
	l := []int{}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == target {
				l = append(l, list[i], list[j])
				log(l)
				return l
			}
		}
	}
	return l
}

func TwoSumProblemOp(list []int, target int) []int {
	low := 0
	high := len(list) - 1

	for low < high {
		sum := list[low] + list[high]
		if sum == target {
			log([]int{list[low], list[high]})
			return []int{list[low], list[high]}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return []int{}
}
