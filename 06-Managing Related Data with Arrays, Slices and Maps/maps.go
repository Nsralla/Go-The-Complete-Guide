package main

import "fmt"

func main(){
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Twitter": "https://www.twitter.com",
	}

	fmt.Println("Websites:")
	for name, url := range websites {
		fmt.Printf("%s : %s \n", name, url)
	}

	fmt.Println("\n accessing a specigic key : ", websites["Google"])

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println("\n After adding LinkedIn : ", websites)

	// to delete a key
	delete(websites, "Twitter")
	fmt.Println("\n After deleting Twitter : ", websites)
}