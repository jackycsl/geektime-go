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
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterArticleServiceServer(srv, InitArticleService())

	g, ctx := errgroup.WithContext(context.Background())

	// Listen for terminate signal
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

	// Start server
	g.Go(func() error {
		log.Printf("server listening at %v", lis.Addr())

		return srv.Serve(lis)
	})

	// Stop server
	g.Go(func() error {
		<-ctx.Done()

		fmt.Println("Stopping http server...")
		srv.GracefulStop()

		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}
