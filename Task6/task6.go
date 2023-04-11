package main

import (
	"context"
)

func main() {
	//1:Каналом
	stopChanel := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopChanel:
				return
			}
		}
	}()
	stopChanel <- struct{}{}
	//2:Контекстом
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	cancel()
}
