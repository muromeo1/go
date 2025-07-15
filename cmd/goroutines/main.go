package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	id    int
	name  string
	email string
}

const count = 1_000

var (
	contentFast sync.Map
	contentSlow sync.Map
)

func main() {
	start := time.Now()

	fmt.Println("Started...")

	go fast(start)
	slow(start)
}

func fast(start time.Time) {
	ch := make(chan User, count)
	for i := range count {
		go createUser(i, ch, "fast")
	}

	for range count {
		<-ch
	}

	fmt.Println("Goroutine took", time.Since(start))
	fmt.Println("Goroutine count", Count(&contentFast))
}

func slow(start time.Time) {
	for i := range count {
		createUser(i, nil, "slow")
	}

	fmt.Println("Normal took", time.Since(start))
	fmt.Println("Normal count", Count(&contentSlow))
}

func createUser(i int, ch chan User, kind string) {
	user := User{
		id:    i,
		name:  fmt.Sprintf("name %d", i),
		email: fmt.Sprintf("email%d@email.com", i),
	}

	switch kind {
	case "fast":
		contentFast.Store(i, user)
	case "slow":
		contentSlow.Store(i, user)
	}

	time.Sleep(time.Millisecond * 200)

	if ch != nil {
		ch <- user
	}
}

func Count(c *sync.Map) int {
	var i int

	c.Range(func(k, v any) bool {
		i++
		return true
	})

	return i
}
