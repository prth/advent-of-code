package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

type CubePack struct {
	color string
	count int
}

type Set struct {
	cubePacks []CubePack
}

type Game struct {
	index int
	sets  []Set
}

type ValidGameConfiguration struct {
	colorWiseCubeCountMap map[string]int
}

var validGameConfiguration = ValidGameConfiguration{
	colorWiseCubeCountMap: map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	},
}

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	for _, game := range input {
		isValidGameForAns1 := true
		gameFewestCubesCountMap := make(map[string]int)

		for _, set := range game.sets {
			for _, cubePack := range set.cubePacks {
				if isValidGameForAns1 {
					if colorMaxCountConfig, ok := validGameConfiguration.colorWiseCubeCountMap[cubePack.color]; !ok {
						log.Fatalf("Invalid state | color config not found | color=[%s]", cubePack.color)
					} else {
						if cubePack.count > colorMaxCountConfig {
							isValidGameForAns1 = false
						}
					}
				}

				if _, ok := gameFewestCubesCountMap[cubePack.color]; !ok {
					gameFewestCubesCountMap[cubePack.color] = cubePack.count
				} else {
					if gameFewestCubesCountMap[cubePack.color] < cubePack.count {
						gameFewestCubesCountMap[cubePack.color] = cubePack.count
					}
				}
			}
		}

		if isValidGameForAns1 {
			answer1 += game.index
		}

		gameProductForAns2 := 1

		for _, count := range gameFewestCubesCountMap {
			gameProductForAns2 *= count
		}

		answer2 += gameProductForAns2
	}

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getInput() ([]Game, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []Game

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parseGame(line))
	}

	return input, scanner.Err()
}

func parseGame(gameStr string) Game {
	indexSetSeparatorPosition := strings.Index(gameStr, ":")
	gameIndex, _ := strconv.Atoi(gameStr[5:indexSetSeparatorPosition])

	var sets []Set

	for _, setStr := range strings.Split(gameStr[indexSetSeparatorPosition+1:], ";") {
		sets = append(sets, parseSet(setStr))
	}

	return Game{
		index: gameIndex,
		sets:  sets,
	}
}

func parseSet(setStr string) Set {
	var cubePacks []CubePack

	for _, cubePackStr := range strings.Split(strings.TrimSpace(setStr), ",") {
		cubePacks = append(cubePacks, parseCubePack(cubePackStr))
	}

	return Set{
		cubePacks: cubePacks,
	}
}

func parseCubePack(cubePackStr string) CubePack {
	colorPackElements := strings.Split(strings.TrimSpace(cubePackStr), " ")

	count, _ := strconv.Atoi(colorPackElements[0])

	return CubePack{
		color: strings.ToLower(colorPackElements[1]),
		count: count,
	}
}
