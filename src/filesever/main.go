package main

import "fmt"

func main() {
	//http.ListenAndServe(":8090", http.FileServer(http.Dir(".")))
	var (
		i int = 2
		j int = 3
	)
	fmt.Println(i, j)
}
