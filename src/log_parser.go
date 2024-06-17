package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func readLogFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type Game struct {
	TotalKills   int
	Players      map[string]struct{}
	Kills        map[string]int
	KillsByMeans map[string]int
}

func parseLog(lines []string) []Game {
	var games []Game
	var currentGame *Game
	killPattern := regexp.MustCompile(`Kill: \d+ \d+ \d+: (.*) killed (.*) by (.*)`)

	for _, line := range lines {
		if strings.Contains(line, "InitGame") {
			if currentGame != nil {
				games = append(games, *currentGame)
			}
			currentGame = &Game{
				Players:      make(map[string]struct{}),
				Kills:        make(map[string]int),
				KillsByMeans: make(map[string]int),
			}
		} else if strings.Contains(line, "Kill:") {
			currentGame.TotalKills++
			match := killPattern.FindStringSubmatch(line)
			if match != nil {
				killer, victim, means := match[1], match[2], match[3]
				if killer == "<world>" {
					currentGame.Kills[victim]--
				} else {
					currentGame.Players[killer] = struct{}{}
					currentGame.Players[victim] = struct{}{}
					currentGame.Kills[killer]++
				}
				currentGame.KillsByMeans[means]++
			}
		}
	}
	if currentGame != nil {
		games = append(games, *currentGame)
	}

	return games
}
