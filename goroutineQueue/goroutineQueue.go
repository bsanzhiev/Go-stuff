package main

// Очередь синхронизированных задач
// Напишите программу на языке программирования Go, реализующую простую очередь задач. 
// Необходимо создать горутины, которые будут добавлять задачи в очередь, 
// и другие горутины, которые будут извлекать и выполнять эти задачи.
// Обеспечьте синхронизацию доступа к очереди так, чтобы избежать гонок данных.

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data string
}

type TaskQueue struct {
	tasks []Task
	mutex sync.Mutex
}

func (q *TaskQueue) enqueue(task Task) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.tasks = append(q.tasks, task)
	fmt.Printf("Task enqueued: ID = %d, Data = %s\n", task.ID, task.Data)
}

func (q *TaskQueue) dequeue() (Task, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.tasks) == 0 {
		return Task{}, false
	}

	task := q.tasks[0]
  // Удаляем первый элемент из очереди
	q.tasks = q.tasks[1:]
	fmt.Printf("Task dequeued: ID = %d, Data = %s\n", task.ID, task.Data)
	return task, true
}

func processTask(workerID int, queue *TaskQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := queue.dequeue()
		if !ok {
			fmt.Printf("Worker %d finished.\n", workerID)
			return
		}

		time.Sleep(500 * time.Millisecond)

		fmt.Printf("Worker %d processed task: ID=%d, Data=%s\n", workerID, task.ID, task.Data)
	}
}

func main() {
	var wg sync.WaitGroup
	taskQueue := TaskQueue{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < 5; j++ {
				task := Task{ID: j, Data: fmt.Sprintf("Task %d", j)}
				taskQueue.enqueue(task)
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go processTask(i, &taskQueue, &wg)
	}
	wg.Wait()
}
