package work_understanding_channels

import (
	"testing"
	"time"
	"sync"
)

func Test_Waiting(t *testing.T){
	//wg := New(50)
	wg := New_i(50)
	p := 0
	var l sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(func(){
			time.Sleep(time.Second / 10)
			l.Lock()
			p++
			l.Unlock()
		})
	}

	wg.Wait()
	if p != 1000 {
		t.Log("p does not equal 1000, instead equals ", p)
		t.Fail()
	}
}
