package main

import "fmt"

/*
You have been asked to manage a basketball league and are going to write a program to help you.
Define two types.
The first one, called `Team`, has a field for the name of the team and a field for the player names.
The second type is called `League` and has a field called `Teams` for the teams in the league
and a field called `Wins` that maps a team’s name to its number of wins.
*/
func main() {
	miamiHeat := Team{Name: "Miami Heat", Players: []string{"Jimmy Buckets"}}
	nba := League{Teams: []string{miamiHeat.Name}, Wins: map[string]int{miamiHeat.Name: 10}}

	fmt.Println(miamiHeat, nba, nba.Wins[miamiHeat.Name])
}

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []string
	Wins  map[string]int
}
