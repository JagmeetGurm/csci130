
package main


import (
	
"log"
	
"os"
"text/template"

)


func main() {
	// Define a template.

	const letter = `

Dear {{.Honorific}}{{.Name}},

{{if .Attended}}

It was great that you attended fundraiser.{{else}}

It is sad that you couldn't attend the fundraiser.{{end}}

{{if.Donated}}

Thank you for the donation.

Reminder:{{else}}

Reminder:{{end}}

{{range .Events}}{{.}}

{{end}}

Best wishes,

Jagmeet

`








// Prepare some data to insert into the template.
	
type Recipient struct {
		
Honorific, Name string
	
	Attended, Donated bool   
Events []string

	

}

	var recipients = []Recipient{
	
	{"Ms."," Mildred", true, true, []string{"Blood donation camp on 3/18/15", "Save-Water on 3/20/15", "Pet-care on 3/22/15"}},
	
	{"Mr." ,"Johnson", true, false,[]string{"Blood donation camp on 3/18/15", "Save-Water on 3/20/15", "Pet-care on 3/22/15"}},
		
{"Mr."," Rodney", false, false,[]string{"Blood donation camp on 3/18/15", "Save-Water on 3/20/15", "Pet-care on 3/22/15"}},

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

