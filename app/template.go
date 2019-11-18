package main

import (
	"html/template"
	"path"
	"runtime"
)

type Alert struct {
	Message string
	Theme   string
}

func LoadTemplate(filename string) *template.Template {
	_, callerFileName, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller information")
	}

	dir := path.Dir(callerFileName)

	layout := dir + "/views/layout.html"
	view := dir + "/views/" + filename

	return template.Must(template.ParseFiles(layout, view))
}
