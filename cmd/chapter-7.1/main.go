package main

import (
	"fmt"
	"strings"
)

// You have been asked to manage a basketball league and are going to write a program to help you.
// Define two types.
// The first one, called `Team`, has a field for the name of the team and a field for the player names.
// The second type is called `League` and has a field called `Teams` for the teams in the league
// and a field called `Wins` that maps a team’s name to its number of wins.
func main() {
	miamiHeat := NewTeam("Miami Heat", "Jimmy Buckets", "Tyler Herro")
	orlandoMagic := NewTeam("Orlando Magic", "Victor Oladipo")
	nba := NewLeague(map[string]int{miamiHeat.Name: 10, orlandoMagic.Name: 3}, miamiHeat.Name, orlandoMagic.Name)

	fmt.Println(nba)
	fmt.Println(miamiHeat)
	fmt.Println(orlandoMagic)
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

type League struct {
	Teams []string
	Wins  map[string]int
}

func NewLeague(wins map[string]int, teams ...string) League {
	return League{Teams: teams, Wins: wins}
}

func (l League) String() string {
	var leagueInfo strings.Builder

	leagueInfo.WriteString("Welcome to the National Basketball Association!\n\n")

	for _, teamName := range l.Teams {
		leagueInfo.WriteString(fmt.Sprintf("%s with %d wins\n", teamName, l.Wins[teamName]))
	}

	return leagueInfo.String()
}
