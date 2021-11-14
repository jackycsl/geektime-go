// DO -> PO
package data

import "github.com/jackycsl/geektime-go/chp4/internal/biz"

type ArticleRepo struct {
	data *Data
}

func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &ArticleRepo{
		data: data,
	}
}

func (ar *ArticleRepo) GetArticle(id int64) (*biz.Article, error) {
	// ENT orm to obtain data from database
	// p, err = ar.data.db.Article.Get(id)
	// if err != nil {
	// 	return nil, err
	// }
	// return &biz.Article{
	// 	ID:      p.ID,
	// 	Title:   p.Title,
	// 	Content: p.Content,
	// }, nil
	return &biz.Article{
		ID:      1,
		Title:   "Apple",
		Content: "It's red",
	}, nil
}
