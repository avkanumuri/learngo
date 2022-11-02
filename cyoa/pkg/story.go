package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

var defaultTmpl = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Chose your own adventure</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraph}}
        	<p>{{.}}</p>
        {{end}}
        <ul>
        {{range .Options}}
            	<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>`

func (s Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultTmpl))
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Unexpected error", 500)
		}
		return
	}

}

// func DispPage(s Story, fallback http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		tpl := template.Must(template.New("").Parse(defaultTmpl))
// 		err := tpl.Execute(w, s["intro"])
// 		if err != nil {
// 			panic(err)
// 		}
// 		fallback.ServeHTTP(w, r)
// 	}

// }

type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title,omitempty"`
	Paragraph []string `json:"story,omitempty"`
	Options   []Option `json:"options,omitempty"`
}

type Option struct {
	Text    string `json:"text,omitempty"`
	Chapter string `json:"arc,omitempty"`
}

func JsonStory(r io.Reader) (Story, error) {
	var book Story
	d := json.NewDecoder(r)
	if err := d.Decode(&book); err != nil {
		return nil, err
	}
	return book, nil

}
