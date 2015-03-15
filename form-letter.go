package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was great that you attended fundraiser.{{else}}
It is sad that you couldn't attend the fundraiser.
Also, there is Air show on Wednesday, Agro-fair on Friday, and Fashion show on Sunday.{{end}}
{{with .Donate}}Thank you for the donation {{.}}.
Also, there is Air show on Wednesday, Agro-fair on Friday, and Fashion show on Sunday.
{{end}}
Best wishes,
Josie
`




// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Donate string
		Attended   bool

	}
	var recipients = []Recipient{
		{"Ms Mildred", "100 dollars", true },
		{"Mr Johnson", "50 dollars", false},
		{"Mr Rodney", "", false},
	}

	// STEP 1 & STEP 2
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	//STEP 3
	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}

// FROM : http://golang.org/pkg/text/template/#example_Template
