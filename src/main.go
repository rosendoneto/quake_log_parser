package main

import (
	"fmt"
)

func main() {
	logFilePath := "data/games.log"
	logLines, err := readLogFile(logFilePath)
	if err != nil {
		fmt.Printf("Error reading log file: %v\n", err)
		return
	}

	games := parseLog(logLines)
	generateReports(games)
	generatePlayerRanking(games)
}
