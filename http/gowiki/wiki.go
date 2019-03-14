package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// the Body field is a []byte rather than string because it is the type expected by the io libraries is used
type Page struct {
	Title string
	Body  []byte
}

// the Page struct describes how page data will be stored in memory
// for persistent storage we create a save method on Page
// this method will save the Page's Body to a text file
// it return an error value because that is the return type of WriteFile
// if all goes well, Page.save() will return nil (the zero-value for pointers, interfaces, and some other types)
// the octal integer literal 0600 indicates that the file should be created with read-write permissions for the current user only (see the Unix man page open(2) for details)
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// in addition to saving pages, we will want to load pages too
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):] // handle URLs prefixed with "/view/"
	p, err := loadPage(title)
	if err != nil {
		// The http.Redirect function adds an HTTP status code of http.StatusFound (302) and a Location header to the HTTP response.
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<textarea name=\"body\">%s</textarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)

	// The function template.ParseFiles will read the contents of edit.html and return a *template.Template.

	// The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter.
	// The .Title and .Body dotted identifiers refer to p.Title and p.Body.
	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)

	renderTemplate(w, "edit", p)
}

// The page title (provided in the URL) and the form's only field, Body, are stored in a new Page.
// The save() method is then called to write the data to a file, and the client is redirected to the /view/ page.
// The value returned by FormValue is of type string.
// We must convert that value to []byte before it will fit into the Page struct.
//We use []byte(body) to perform the conversion.
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
