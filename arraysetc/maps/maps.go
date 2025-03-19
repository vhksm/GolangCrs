package maps

import (
	"fmt"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://www.aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	websites["Linkedin"] = "https://www.linkedin.com.br"
	fmt.Println(websites)
}
