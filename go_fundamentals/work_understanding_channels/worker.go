package work_understanding_channels

type WG struct {
	main chan func()
	allDone chan bool
}

func New(n int) WG {
	res := WG {
		main : make(chan func()),
		allDone: make(chan bool),
	}

	procDone := make(chan bool)
	for i := 0; i < n; i++ { // run as many workers were required
		go func(){
			for {
				f := <-res.main
				if f == nil {
					procDone <- true
					return
				}
				f()
			}
		}()
	}
	go func(){
		for i := 0; i < n; i++ { //this also control the rate at which this completes, weather it takes only that much time it should or more
			_ = <-procDone
		}
		res.allDone <- true
	}()
	return res
}

func New_i(n int) WG {
	res := WG {
		main : make(chan func()),
		allDone: make(chan bool),
	}

	procDone := make(chan bool)
	for i := 0; i < n; i++ { // run as many workers were required
		go func(){
			for f := range res.main {
				f()
			}
		}()
	}
	go func(){
		for i := 0; i < n; i++ { //this also control the rate at which this completes, weather it takes only that much time it should or more
			_ = <-procDone
		}
		res.allDone <- true
	}()
	return res
}


func (wg WG) Add(f func()) {
	wg.main <- f
}

func (wg WG) Wait() {
	close(wg.main)
	_ = <- wg.allDone
}
