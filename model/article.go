package model

import (
	"GinHello/initDB"
	"log"
)

type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (article Article) Insert() int {
	result, e := initDB.Db.Exec("insert into ` article` (type, content) values (?, ?);", article.Type, article.Content)
	if e != nil {
		log.Panicln("文章添加失败", e.Error())
	}
	i, _ := result.LastInsertId()
	return int(i)
}

func (article Article) FindAll() []Article {
	rows, e := initDB.Db.Query("select * from ` article`;")
	if e != nil {
		log.Panicln("查询数据失败")
	}

	var articles []Article

	for rows.Next() {
		var a Article
		if e := rows.Scan(&a.Id, &a.Type, &a.Content); e == nil {
			articles = append(articles, a)
		}
	}
	return articles
}

func (article Article) FindById() Article {
	row := initDB.Db.QueryRow("select * from ` article` where id=?;", article.Id)
	if e := row.Scan(&article.Id, &article.Type, &article.Content); e != nil {
		log.Panicln("绑定发生错误", e.Error())
	}
	return article
}

func (article Article) DeleteOne() {
	if _, e := initDB.Db.Exec("delete from ` article` where id = ?", article.Id); e != nil {
		log.Panicln("数据发生错误，无法删除")
	}
}
