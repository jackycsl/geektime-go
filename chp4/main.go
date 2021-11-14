package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/jackycsl/geektime-go/chp4/api"
	"github.com/jackycsl/geektime-go/chp4/internal/biz"
	"github.com/jackycsl/geektime-go/chp4/internal/data"
	"github.com/jackycsl/geektime-go/chp4/internal/service"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterArticleServiceServer(s, InitArticleService())

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case s := <-c:
			return fmt.Errorf("caught signal: %v", s.String())
		}

	})

	g.Go(func() error {
		log.Printf("server listening at %v", lis.Addr())

		go func() {
			<-ctx.Done()

			fmt.Println("Stopping grpc server...")
			s.GracefulStop()
		}()
		return s.Serve(lis)
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}

// wire generate 时遇到 error, 暂时用手写， 请问老师可以帮忙看为什么会有error
func InitArticleService() *service.ArticleService {
	db := data.NewData()
	articleRepo := data.NewArticleRepo(db)
	articleBiz := biz.NewArticleUseCase(articleRepo)
	articleService := service.NewArticleService(articleBiz)
	return articleService
}
