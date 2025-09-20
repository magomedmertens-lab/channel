cat > main.go <<'EOF'
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Измените N по необходимости
	N := 5 * time.Second

	// Отправитель
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 1
		for {
			select {
			case <-done:
				fmt.Println("sender: stopped")
				return
			case ch <- i:
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Приёмник
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("receiver: stopped")
				return
			case val := <-ch:
				fmt.Println("Read:", val)
			}
		}
	}()

	// Ждём N секунд и даём сигнал на завершение
	<-time.After(N)
	close(done)

	// Ждём, пока горутины корректно завершатся
	wg.Wait()
	fmt.Println("Время вышло, программа завершена.")
}
EOF
