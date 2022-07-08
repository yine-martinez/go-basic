package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

//ch:= make(chan int)
//(ch <-chan<- int)
func main() {

	//chanel
	/*wgc := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wgc.Add(2)
	go func(ch <-chan int, wgc *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wgc.Done()
	}(ch, wgc)
	go func(ch chan<- int, wgc *sync.WaitGroup) {
		//close(ch)
		ch <- 42
		wgc.Done()
	}(ch, wgc)

	wgc.Wait()*/
	//chanel
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id, m); ok {
				/*fmt.Println("from cache")
				fmt.Println(b)*/
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id, m); ok {
				//fmt.Println("from database")
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
				//fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from Cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}

	return Book{}, false
}
