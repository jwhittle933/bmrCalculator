package bodymeasures

import (
	"bufio"
	"fmt"
	"strconv"
)

func getAge(scanner *bufio.Scanner) float64 {
	fmt.Println("What's your age?")
	scanner.Scan()
	age, _ := strconv.ParseFloat(scanner.Text(), 64)
	return age
}
func getWeight(scanner *bufio.Scanner) float64 {
	fmt.Println("What's your weight in pounds?")
	scanner.Scan()
	weight, _ := strconv.ParseFloat(scanner.Text(), 64)
	return weight
}
func getHeight(scanner *bufio.Scanner) float64 {
	fmt.Println("What's your height in pounds?")
	scanner.Scan()
	height, _ := strconv.ParseFloat(scanner.Text(), 64)
	return height
}
func getActivity(scanner *bufio.Scanner) float64 {
	fmt.Println("What's your activity level per day(between 1.0 and 2.0)?")
	scanner.Scan()
	height, _ := strconv.ParseFloat(scanner.Text(), 64)
	return height
}

func getGender(reader *bufio.Reader) string {
	fmt.Println("Are you male or female?")
	gender, _ := reader.ReadString('\n')
	return gender
}
