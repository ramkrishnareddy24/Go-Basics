package maps

import "fmt"

func main(){
	websites := map[string]string{
		"Google": "https://google.com",
		"Aws": "https://amazonwebservices.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Aws"])
	websites["LinkedIn"] = "https://linkedin.com"

	delete(websites,"Google")
	fmt.Println(websites)
}

