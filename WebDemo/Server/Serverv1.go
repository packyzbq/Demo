// 第一版server，简单调用
package Server

import (
	"WebDemo/Data"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func viewHandlerHardcode(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view"):]
	p, _ := Data.LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandlerHardcode(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := Data.LoadPage(title)
	if err != nil {
		p = &Data.Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	log.Printf("trigger edit, title = %s", title)
	p, err := Data.LoadPage(title)
	if err != nil {
		p = &Data.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	log.Printf("trigger view, title = %s", title)
	p, err := Data.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Data.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("trigger index")
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Data.Page) {
	// 每次调用 ParseFile 效率低，用 tmpl.Must 将所有模板 parse，并在需要时进行 render
	//t, err := tmpl.ParseFiles("tmpl" + tmpl + ".html")
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := getTitle(r)
		if err != nil {
			http.NotFound(w, r)
		}
		fn(w, r, m)
	}
}

func StartServer() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
