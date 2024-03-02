package scheduler

import (
	"context"
	"job-scheduler/internal/config"
	"job-scheduler/internal/jobs"
	"job-scheduler/internal/utils"
	"job-scheduler/pkg/logging"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func Run(cfg *config.Config) {
	ctx := context.WithValue(context.Background(), "caller", "scheduler")
	logger := logging.FromContext(ctx)

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		logger.Errorw("Could not create gocron scheduler", "err", err)
	}

	logger.Info("Scheduler.Run scheduler initialized")

	sevenDayVideoViews := jobs.PythonJob{
		Dir:  "jobs/delayed_video_views",
		Args: []string{"2024-02-28 00:09:00", utils.CreatePostgresConnectionString(cfg), "AIzaSyCLntcNQGAdjFbLdOJHMQDCrWze9QCRXEg"},
	}

	// Refresh subscription job
	_, err = s.NewJob(
		gocron.CronJob(cfg.SchedulerConfig.SevenDayVideoViewsCronSchedule, false),
		gocron.NewTask(
			sevenDayVideoViews.Execute,
			context.Background(),
		),
	)
	if err != nil {
		logger.Errorw("Could not run job", "err", err)
	}
	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(30 * time.Minute):
	}
	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
	}
	return
}
