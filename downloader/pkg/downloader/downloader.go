package downloader


import (
	"time"
)

// DownloadTask represents a single download task
type DownloadTask struct {
	ID          int
	URL         string
	Destination string
	Status      string // "pending", "in-progress", "completed", "failed"
	RetryCount  int
	CreatedAt   time.Time
}

// NewDownloadTask creates a new download task
func NewDownloadTask(id int, url, destination string) *DownloadTask {
	return &DownloadTask{
		ID:          id,
		URL:         url,
		Destination: destination,
		Status:      "pending",
		RetryCount:  0,
		CreatedAt:   time.Now(),
	}
}