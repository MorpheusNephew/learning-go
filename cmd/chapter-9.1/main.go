package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

/*
Create a sentinel error to represent an invalid ID.
In `main`, use `errors.Is` to check for the sentinel error,
and print a message when it is found.
*/
func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		if err != nil {
			fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			if errors.Is(err, ErrInvalidID) {
				fmt.Println("Hey we found the invalid id error")
			}
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

type Sentinel string

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

const ErrInvalidID = Sentinel("invalid ID")

func (s Sentinel) Error() string {
	return string(s)
}

func ValidateEmployee(e Employee) error {
	if len(e.ID) == 0 {
		return errors.New("missing ID")
	}
	if !validID.MatchString(e.ID) {
		return ErrInvalidID
	}
	if len(e.FirstName) == 0 {
		return errors.New("missing FirstName")
	}
	if len(e.LastName) == 0 {
		return errors.New("missing LastName")
	}
	if len(e.Title) == 0 {
		return errors.New("missing Title")
	}
	return nil
}
