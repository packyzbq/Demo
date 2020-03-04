package Data

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

const (
	POST_PATH = "WikiPost/"
)

func (p *Page) save() error {
	filename := POST_PATH + p.Title + ".txt"
	// 写入文件
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func (p *Page) Save() error {
	return p.save()
}

func LoadPage(title string) (*Page, error) {
	filename := POST_PATH + title + ".txt"
	body, error := ioutil.ReadFile(filename)
	if error != nil {
		return nil, error
	}
	return &Page{Title: title, Body: body}, nil
}
