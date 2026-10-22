// quiz game

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type quizzer struct {
	/* represents a scoreboard of right/wrong answers */
	total     int
	correct   int
	incorrect int
}

func show_score(q *quizzer) {
	// prints resulting score
	fmt.Printf("\nTotal Questions: %d\nCorrect: %d\nIncorrect: %d\n", q.total, q.correct, q.incorrect)

}

func make_fields(r *[][]string, num_records int) {
	/*
		Generates num_records amount of {question,answer} fields
		r is a pointer to a slice of string slices
	*/
	for i := 0; i < num_records; i++ {
		/* generate a field in format {quesiton, answer} */
		op1 := rand.Intn(11)
		op2 := rand.Intn(11)

		qa := strconv.Itoa(op1) + "+" + strconv.Itoa(op2)
		record := []string{qa, strconv.Itoa(op1 + op2)}
		*r = append(*r, record)

	}
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

func get_num_records(filename string) int {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	return len(records)

}

func start_quiz(fileName string, q *quizzer, num_records int) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	re := csv.NewReader(f) // uses re for reader to not conflict with r ptr

	for i := 0; i < num_records; i++ {
		records, err := re.Read()
		if err == io.EOF { // end of file reached
			break // end of file
		}
		if err != nil {
			log.Fatal(err)
		}

		var answer int
		fmt.Printf("Problem #%d: %v = ", i+1, records[0])
		fmt.Scanf("%d", &answer)

		q.total++
		if strconv.Itoa(answer) == records[1] {
			q.correct++
		} else {
			q.incorrect++
		}
	}
	show_score(q)

}

func checkFileExists(filePath string) (string, bool) {
	_, error := os.Stat(filePath)
	return filePath, !errors.Is(error, os.ErrNotExist)
}

func main() {
	// check if problems.csv file exists, if not, create one, skip if it exists
	// var fileName, fileExists := checkFileExists("problems.csv")
	var fileName string
	var fileExists bool
	var num_records int
	fileName, fileExists = checkFileExists("problems.csv")

	if fileExists {
		num_records = get_num_records(fileName)
		start_quiz(fileName, &quizzer{0, 0, 0}, num_records)
	} else {
		num_records = 5                    // creates n number of records;
		records := [][]string{}            // create slice of record slices
		make_fields(&records, num_records) // create question,answer field
		fileName = make_csv(&records)

	}

	start_quiz(fileName, &quizzer{0, 0, 0}, num_records)

}
