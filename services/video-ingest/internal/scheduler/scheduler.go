package scheduler

import (
	"context"
	channelDB "video-ingest/internal/channels/database"
	"video-ingest/internal/config"
	"video-ingest/pkg/logging"

	"github.com/go-co-op/gocron/v2"
)

func Run(cfg *config.Config, channelDB channelDB.ChannelDB) {
	ctx := context.WithValue(context.Background(), "caller", "scheduler")
	logger := logging.FromContext(ctx)

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		// handle error
	}

	logger.Info("Scheduler.Run scheduler initialized")

	// Refresh subscription job
	_, err = s.NewJob(
		gocron.CronJob(cfg.Scheduler.RefreshAllSubscriptionsCronSchedule, false),
		gocron.NewTask(
			RefreshAllSubscriptions,
			context.Background(),
			cfg,
			channelDB,
		),
	)
	if err != nil {
		// handle error
	}
	// start the scheduler
	s.Start()
	return
}
