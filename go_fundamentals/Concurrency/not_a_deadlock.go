package main

func main() {
	done := make(chan bool)
	close(done)

	// Will not block and will print false twice
	// because it’s the default value for bool type
	println(<-done)
	println(<-done)
}