package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const inputFilePath = "input.txt"

type Gear struct {
	partNumberCount int
	product         int
}

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	gearProductMap := make(map[string]*Gear)

	for lineIndex, line := range input {
		numberTrackerValue := 0
		numberTrackerStartIndex := -1
		numberTrackerEndIndex := -1
		symbolFoundForAns1 := false

		for trackNumIndex, ch := range line {
			num, err := strconv.Atoi(string(ch))

			if err == nil {
				numberTrackerValue = 10*numberTrackerValue + num

				if numberTrackerStartIndex == -1 {
					numberTrackerStartIndex = trackNumIndex
				}

				numberTrackerEndIndex = trackNumIndex
			}

			if err != nil || trackNumIndex == len(line)-1 {
				if numberTrackerValue > 0 {
					for checkCharIndex := numberTrackerStartIndex - 1; checkCharIndex <= numberTrackerEndIndex+1; checkCharIndex++ {
						if lineIndex > 0 && checkCharIndex > 0 && checkCharIndex < len(line)-1 {
							upperChar := string(input[lineIndex-1][checkCharIndex])
							if checkSpecialSymbol(upperChar) {
								symbolFoundForAns1 = true

								if upperChar == "*" {
									appendPartNumberToGearMap(gearProductMap, numberTrackerValue, lineIndex-1, checkCharIndex)
									break
								}
							}
						}

						if lineIndex < len(input)-1 && checkCharIndex > 0 && checkCharIndex < len(line)-1 {
							lowerChar := string(input[lineIndex+1][checkCharIndex])
							if checkSpecialSymbol(lowerChar) {
								symbolFoundForAns1 = true

								if lowerChar == "*" {
									appendPartNumberToGearMap(gearProductMap, numberTrackerValue, lineIndex+1, checkCharIndex)
									break
								}
							}
						}
					}

					if numberTrackerStartIndex > 0 {
						leftChar := string(input[lineIndex][numberTrackerStartIndex-1])
						if checkSpecialSymbol(leftChar) {
							symbolFoundForAns1 = true

							if leftChar == "*" {
								appendPartNumberToGearMap(gearProductMap, numberTrackerValue, lineIndex, numberTrackerStartIndex-1)
							}
						}
					}

					if numberTrackerEndIndex < len(line)-1 {
						rightChar := string(input[lineIndex][numberTrackerEndIndex+1])
						if checkSpecialSymbol(rightChar) {
							symbolFoundForAns1 = true

							if rightChar == "*" {
								appendPartNumberToGearMap(gearProductMap, numberTrackerValue, lineIndex, numberTrackerEndIndex+1)
							}
						}
					}
				}

				if symbolFoundForAns1 {
					answer1 += numberTrackerValue
				}

				numberTrackerValue = 0
				numberTrackerStartIndex = -1
				numberTrackerEndIndex = -1
				symbolFoundForAns1 = false
			}
		}
	}

	for _, val := range gearProductMap {
		if val.partNumberCount > 1 {
			answer2 += val.product
		}
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

func checkSpecialSymbol(ch string) bool {
	if ch == "." {
		return false
	}

	if _, err := strconv.Atoi(string(ch)); err != nil {
		return true
	}

	return false
}

func appendPartNumberToGearMap(gearProductMap map[string]*Gear, partNumber int, specialSymbolLineIndex int, specialSymbolCharIndex int) {
	specialSymbolCoordinates := string(specialSymbolLineIndex) + ":" + string(specialSymbolCharIndex)

	if _, ok := gearProductMap[specialSymbolCoordinates]; !ok {
		gearProductMap[specialSymbolCoordinates] = &Gear{
			partNumberCount: 1,
			product:         partNumber,
		}
	} else {
		gearProductMap[specialSymbolCoordinates].partNumberCount++
		gearProductMap[specialSymbolCoordinates].product *= partNumber
	}
}
