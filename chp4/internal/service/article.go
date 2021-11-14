// DTO -> DO

package service

import (
	"context"

	pb "github.com/jackycsl/geektime-go/chp4/api"
	"github.com/jackycsl/geektime-go/chp4/internal/biz"
)

// var ProviderSet = wire.NewSet(NewArticleService)

type ArticleService struct {
	pb.UnimplementedArticleServiceServer
	article *biz.ArticleUseCase
}

func NewArticleService(article *biz.ArticleUseCase) *ArticleService {
	return &ArticleService{
		article: article,
	}
}

func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	p, err := s.article.Get(req.Id)
	// this error is send to external. Can consider to change the error if don't want to expose the error.
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleReply{Article: &pb.Article{Id: p.ID, Title: p.Title, Content: p.Content}}, nil
}
