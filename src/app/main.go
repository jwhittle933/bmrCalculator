package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var reader = bufio.NewReader(os.Stdin)

type user struct {
	Age      float64
	Weight   float64
	Height   float64
	Activity float64
	Gender   string
}

func feetToInches(feet int) (inches int) {
	inches = feet * 12
	return inches
}

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

func mascBmr(weight, height, age, activity float64) float64 {
	bmr := math.Round((66 + (6.3 * weight) + (12.9 * height) - (6.8 * age)) * activity)
	return bmr
}

func femBmr(weight, height, age, activity float64) float64 {
	bmr := math.Round((655 + (4.3 * weight) + (4.7 * height) - (4.7 * age)) * activity)
	return bmr
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	inputs := user{getAge(scanner), getWeight(scanner), getHeight(scanner), getActivity(scanner), getGender(reader)}

	if inputs.Gender == "male" {
		fmt.Println("You should eat around", mascBmr(inputs.Age, inputs.Height, inputs.Age, inputs.Activity), "calories per day.")
	} else {
		fmt.Println("You should eat around", femBmr(inputs.Weight, inputs.Height, inputs.Age, inputs.Activity), "calories per day.")
	}
}
