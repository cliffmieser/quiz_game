// quiz game

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func make_fields(r *[][]string, num_records int) {
	/*
		Generates num_records amount of {question,answer} fields
		r is a pointer to a slice of string slices
	*/
	for i := 0; i <= num_records; i++ {
		// generate a field in format {quesiton, answer}
		op1 := rand.Intn(11)
		op2 := rand.Intn(11)

		// record := strconv.Itoa(op1) + "+" + strconv.Itoa(op2) + "," + strconv.Itoa(op1+op2)
		qa := strconv.Itoa(op1) + "+" + strconv.Itoa(op2)
		record := []string{qa, strconv.Itoa(op1 + op2)}
		fmt.Println(record)

		// record_slice := []string{record}
		// records = append(records, record_slice)
		*r = append(*r, record)

	}
	fmt.Println(*r)
}

func make_csv(r *[][]string) {
	/* Creates csv file and writes all of {question, answer} records
	To the csv*/
	f, err := os.Create("problems.csv")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(*r) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	num_records := 5

	records := [][]string{} // create slice of record slices
	make_fields(&records, num_records)
	make_csv(&records)

}
