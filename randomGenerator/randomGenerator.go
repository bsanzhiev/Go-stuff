package main

// Написать генератор случайных чисел
// В принципе, легкая задача, на базовые знания по асинхронному взаимодействию в Go. 
// Для решения я бы использовал небуфферезированный канал. 
// Будем асинхронно писать туда случайные числа и закроем его, когда закончим писать.

import (
  "fmt"
  "math/rand"
  "time"
)

func randNumGenerator(n int) <-chan int {
  // генерируем случайные числа
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  // создаем канал
  out := make(chan int)
  // запускаем горутину 
  go func() {
    for i := 0; i < n; i++ {
      out <- r.Intn(n)
      }
      close(out)
  }()
  return out  
}

func main() {
  for num := range randNumGenerator(10) {
    fmt.Println(num)
  }
}