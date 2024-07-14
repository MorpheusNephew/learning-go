package main

import (
	"fmt"
	"slices"
	"strings"
)

// Add two methods to `League`. The first method is called `MatchResult`.
// It takes four parameters: the name of the first team, its score in the game, the name of the second team, and its score in the game.
// This method should update the `Wins` field in `League`.
// Add a second method to `League` called `Ranking` that returns a slice of the team names in order of wins.
// Build your data structures and call these methods from the `main` function in your program using some sample data.
func main() {
	miamiHeat := NewTeam("Miami Heat", "Jimmy Butler")
	atlantaHawks := NewTeam("Atlanta Hawks", "Trey Young")
	bostonCeltics := NewTeam("Boston Celtics", "Jaylon Brown", "Jason Tatum")
	newYorkKnicks := NewTeam("New York Knicks", "Jaylen Brunson")

	teamNames := []string{atlantaHawks.Name, bostonCeltics.Name, miamiHeat.Name, newYorkKnicks.Name}

	nba := NewLeague(teamNames...)

	fmt.Println(nba)

	nba.MatchResult(miamiHeat.Name, 70, bostonCeltics.Name, 34)
	nba.MatchResult(newYorkKnicks.Name, 35, atlantaHawks.Name, 56)
	nba.MatchResult(bostonCeltics.Name, 45, atlantaHawks.Name, 34)
	nba.MatchResult(miamiHeat.Name, 4, newYorkKnicks.Name, 2)
	nba.MatchResult(miamiHeat.Name, 6, atlantaHawks.Name, 3)
	nba.MatchResult(miamiHeat.Name, 6, bostonCeltics.Name, 3)
	nba.MatchResult(miamiHeat.Name, 6, atlantaHawks.Name, 8)
	nba.MatchResult(newYorkKnicks.Name, 6, atlantaHawks.Name, 9)

	fmt.Println(nba)
	fmt.Println(nba.Ranking())
	fmt.Println(nba)
}

func (l League) MatchResult(t1 string, s1 int, t2 string, s2 int) {
	if s1 > s2 {
		l.Wins[t1]++
	} else {
		l.Wins[t2]++
	}
}

func (l League) Ranking() []string {
	teamsCopy := slices.Clone(l.Teams)

	slices.SortFunc(teamsCopy, func(a, b string) int {
		return l.Wins[b] - l.Wins[a]
	})

	return teamsCopy
}

func NewLeague(teams ...string) League {
	teamsWins := TeamsWins{}

	for _, team := range teams {
		teamsWins[team] = 0
	}

	return League{Teams: teams, Wins: teamsWins}
}

type Team struct {
	Name    string
	Players []string
}

func NewTeam(name string, players ...string) Team {
	return Team{Name: name, Players: players}
}

func (t Team) String() string {
	return fmt.Sprintf("It's your %s!\nHere are your players! %s\n", t.Name, strings.Join(t.Players, ", "))
}

type TeamsWins map[string]int

type League struct {
	Teams []string
	Wins  TeamsWins
}

func (l League) String() string {
	var leagueInfo strings.Builder

	leagueInfo.WriteString("Welcome to the National Basketball Association!\n\n")

	for _, teamName := range l.Teams {
		leagueInfo.WriteString(fmt.Sprintf("%s with %d wins\n", teamName, l.Wins[teamName]))
	}

	return leagueInfo.String()
}
