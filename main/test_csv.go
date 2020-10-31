package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("main.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var records [][]string
	var str string
	for {
		str, err = a(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		strReader := strings.NewReader(str)
		csvReader := csv.NewReader(strReader)
		record, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record[0])
	}
	for g := 0; g < len(records[0]); g++ {
		fmt.Printf("%4d ", g)
		for i := 0; i < 8; i++ {
			fmt.Printf("%s|", records[i][g])
		}
		fmt.Println()
	}

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

func a(file *os.File) (string, error) {
	reader := bufio.NewReaderSize(file, 4096)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	_, err = file.Seek(int64(-reader.Buffered()), 1)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
