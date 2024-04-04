package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

type TemperatureData struct {
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])

	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:] // removing element from slice
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println("Merged two dynamic lists: ", prices)

}

func main_old() {
	var productNames [4]string
	productNames = [4]string{"A Book"}

	productNames[2] = "A carpet"

	fmt.Println(productNames)

	prices := [4]float64{10.99, 9.99, 20.0, 14.29}
	fmt.Println(prices)
	fmt.Println(prices[2])

	featuredPrices := prices[1:3]
	featuredPrices[0] = 199.99
	fmt.Println("Sliced prices: ", featuredPrices)
	fmt.Println("Prices after updated nslice: ", prices)

	featuredPrices = prices[:3]
	fmt.Println("Sliced prices without start: ", featuredPrices)

	fmt.Println(len(featuredPrices), cap(featuredPrices))
}
