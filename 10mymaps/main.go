package main

import "fmt"

func main() {
	fmt.Println("welcome to maps")

	languages := make(map[string]string)

	languages["JS"] = "Java Script"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS is short of: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//loops are interesting in go lang

	for key, value := range languages {
		fmt.Printf("For Key %v language is %v\n", key, value)
	}
}
