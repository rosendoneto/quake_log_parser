package main

import (
	"fmt"
	"sort"
	"strings"
)

func generateReports(games []Game) {
	for i, game := range games {
		fmt.Printf("Game-%d:\n", i+1)
		fmt.Printf("  Total Kills: %d\n", game.TotalKills)
		fmt.Printf("  Players: %s\n", keysToString(game.Players))
		fmt.Printf("  Kills:\n")
		for player, kills := range game.Kills {
			fmt.Printf("    %s: %d\n", player, kills)
		}
		fmt.Printf("  Kills by Means:\n")
		for means, count := range game.KillsByMeans {
			fmt.Printf("    %s: %d\n", means, count)
		}
		fmt.Println()
	}
}

func keysToString(m map[string]struct{}) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return strings.Join(keys, ", ")
}

func generatePlayerRanking(games []Game) {
	playerScores := make(map[string]int)
	for _, game := range games {
		for player, kills := range game.Kills {
			playerScores[player] += kills
		}
	}

	type playerScore struct {
		player string
		score  int
	}
	var rankedPlayers []playerScore
	for player, score := range playerScores {
		rankedPlayers = append(rankedPlayers, playerScore{player, score})
	}

	sort.Slice(rankedPlayers, func(i, j int) bool {
		return rankedPlayers[i].score > rankedPlayers[j].score
	})

	fmt.Println("Player Ranking:")
	for _, ps := range rankedPlayers {
		fmt.Printf("  %s: %d\n", ps.player, ps.score)
	}
}
