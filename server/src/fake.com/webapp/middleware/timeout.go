package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel() // releases resources if operation completes before timeout elapses

	r.WithContext(ctx)

	ch := make(chan struct{})
	go func() {
		fmt.Printf("%s\tErrors before serving next: %v.\n", time.Now(), ctx.Err())
		if ctx.Err() == nil {
			fmt.Printf("%s:\tServing new request.\n", time.Now())
			tm.Next.ServeHTTP(w, r)
			ch <- struct{}{}
		}
	}()

	select {
	case <-ch:
		fmt.Printf("%s:\tContext not finished: %v \n", time.Now(), ctx.Err())
		return
	case <-ctx.Done():
		//fmt.Printf("%s: Context is finished: %v \n", time.Now(), ctx.Err())
		w.WriteHeader(http.StatusRequestTimeout)
		fmt.Printf("%s:\tContext is finished: %v \n", time.Now(), ctx.Err())
	}
}
