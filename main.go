package main

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
		//fmt.Println("Generate a ", i)
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		//print("filter ", i, " prime ", prime, "\n")
		if i%prime != 0 {
			out <- i
			//print("out ", i, " prime ", prime, "\n")
		}
	}
}

func main() {
	ch := make(chan int)
	go Generate(ch)
	for {
		prime := <-ch
		if prime > 100 {
			break
		}
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1

	}
}
