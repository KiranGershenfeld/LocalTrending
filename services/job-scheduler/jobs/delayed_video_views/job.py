#!/usr/bin/env python3
from datetime import datetime, timedelta
import sys
import logging
from sqlalchemy import create_engine, text, exc
from jobs_common import youtube 

def get_videos_from_range(engine, start_date, end_date):
        # Define the query to select records from the "videos" table created within the past hour
    stmt = text(f"""
        SELECT youtube_id, channel_id, upload_time FROM videos
        WHERE upload_time <= '{start_date}' AND
        upload_time >= '{end_date}';
    """)
    print(stmt)
    with engine.connect() as connection:
        records = connection.execute(stmt)
        connection.commit()
    return records

def get_batch_video_views(video_records, api_client):
    logger = logging.getLogger(__name__)

    records_with_views = []
    for record in video_records:
        video_id = record[0]
        try:
            views = api_client.get_video_views(video_id, method='api')
        except Exception as e: 
            logger.error(f"Could not get video views for video {video_id}. Error: {e}")
            continue

        record_with_views = tuple(record) + (views,)
        records_with_views.append(record_with_views)

    return records_with_views

def update_video_views_table(engine, records_with_views, table_name):
    stmt = f"""INSERT INTO {table_name} (video_id, channel_id, upload_time, view_count)
              VALUES (:video_id, :channel_id, :upload_time, :view_count)
           """
    with engine.connect() as connection:
        for record_with_views in records_with_views:
            video_id, channel_id, upload_time, view_count = record_with_views
            connection.execute(text(stmt), {"video_id": video_id, "channel_id": channel_id, "upload_time": upload_time, "view_count": view_count})
        connection.commit()

def create_video_views_table(engine, table_name):
    stmt = f"""CREATE TABLE IF NOT EXISTS {table_name} (
        video_id VARCHAR(255),
        channel_id VARCHAR(255),
        upload_time TIMESTAMP,
        view_count INT
    );"""
    with engine.connect() as connection:
        connection.execute(text(stmt))
        connection.commit()

if __name__ == "__main__":    
    try:
        execution_date = sys.argv[1]
        postgres_conn_string = sys.argv[2]
        youtube_api_key = sys.argv[3]
    except IndexError as e:
        print(f"Job requires three arguments but got {len(sys.argv)}, {sys.argv}")

    delay_days = 7
    job_interval_seconds = 10
    delayed_video_views_table_name = f"video_views_{delay_days}d"


    execution_date = datetime.strptime(execution_date, '%Y-%m-%d %H:%M:%S')
    engine = create_engine(postgres_conn_string)
    print(f"executing with date {execution_date}")


    print(f"Create seve day video views table...")
    create_video_views_table(engine, delayed_video_views_table_name)
    
    print(f"Getting 7 day old videos...")
    video_records = get_videos_from_range(engine, execution_date - timedelta(days=delay_days), execution_date - timedelta(days=delay_days, minutes=job_interval_seconds))
    

    print("Getting video views...")
    api_client = youtube.YouTubeAPI(youtube_api_key)
    records_with_views = get_batch_video_views(video_records, api_client)

    print("Updating one day video views")
    update_video_views_table(engine, records_with_views, delayed_video_views_table_name)