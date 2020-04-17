package test_test

import (
	"fmt"
	"speed/app/lib/workerPool"
	"testing"
)

func TestPool(t *testing.T) {

	p := workerPool.NewPool(5)
	ts := &workerPool.Task{}
	ts.SetHandel(func() error {
		fmt.Println("打酱油任务..........")
		return nil
	})

	//p.ExternalChain <- ts

	go func() {
		for i := 0; i < 50; i++ {
			p.ExternalChain <- ts
		}
	}()

	go p.Run()
	select {

	}

}
