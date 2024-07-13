package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

/*
Rather than returning the first error found,
return back a single error that contains all errors discovered during validation.
Update the code in `main` to properly report multiple errors.
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
			switch err := err.(type) {
			case interface{ Unwrap() error }:
				innerErr := err.Unwrap()
				fmt.Printf("internal error: %v", innerErr)
			case interface{ Unwrap() []error }:
				innerErrs := err.Unwrap()

				fmt.Println("Here are the inner errors:")
				for i, err := range innerErrs {
					fmt.Printf("%d: %v\n", i+1, err)
				}

			default:
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
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

func ValidateEmployee(e Employee) error {
	ve := ValidationError{}

	if len(e.ID) == 0 {
		ve.Errors = append(ve.Errors, errors.New("missing ID"))
	}
	if !validID.MatchString(e.ID) {
		ve.Errors = append(ve.Errors, errors.New("invalid ID"))
	}
	if len(e.FirstName) == 0 {
		ve.Errors = append(ve.Errors, errors.New("missing FirstName"))
	}
	if len(e.LastName) == 0 {
		ve.Errors = append(ve.Errors, errors.New("missing LastName"))
	}
	if len(e.Title) == 0 {
		ve.Errors = append(ve.Errors, errors.New("missing Title"))
	}

	if len(ve.Errors) > 0 {
		return ve
	}

	return nil
}

type ValidationError struct {
	Errors []error
}

func (ve ValidationError) Error() string {
	return errors.Join(ve.Errors...).Error()
}

func (ve ValidationError) Unwrap() []error {
	return ve.Errors
}
