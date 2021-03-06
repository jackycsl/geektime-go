// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	g, ctx := errgroup.WithContext(context.Background())

	// Listen for signal terminate
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

	// start server
	g.Go(func() error {
		fmt.Println("Starting http server...")

		return srv.ListenAndServe()
	})

	// stop server
	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		fmt.Println("Stopping http server...")
		err := srv.Shutdown(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}
