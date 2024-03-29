package database

import (
	"context"
	"video-ingest/internal/channels/model"
	"video-ingest/internal/database"
	"video-ingest/pkg/logging"

	"gorm.io/gorm"
)

type ChannelDB interface {
	// SaveArticle saves a given article with tags.
	// if not exist tags, then save a new tag
	AddChannel(ctx context.Context, channel *model.Channel) error
	GetActiveSubscriptions(ctx context.Context) (activeChannels []model.Channel, err error)
}

// NewArticleDB creates a new article db with given db
func NewChannelDB(db *gorm.DB) ChannelDB {
	return &channelDB{db: db}

}

type channelDB struct {
	db *gorm.DB
}

func (c *channelDB) AddChannel(ctx context.Context, channel *model.Channel) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, c.db)

	if err := db.WithContext(ctx).Create(channel).Error; err != nil {
		if database.IsKeyConflictErr(err) {
			return database.ErrKeyConflict
		}
		logger.Errorw("channel.db.AddChannel failed to save channel", "err", err)

		return err
	}
	return nil
}

func (c *channelDB) GetActiveSubscriptions(ctx context.Context) (activeChannels []model.Channel, err error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, c.db)

	if err := db.WithContext(ctx).Where("ingestion_active = ?", true).Find(&activeChannels).Error; err != nil {
		logger.Errorw("channel.db.GetActiveSubscriptions failed to get active channels", "err", err)
		return nil, err
	}

	return activeChannels, nil
}
