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

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		return fmt.Errorf("caught signal: %v", s.String())
	})

	g.Go(func() error {
		fmt.Println("Starting http server...")

		go func() {
			<-ctx.Done()

			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			fmt.Println("Stopping http server...")
			err := srv.Shutdown(ctx)
			if err != nil {
				return
			}
		}()
		return srv.ListenAndServe()
	})

	err := g.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}
