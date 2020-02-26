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

// Create a new article with the title and content provided
func CreateNewArticle(title, content string) (*Article, error) {
	// Set the ID of a new article to one more than the number of articles
	a := Article{ID: len(articleList) + 1, Title: title, Content: content}

	// Add the article to the list of articles
	articleList = append(articleList, a)

	return &a, nil
}

// Update a article
func UpdateArticle(id int, title, content string) (*Article, error) {
	a, err := GetArticleByID(id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		a.Title = title
	}
	if content != "" {
		a.Content = content
	}

	return a, nil
}

// Delete a article
func DeleteArticle(id int) error {
	for k, a := range articleList {
		if a.ID == id {
			articleList = append(articleList[:k], articleList[k+1:]...)
			return nil
		}
	}
	return errors.New("Article not found")
}
