package jobs

import (
	"fmt"
	"job-scheduler/internal/database"
	"time"
)

type GetVideoViews struct {
	destinationTable       string
	timeSinceUploadSeconds int
	jobRunIntervalSeconds  int
	db                     database.Database
}

func (job *GetVideoViews) Run() error {
	//Get videos in time range
	db, err := job.db.Connect()
	if err != nil {
		return err
	}

	// Get the current time
	currentTime := time.Now()

	// Calculate the start time for the query
	startTime := currentTime.Add(-time.Duration(job.timeSinceUploadSeconds) * time.Second)

	// Calculate the end time for the query
	endTime := currentTime.Add(-time.Duration(job.timeSinceUploadSeconds+job.jobRunIntervalSeconds) * time.Second)

	rows, err := job.db.Query(db, fmt.Sprintf(`
		SELECT youtube_id, channel_id, upload_time FROM videos
		WHERE upload_time <= '%v' AND
		upload_time >= '%v';
	`, startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")))

	if err != nil {
		return err
	}
	//Look up video view counts
	//Save to designated table
}

func (job *GetVideoViews) Fail() error {

}
