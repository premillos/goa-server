package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"com.goa/internal/config"
)

func Bootstrap(ctx context.Context, cfg string) {

	config.MustLoad(cfg)

	fmt.Printf("%+v", config.C)
	clearFunc1, _ := InitDb(ctx)
	Run(ctx, func(ctx context.Context) (func(), error) {
		_, clearFunc2, _ := InitEngine(ctx)
		return func() {
			clearFunc1()
			clearFunc2()
		}, nil
	})

}

// The Run function sets up a signal handler and executes a handler function until a termination signal
// is received.
func Run(ctx context.Context, handler func(ctx context.Context) (func(), error)) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFn, err := handler(ctx)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		fmt.Printf("接收到信号 %s", sig.String())

		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFn()

	time.Sleep(time.Millisecond * 100)
	os.Exit(state)
	return nil
}
