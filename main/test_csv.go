package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("main.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for g := 0; g < len(records[0]); g++ {
		fmt.Printf("%4d ", g)
		for i := 0; i < 2; i++ {
			fmt.Printf("%s| ", records[i][g])
		}
		fmt.Println()
	}

}
