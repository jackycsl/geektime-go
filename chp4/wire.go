//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jackycsl/geektime-go/chp4/internal/biz"
	"github.com/jackycsl/geektime-go/chp4/internal/data"
	"github.com/jackycsl/geektime-go/chp4/internal/service"
)

func InitArticleService() *service.ArticleService {
	// wire.Build(service.NewArticleService, biz.NewArticleUseCase, data.NewArticleRepo, data.NewData)
	wire.Build(service.ProviderSet, biz.ProviderSet, data.ProviderSet)
	return &service.ArticleService{}
}
