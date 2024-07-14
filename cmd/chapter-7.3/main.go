package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Define an interface called `Ranker` that has a single method called `Ranking` that returns a slice of strings.
// Write a function called `RankPrinter` with two parameters, the first of type `Ranker` and the second of type `io.Writer`.
// Use the `io.WriteString` function to write the values returned by `Ranker` to the `io.Writer`, with a newline separating each result.
// Call this function from `main`.
func main() {
	nba := SetupLeague()

	RankPrinter(nba, os.Stdout)
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {

	ranking := r.Ranking()

	for _, rank := range ranking {
		io.WriteString(w, fmt.Sprintf("%s\n", rank))
	}
}

func SetupLeague() League {
	miamiHeat := NewTeam("Miami Heat", "Jimmy Butler")
	atlantaHawks := NewTeam("Atlanta Hawks", "Trey Young")
	bostonCeltics := NewTeam("Boston Celtics", "Jaylon Brown", "Jason Tatum")
	newYorkKnicks := NewTeam("New York Knicks", "Jaylen Brunson")

	teamNames := []string{atlantaHawks.Name, bostonCeltics.Name, miamiHeat.Name, newYorkKnicks.Name}

	nba := NewLeague(teamNames...)

	nba.MatchResult(miamiHeat.Name, 70, bostonCeltics.Name, 34)
	nba.MatchResult(newYorkKnicks.Name, 35, atlantaHawks.Name, 56)
	nba.MatchResult(bostonCeltics.Name, 45, atlantaHawks.Name, 34)
	nba.MatchResult(miamiHeat.Name, 4, newYorkKnicks.Name, 2)
	nba.MatchResult(miamiHeat.Name, 6, atlantaHawks.Name, 3)
	nba.MatchResult(miamiHeat.Name, 6, bostonCeltics.Name, 3)
	nba.MatchResult(miamiHeat.Name, 6, atlantaHawks.Name, 8)
	nba.MatchResult(newYorkKnicks.Name, 6, atlantaHawks.Name, 9)

	return nba
}

// From the previous exercise
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
