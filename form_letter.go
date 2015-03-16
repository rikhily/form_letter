package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
	Dear {{.Honorific}}{{.LastName}},
	{{if .Attended}}
	It was good to see you at the fund raiser.{{else}}
	I am sorry that you could not make it to the fund raiser, we were expecting you.{{end}}
	{{if .Donated}}
	Thanks a lot for your donation.
	Remainder for upcoming events:{{else}}
	Remainder for upcoming events:{{end}}
	{{range .Events}}{{.}}
	{{end}}
	Best wishes,
	Ricky
	`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Honorific, LastName, string
		Attended, Donated bool
		Events []string
	}

	var recipients = []Recipient{
		{"Mr.", "Goku", true, true, []string{"Date:Tuesday, March 17 Yoga on Stage at Historic Warnors Theatre - 5:45 pm", "Date:Tuesday, March 17 St Paddy's Day w/ Celtic Alchemy at Peeve's Public House - 7:30pm", "Date:Tuesday, March 17 Alton Brown: Edible Inevitable Tour - 7:30pm", "Date:Wednesday, March 18 San Joaquin Valley Town Hall: Man's Inhumanity to Man", "Date:Thursday, March 19 Government Contracting Workshop - 9:00am"}},
		{"Mrs.", "Chi-Chi",  true, false, []string{"Date:Tuesday, March 17 St Paddy's Day w/ Celtic Alchemy at Peeve's Public House - 7:30pm", "Date:Tuesday, March 17 Alton Brown: Edible Inevitable Tour - 7:30pm"}},
		{"Ms", "Bulma", false, false, []string{"Date:Tuesday, March 17 St Paddy's Day w/ Celtic Alchemy at Peeve's Public House - 7:30pm"}},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
