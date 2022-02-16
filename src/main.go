package main

import (
	"fmt"

	"github.com/gammazero/workerpool"

	"github.com/David-Lor/go-example/internal/foo"
	"github.com/David-Lor/go-example/internal/foo/bar"
)

func main() {
	pool := workerpool.New(1)
	pool.SubmitWait(func() {
		fmt.Println(foo.Func())
		fmt.Println(bar.Func())
	})
}
