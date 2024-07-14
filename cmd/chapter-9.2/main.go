package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Define a custom error type to represent an empty field error.
// This error should include the name of the empty `Employee` field.
// In `main`, use `errors.As` to check for this error.
// Print out a message that includes the field name.
func main() {
	d := json.NewDecoder(strings.NewReader(data))
	var emptyErr EmptyFieldErr

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
			if errors.As(err, &emptyErr) {
				fmt.Println("We've found an empty err", err)
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

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

type EmptyFieldErr struct {
	Field   string
	Message string
}

func (ef EmptyFieldErr) Error() string {
	return ef.Message
}

func ValidateEmployee(e Employee) error {
	if len(e.ID) == 0 {
		return EmptyFieldErr{Field: "ID", Message: "missing ID"}
	}
	if !validID.MatchString(e.ID) {
		return errors.New("invalid ID")
	}
	if len(e.FirstName) == 0 {
		return EmptyFieldErr{Field: "FirstName", Message: "missing FirstName"}
	}
	if len(e.LastName) == 0 {
		return EmptyFieldErr{Field: "LastName", Message: "missing LastName"}
	}
	if len(e.Title) == 0 {
		return EmptyFieldErr{Field: "Title", Message: "missing Title"}
	}
	return nil
}
