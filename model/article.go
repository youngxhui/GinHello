package model

import (
	"GinHello/initDB"
)

type Article struct {
	Id      int    `json:"id" example:"1"`
	Type    string `json:"type" example:"js"`
	Content string `json:"content" example:"hello js"`
}

func (article Article) TableName() string {
	return "article"
}

func (article Article) Insert() int {
	create := initDB.Db.Create(&article)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (article Article) FindAll() []Article {
	var articles []Article
	initDB.Db.Find(&articles)
	return articles
}

func (article Article) FindById() Article {
	initDB.Db.First(&article, article.Id)
	return article
}

func (article Article) DeleteOne() {
	initDB.Db.Delete(article)
}
