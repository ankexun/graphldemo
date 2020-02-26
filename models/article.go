package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
var articleList = []Article{
	Article{ID: 1, Title: "Article 1,文章1", Content: "Article 1 body,中国加油"},
	Article{ID: 2, Title: "Article 2,文章2", Content: "Article 2 body,武汉加油"},
}

// Return a list of all the articles
func GetAllArticles() []Article {
	return articleList
}

// Get article
func GetArticleByID(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
