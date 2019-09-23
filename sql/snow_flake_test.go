package sql

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	worker, err := NewNode(1)
	if err != nil {
		return
	}

	ch := make(chan int64)
	count := 100000
	for i := 0; i < count; i++ {
		go func() {
			id := worker.Generate()
			ch <- id
		}()
	}

	// defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		_, ok := m[id]
		if ok {
			t.Error("ID is not unique!")
			return
		}
		m[id] = i
	}
	fmt.Println("All", count, " successed!")
}

func TestServicesGenIds(t *testing.T) {
	ids, err := GenerateServiceIdentities(ServiceAPIPlatForm, 100)
	fmt.Print(ids)
	fmt.Print(err)
}