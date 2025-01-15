package downloader


import (
	"fmt"
	"sync"
	"time"
)

// DownloadManager manages multiple download tasks
type DownloadManager struct {
	Tasks     []*DownloadTask
	TaskQueue chan *DownloadTask
	wg        sync.WaitGroup
}

// NewDownloadManager creates a new download manager
func NewDownloadManager(bufferSize int) *DownloadManager {
	return &DownloadManager{
		Tasks:     []*DownloadTask{},
		TaskQueue: make(chan *DownloadTask, bufferSize),
	}
}

// AddTask adds a new download task
func (dm *DownloadManager) AddTask(task *DownloadTask) {
	dm.Tasks = append(dm.Tasks, task)
	dm.TaskQueue <- task
	fmt.Printf("Task added: %s\n", task.URL)
}

// Start processes the task queue with concurrency
func (dm *DownloadManager) Start(workers int) {
	for i := 0; i < workers; i++ {
		go dm.worker()
	}

	go func() {
		dm.wg.Wait()
		close(dm.TaskQueue)
	}()
}

// worker processes tasks from the queue
func (dm *DownloadManager) worker() {
	for task := range dm.TaskQueue {
		dm.wg.Add(1)
		err := DownloadFile(task)
		if err != nil {
			fmt.Printf("Error downloading %s: %v\n", task.URL, err)
			task.RetryCount++
			if task.RetryCount < 3 {
				fmt.Printf("Retrying %s...\n", task.URL)
				dm.TaskQueue <- task
			}
		}
		dm.wg.Done()
	}
}
