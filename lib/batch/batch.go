package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	status := make(chan struct{}, pool)
	var mx sync.Mutex
	var wg sync.WaitGroup
	var i int64
	for i = 0; i < n; i++ {
		status <- struct{}{}
		wg.Add(1)
		go func(next int64) {
			nextUser := getOne(next)
			mx.Lock()
			res = append(res, nextUser)
			mx.Unlock()
			wg.Done()
			<-status
		}(i)
	}
	wg.Wait()
	return res
}
