package main

// Координация горутин с использованием каналов
// Напишите программу на языке программирования Go, использующую
// каналы для координации между горутинами.
// Программа должна имитировать работу системы по отправке и получению
// сообщений между различными компонентами.

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	ID   int
	Data string
}

func sender(messages chan<- Message, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		message := Message{ID: i, Data: fmt.Sprintf("Message %d", i)}
		messages <- message
		fmt.Printf("Sender sent message: ID=%d, Data=%s\n", message.ID, message.Data)
		time.Sleep(300 * time.Millisecond)
	}

	close(messages)
	done <- struct{}{}
}

func receiver(messages <-chan Message, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()

	for {
		select {
		case message, ok := <-messages:
			if !ok {
				fmt.Println("Receiver finished.")
				return
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Reciever received message: ID=%d, Data=%s\n", message.ID, message.Data)
		case <-done:
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	messages := make(chan Message, 3)
	done := make(chan struct{})

	wg.Add(1)
	go sender(messages, &wg, done)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go receiver(messages, &wg, done)
	}
	wg.Wait()
	close(done)
}
