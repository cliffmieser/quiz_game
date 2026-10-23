// quiz game

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
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

func get_records(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func checkFileExists(filePath string) (string, bool) {
	_, error := os.Stat(filePath)
	return filePath, !errors.Is(error, os.ErrNotExist)
}

func start_quiz(records [][]string, q *quizzer) {
	var s int
	fmt.Printf("Enter seconds (integer): ")
	fmt.Scanf("%d", &s)

	if s <= 0 {
		s = 30
	}

	defer show_score(q) // will be called on any return path

	timer := time.NewTimer(time.Duration(s) * time.Second) // stops program after s seconds
	defer timer.Stop()

	results := make(chan bool) // true -> answered correctly

	go func() { // this runs the loop of questions
		for i, record := range records {
			var answer int
			fmt.Printf("Problem #%d: %v = ", i+1, record[0])
			fmt.Scanf("%d", &answer)
			results <- strconv.Itoa(answer) == record[1]
			q.total++
		}
		close(results) // no more questions
	}()

	// while the above go routine provides questions, we wait for correct/incorrect
	// answers from the results channel, which determines how to update the score
	for {
		select {
		case correct, ok := <-results:
			if !ok {
				return // all questions answered
			}
			q.total++
			if correct {
				q.correct++
			} else {
				q.incorrect++
			}
		case <-timer.C:
			fmt.Println("\nTime's up!")
			return
		}
	}
}

func main() {
	// check if problems.csv file exists, if not, create one, skip if it exists
	fileName, fileExists := checkFileExists("problems.csv")
	var records [][]string

	if fileExists {
		var err error
		records, err = get_records(fileName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		records = [][]string{}
		make_fields(&records, 5)
	}

	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	if !fileExists {
		make_csv(&records)
	}

	start_quiz(records, &quizzer{0, 0, 0})

}

/*
Next: Add timers that are customizable with flags

*/
