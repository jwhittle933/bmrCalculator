package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func handleInput(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func feetToInches(feet int) (inches int) {
	inches = feet * 12
	return inches
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
	type User struct {
		Age      float64
		Weight   float64
		Height   float64
		Activity float64
		Gender   string
	}

	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What's your age?")
	scanner.Scan()
	age, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("What's your weight in pounds?")
	scanner.Scan()
	weight, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("What's your height in inches?")
	scanner.Scan()
	height, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("What's your activity level per day(between 1.0 and 2.0)?")
	scanner.Scan()
	activity, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Are you male or female?")
	gender, _ := reader.ReadString('\n')

	inputs := User{age, weight, height, activity, gender}

	if inputs.Gender == "male" {
		fmt.Println("You should eat around", mascBmr(inputs.Weight, inputs.Height, inputs.Age, inputs.Activity), "calories per day.")
	} else {
		fmt.Println("You should eat around", femBmr(inputs.Weight, inputs.Height, inputs.Age, inputs.Activity), "calories per day.")
	}
}
