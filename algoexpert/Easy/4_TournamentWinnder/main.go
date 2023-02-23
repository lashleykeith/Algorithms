// O(n) time | O(k) space

package main

const HOME_TEAM_WON = 1

func updateScores(team string, points int, scores map[string]int){
    scores[team] += points
}

func TournamentWinner(competitions [][]string, results []int) string{
    currentBestTeam := ""
    scores := make(map[string]int)

    for idx, competition := range competitions {
        result := results[idx]
        homeTeam, awayTeam := competition[0], competition[1]

        winningTeam := awayTeam

        if result == HOME_TEAM_WON{
            winningTeam = homeTeam
        }

        updateScores(winningTeam, 3, scores)

        if scores[winningTeam] > scores[currentBestTeam]{
            currentBestTeam = winningTeam
        }
    }
    return currentBestTeam
}
