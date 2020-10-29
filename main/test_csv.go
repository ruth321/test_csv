package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("main.csv")
	if err != nil {
		log.Fatal(err)
	}
	a(file, 0)
	a(file, 0)
	//_, err = file.Seek(21, 0)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//reader := bufio.NewReader(file)
	//bytes, err := reader.Peek(1033)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(bytes))
	//_, err = file.Seek(0, 1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//reader = bufio.NewReader(file)
	//bytes, err = reader.Peek(81)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(bytes))
	//defer file.Close()
	//r := csv.NewReader(file)
	//records, err := r.ReadAll()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for g := 0; g < len(records[0]); g++ {
	//	fmt.Printf("%4d ", g)
	//	for i := 0; i < 8; i++ {
	//		fmt.Printf("|%s", records[i][g])
	//	}
	//	fmt.Println()
	//}
}

func a(file *os.File, g int64) {
	length := 84
	reader := bufio.NewReader(file)
	bytes, err := reader.Peek(length)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	for bytes[len(bytes)-1] != 13 {
		length++
		bytes, err = reader.Peek(length)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(string(bytes[:len(bytes)-1]))
}
