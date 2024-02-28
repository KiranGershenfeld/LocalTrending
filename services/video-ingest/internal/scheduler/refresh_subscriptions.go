package scheduler

import (
	"context"
	"video-ingest/internal/channels"
	channelDB "video-ingest/internal/channels/database"
	"video-ingest/internal/config"
	"video-ingest/pkg/logging"
)

func RefreshAllSubscriptions(ctx context.Context, cfg *config.Config, db channelDB.ChannelDB) error {
	logger := logging.FromContext(ctx)
	logger.Info("Scheduler.RefreshAllSubscriptions running refresh subscriptions job")

	activeChannels, err := db.GetActiveSubscriptions(ctx)
	if err != nil {
		return err
	}
	for _, channel := range activeChannels {
		err := channels.SubscribeToChannel(cfg, channel.YoutubeID)
		if err != nil {
			logger.Errorw("Scheduler.RefreshAllSubscriptions subscription to channel", channel.YoutubeID, "err", err)
			return err
		}
	}
	return nil

}
