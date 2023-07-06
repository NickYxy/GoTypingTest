package GoTypingTest

import (
	"fmt"
	"sync"
)

type packet struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Flight struct {
	mu     sync.Mutex
	flight map[string]*packet
}

func (f *Flight) Fly(key string, fn func() (interface{}, error)) (interface{}, error) {
	f.mu.Lock()
	if f.flight == nil {
		f.flight = make(map[string]*packet)
	}

	if p, ok := f.flight[key]; ok {
		f.mu.Unlock()
		p.wg.Wait()
		return p.val, p.err
	}

	p := new(packet)
	p.wg.Add(1)
	f.flight[key] = p
	f.mu.Unlock()

	p.val, p.err = fn()
	p.wg.Done()

	f.mu.Lock()
	delete(f.flight, key)
	f.mu.Unlock()

	return p.val, p.err
}

// GetPeer cache consistency 一致性哈希
//func (c *Consistence) GetPeer(key string) string {
//	if len(c.ring) == 0 {
//		return ""
//	}
//	hashValue := int(c.hash([]byte(key)))
//	idx := sort.Search(len(c.ring), func(i int) bool {
//		return c.ring[i] >= hashValue
//	})
//
//	return c.hashMap[c.ring[idx%len(c.ring)]]
//}

func main() {
	fmt.Println(any("aaa"))
}
