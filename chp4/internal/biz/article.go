// DO

package biz

import (
	"errors"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewArticleUseCase)

type Article struct {
	ID      int64
	Title   string
	Content string
}

type ArticleRepo interface {
	GetArticle(id int64) (*Article, error)
}

type ArticleUseCase struct {
	repo ArticleRepo
}

func NewArticleUseCase(repo ArticleRepo) *ArticleUseCase {
	return &ArticleUseCase{repo: repo}
}

func (uc *ArticleUseCase) Get(id int64) (p *Article, err error) {
	if id == 0 {
		return nil, errors.New("invalid article id")
	}
	p, err = uc.repo.GetArticle(id)
	// By right, the repo already handle SQLNoRowsError.
	// Here can consider to change the error to "article not found"
	if err != nil {
		return nil, err
	}
	return p, err
}
