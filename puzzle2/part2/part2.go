package part2

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
	minRed, minGreen, minBlue := 0, 0, 0

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

			switch cubeColor {
			case "red":
				if cubeNumber > minRed {
					minRed = cubeNumber
				}
			case "green":
				if cubeNumber > minGreen {
					minGreen = cubeNumber
				}
			case "blue":
				if cubeNumber > minBlue {
					minBlue = cubeNumber
				}
			}
		}
	}
	power := minRed * minGreen * minBlue
	return power
}
