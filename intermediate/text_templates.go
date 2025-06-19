package main

import (
	"html/template"
	"os"
)

func main() {
	// Define a simple template with a placeholder
	tmpl := `Hello, {{.Name}}! Welcome to the Go programming world.`

	// Create a new template and parse the letter into it
	t, err := template.New("greeting").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Define data to be passed to the template
	data := struct {
		Name string
	}{
		Name: "Alice",
	}

	// Execute the template and print the result
	err = t.ExecuteTemplate(os.Stdout, "greeting", data)
	if err != nil {
		panic(err)
	}
}
