package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var country = "us"

const (
	urlPrefix string = "https://newsapi.org/v2/top-headlines?country="
	password  string = "9621ea2f064b4caebb083ccaafff7fa6"
	apikeyfix string = "&apiKey="
)

type source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type articles struct {
	Source      source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

//NewsListCall blah blah
type NewsListCall struct {
	//yolo
	Status       string     `json:"status"`
	TotalResults int        `json:"totalResults"`
	Articles     []articles `json:"articles"`
}

func makeAPICall(prefix string, pass string, countryCode string) NewsListCall {
	req, err := http.NewRequest("GET", "https://newsapi.org/v2/top-headlines?country=us", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("X-Api-Key", "9621ea2f064b4caebb083ccaafff7fa6")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	newsListCall := NewsListCall{}
	jsonError := json.Unmarshal(body, &newsListCall)

	if err != nil {
		log.Println(jsonError)
	}

	return newsListCall
}

type Article struct {
	Source string
	Author string
	Title  string
	URL    string
}

type NewsPage struct {
	Header   string
	News     string
	Articles []Article
}

var Arts []Article

func newsDisplayHandler(w http.ResponseWriter, r *http.Request) {
	newsPage := NewsPage{
		Header:   "This is your Mostly Fake news update for " + time.Now().Format("Mon 2006-01-2"),
		News:     "Here's your daily dose of mostly Fake News",
		Articles: Arts}
	parse, _ := template.ParseFiles("template/newsPage.html")
	parse.Execute(w, newsPage)
}

func main() {
	//Get News JSON
	news := makeAPICall(urlPrefix, password, country)

	//Get News Article from JSON and format HTML presentation
	for _, ne := range news.Articles {
		article := Article{
			Source: ne.Source.Name,
			Author: ne.Author,
			Title:  ne.Title,
			URL:    ne.URL}

		Arts = append(Arts, article)
	}

	//	fmt.Println(news.TotalResults)
	fmt.Println(len(Arts))
	http.HandleFunc("/news/", newsDisplayHandler)
	http.ListenAndServe(":9003", nil)

}
