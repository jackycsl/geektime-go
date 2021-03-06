// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/jackycsl/geektime-go/chp4/internal/biz"
	"github.com/jackycsl/geektime-go/chp4/internal/data"
	"github.com/jackycsl/geektime-go/chp4/internal/service"
)

// Injectors from wire.go:

func InitArticleService() *service.ArticleService {
	dataData := data.NewData()
	articleRepo := data.NewArticleRepo(dataData)
	articleUseCase := biz.NewArticleUseCase(articleRepo)
	articleService := service.NewArticleService(articleUseCase)
	return articleService
}
