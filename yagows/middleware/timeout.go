package middleware

import (
	"context"
	"time"
	. "yagows"
)

func Timeout(duration int) RequestHandler {
	return func(c *Context) {
		finishChan := make(chan struct{})
		panicChan := make(chan interface{})

		timeoutCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(duration)*time.Millisecond)
		defer cancel()

		go func() {
			defer func() {
				if err := recover(); err != nil {
					panicChan <- err
				}
			}()

			c.Next()

			close(finishChan)
		}()

		select {
		case err := <-panicChan:
			panic(err)
		case <-finishChan:
			return
		case <-timeoutCtx.Done():
			panic("timeout")
		}
	}
}
