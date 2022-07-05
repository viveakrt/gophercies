package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Ques struct {
	ques string
	ans  string
}

func main() {
	fileName := flag.String("csv", "problems", "csv is used to pass the name of file in csv")
	shuffle := flag.Bool("s", false, "used to shuffle the questions")
	limit := flag.Int("t", 30, "time in seconds for quiz")
	flag.Parse()
	data := readCsvFileQues(*fileName)
	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	}

	fmtGreen("Quiz GAME ")
	fmtPurple("Press Any Key to Start the game")
	var new string
	fmt.Scanln(&new)

	timeNow := time.NewTimer(time.Duration(*limit) * time.Second)
	count := 0
problemLoop:
	for _, question := range data {
		fmtBlue(question.ques)
		ansChan := make(chan string)
		go func() {
			var sol string
			fmt.Scanf("%s\n", &sol)
			ansChan <- sol
		}()
		select {
		case <-timeNow.C:
			fmtGreen("")
			break problemLoop
		case ans := <-ansChan:
			if ans == question.ans {
				count += 1
			}
		}
	}

	fmtGreen("You score " + strconv.Itoa(count) +
		" out of " + strconv.Itoa(len(data)))

}

func readCsvFileQues(filePath string) []Ques {
	filePath = "../" + filePath + ".csv"
	file, err := os.Open(filePath)
	if err != nil {
		logError("While Opening file", "readCsvFile")
	}
	r := csv.NewReader(file)
	var questions []Ques
	var q Ques
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		q.ques = record[0]
		q.ans = record[1]
		questions = append(questions, q)
	}
	file.Close()
	return questions
}

func logError(message string, function string) {
	var reset = "\033[0m"
	var red = "\033[31m"
	var blue = "\033[34m"
	var white = "\033[97m"
	fmt.Println("{" + white + "Error : " + red + message +
		white + " |" + white + " Method : " + blue +
		function + reset + "}\n")
}
func fmtGray(s string) {
	var reset = "\033[0m"
	var gray = "\033[37m"
	fmt.Println(gray + s + reset)
}
func fmtCyan(s string) {
	var reset = "\033[0m"
	var cyan = "\033[36m"
	fmt.Println(cyan + s + reset)
}
func fmtPurple(s string) {
	var reset = "\033[0m"
	var purple = "\033[35m"
	fmt.Println(purple + s + reset)
}
func fmtBlue(s string) {
	var reset = "\033[0m"
	var blue = "\033[34m"
	fmt.Println(blue + s + reset)
}
func fmtRed(s string) {
	var reset = "\033[0m"
	var red = "\033[31m"
	fmt.Println(red + s + reset)
}
func fmtGreen(s string) {
	var reset = "\033[0m"
	var green = "\033[32m"
	fmt.Println(green + s + reset)
}
func fmtYellow(s string) {
	var reset = "\033[0m"
	var Green = "\033[33m"
	fmt.Println(Green + s + reset)
}
