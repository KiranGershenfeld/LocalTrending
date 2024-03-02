CREATE TABLE IF NOT EXISTS video_views_7d (
        video_id VARCHAR(255),
        channel_id VARCHAR(255),
        upload_time TIMESTAMP,
        view_count INT
);

INSERT INTO videos (youtube_id, channel_id, title, thumbnail_url, upload_time, created_at, updated_at, deleted_at)
VALUES 
    ('t5S9LfvBnOg', 'channel123', 'Video Title 1', 'http://example.com/thumbnail1.jpg', '2024-02-21 00:00:00', NOW(), NOW(), NULL),
    ('FWdv_DGiD5Q', 'channel456', 'Video Title 2', 'http://example.com/thumbnail2.jpg', '2024-02-21 00:05:00', NOW(), NOW(), NULL),
    ('yD0Kfup7U3U', 'channel789', 'Video Title 3', 'http://example.com/thumbnail3.jpg', '2024-02-21 00:07:00', NOW(), NOW(), NULL);