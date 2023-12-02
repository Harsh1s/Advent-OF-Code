package part1

import (
	"strconv"
	"strings"
)

func Solve(s string) int {
	s = strings.TrimSpace(s)
	gameData := strings.Split(s, ":")
	if len(gameData) < 2 {
		panic("invalid game data")
	}

	gameNumberData := strings.Split(strings.TrimSpace(gameData[0]), " ")
	if len(gameNumberData) < 2 {
		panic("invalid game number data")
	}

	gameNumber, err := strconv.Atoi(gameNumberData[1])
	if err != nil {
		panic(err)
	}

	hands := strings.Split(strings.TrimSpace(gameData[1]), ";")
	for _, hand := range hands {
		handData := strings.Split(strings.TrimSpace(hand), ",")
		for _, cubes := range handData {
			cubeData := strings.Split(strings.TrimSpace(cubes), " ")
			if len(cubeData) < 2 {
				panic("invalid cube data")
			}

			cubeColor := cubeData[1]
			cubeNumber, err := strconv.Atoi(cubeData[0])
			if err != nil {
				panic(err)
			}

			if cubeColor == "red" && cubeNumber > 12 {
				return 0
			}
			if cubeColor == "green" && cubeNumber > 13 {
				return 0
			}
			if cubeColor == "blue" && cubeNumber > 14 {
				return 0
			}
		}
	}

	return gameNumber
}
