package main

import (
    "html/template"
    "net/http"
)

type ColorList struct {
    ListTitle string
    Colors    []Color
}

type Color struct {
    Name string
    Rgb  string
}

func main() {
    index := template.Must(template.ParseFiles("index.html"))

    data := ColorList{
        ListTitle: "Colors",
        Colors: []Color{
            {Name: "Red", Rgb: "rgb(255,0,0)"},
            {Name: "Green", Rgb: "rgb(0,128,0)"},
            {Name: "Blue", Rgb: "rgb(0,0,255)"},
        },
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        index.Execute(w, data)
    })

    fs := http.FileServer(http.Dir("assets/"))
    http.Handle("/assets/", http.StripPrefix("/assets/", fs))

    http.ListenAndServe(":3000", nil)
}
