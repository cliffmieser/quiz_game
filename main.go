// quiz game

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type quizzer struct {
	correct   int
	incorrect int
	// records   [][]string
}

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
	// fmt.Println(*r)
}

func make_csv(r *[][]string) string {
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

	return f.Name() // returns the name of the csv
}

func start_quiz(fileName string, q *quizzer) {

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	re := csv.NewReader(f) // uses re for reader to not conflict with r ptr

	records, err := re.Read()

	if err == io.EOF { // end of file reached
		return // end of file
	}

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print(records)

	for value := range records {
		fmt.Printf("%s\n", records[value])
	}
	fmt.Printf("Correct: %d\nIncorrect: %d\n", q.correct, q.incorrect)

}

func main() {
	num_records := 5

	records := [][]string{} // create slice of record slices
	make_fields(&records, num_records)
	fileName := make_csv(&records)

	start_quiz(fileName, &quizzer{0, 0})

}
