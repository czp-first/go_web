package main

import (
	"os"
	"text/template"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的pipeline if demo: {{if `anything`}} 我有内涵\n {{end}}"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("templat test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if ``}} if部分\n {{else}} else部分\n {{end}}"))
	tIfElse.Execute(os.Stdout, nil)
}
