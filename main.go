package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Adam"
	userNames[1] = "Julia"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Michał")
	userNames = append(userNames, "Maciek")

	fmt.Println(userNames, " ", len(userNames))

	coursesRatings := make(floatMap, 3)

	coursesRatings["go"] = 4.7
	coursesRatings["react"] = 4.8
	coursesRatings["angular"] = 4.7

	coursesRatings.output()

	for index, value := range userNames {
		fmt.Println(index, " has value: ", value)
	}

	for key, value := range coursesRatings {
		fmt.Println("key: ", key, " has value: ", value)
	}
}
