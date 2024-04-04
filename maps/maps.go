package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"Google":              "http://google.com",
		"Amazon Web Services": "http://aws.com",
	}
	fmt.Println(websites)

	fmt.Println("Selecting AWS url:", websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println("Addition of new key:", websites)

	delete(websites, "Google")
	fmt.Println("After deletion Google key:", websites)

}
