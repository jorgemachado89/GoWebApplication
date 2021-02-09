package main

import (
	"html/template"
	"os"
)

const tax = 6.75 / 100

type Product struct {
	Name  string
	Price float32
	SomeArray []string
}

const templateString = `
{{- "Item Information" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price }}
Price with Tax: {{ calcTax .Price | printf "$%.2f" }}
Length: {{ len .SomeArray }}
`

func main() {
	p := Product{
		Name:  "Lemonade",
		Price: 2.16,
		SomeArray: []string{"a", "b"},
	}
	fm := template.FuncMap{}
	fm["calcTax"] = func(price float32) float32 {
		return p.Price * (1 + tax)
	}
	t := template.Must(template.New("").Funcs(fm).Parse(templateString))
	t.Execute(os.Stdout, p)
}
