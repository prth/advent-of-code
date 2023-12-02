package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const inputFilePath = "input.txt"

var digitsToWordMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var digitWordLengths = [3]int{3, 4, 5}

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	for _, line := range input {
		firstDigitByOnlyNum := -1
		lastDigitByOnlyNum := -1

		firstDigitByNumOrWord := -1
		lastDigitByNumOrWord := -1

		for i := 0; i < len(line); i++ {
			num, err := strconv.Atoi(string(line[i]))
			eachDigitOnlyNum := -1
			eachDigitNumOrWord := -1

			if err == nil {
				eachDigitOnlyNum = num
				eachDigitNumOrWord = num
			} else {
				for _, wordLength := range digitWordLengths {
					if len(line)-i >= wordLength {
						if num, ok := digitsToWordMap[string(line[i:i+wordLength])]; ok {
							eachDigitNumOrWord = num
						}
					}
				}
			}

			if eachDigitOnlyNum != -1 {
				if firstDigitByOnlyNum == -1 {
					firstDigitByOnlyNum = eachDigitOnlyNum
				}

				lastDigitByOnlyNum = eachDigitOnlyNum
			}

			if eachDigitNumOrWord != -1 {
				if firstDigitByNumOrWord == -1 {
					firstDigitByNumOrWord = eachDigitNumOrWord
				}

				lastDigitByNumOrWord = eachDigitNumOrWord
			}
		}

		if firstDigitByOnlyNum == -1 || lastDigitByOnlyNum == -1 {
			log.Fatalf("Invalid state by only num | line=[%s]", line)
		}

		if firstDigitByNumOrWord == -1 || lastDigitByNumOrWord == -1 {
			log.Fatalf("Invalid state by num or word | line=[%s]", line)
		}

		answer1 += firstDigitByOnlyNum*10 + lastDigitByOnlyNum
		answer2 += firstDigitByNumOrWord*10 + lastDigitByNumOrWord
	}

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getInput() ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input, scanner.Err()
}
