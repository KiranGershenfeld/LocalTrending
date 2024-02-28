#!/usr/bin/env python3

from datetime import datetime, timedelta
import os
import sys
import logging
import time
from dotenv import load_dotenv
from sqlalchemy import create_engine, text, exc
import youtube #common lib defined locally in airflow/common
import pandas as pd
import numpy as np
from collections import defaultdict

def get_videos_from_range(engine, start_date, end_date):
        # Define the query to select records from the "videos" table created within the past hour
    execution_date = datetime.strptime(execution_date, '%Y-%m-%d %H:%M:%S')
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

def get_batch_video_views(engine):
    logger = logging.getLogger(__name__)

    client = youtube.YouTubeAPI(os.environ.get("YOUTUBE_API_KEY"))

    with engine.connect() as connection:
        records = connection.execute(text("SELECT * FROM videos_1d_tmp"))
        connection.commit()
    print(records)
    if not records:
        return

    records_with_views = []

    for record in records:
        video_id = record[0]
        try:
            views = client.get_video_views(video_id, method='api')
        except Exception as e: 
            logger.error(f"Could not get video views for video {video_id}. Error: {e}")
            continue

        record_with_views = record + (views,)
        records_with_views.append(record_with_views)

    return records_with_views

def update_one_day_video_views(engine, records_with_views):
    stmt = """INSERT INTO video_views_1d (youtube_id, channel_id, upload_time, view_count)
              VALUES (:youtube_id, :channel_id, :upload_time, :view_count)
           """
    with engine.connect() as connection:
        for record_with_views in records_with_views:
            connection.execute(text(stmt), **record_with_views)
        connection.commit()

def create_seven_day_views_table(engine):
    stmt = """CREATE TABLE IF NOT EXISTS seven_day_video_views (
        video_id VARCHAR(255),
        channel_id VARCHAR(255),
        upload_time TIMESTAMP,
        view_count INT,
    );"""
    with engine.connect() as connection:
        connection.execute(text(stmt))
        connection.commit()

def db_conn():
    username = os.environ.get("DB_CREDENTIALS_USERNAME")
    password = os.environ.get("DB_CREDENTIALS_PASSWORD")
    host = os.environ.get("DB_CREDENTIALS_HOST")
    port = os.environ.get("DB_CREDENTIALS_PORT")
    name = os.environ.get("DB_CREDENTIALS_NAME")

    DATABASE_URL = f'postgresql://{username}:{password}@{host}:{port}/{name}'
    print(DATABASE_URL)
    engine = create_engine(DATABASE_URL)

    return engine

if __name__ == "__main__":
    load_dotenv()
    if len(sys.argv) < 2:
        execution_date = datetime.now().strftime('%Y-%m-%d %H:%M:%S') #Default value
    else:
        execution_date = sys.argv[1]

    print(f"executing with date {execution_date}")
    engine = db_conn()
    with engine.connect() as conn:
        print (conn.execute(text("SELECT 1")))

    print(f"Create seve day video views table...")
    create_seven_day_views_table(engine)
    
    print(f"Getting 7 day old videos...")
    video_records = get_videos_from_range(engine, execution_date - timedelta(days=1), execution_date - timedelta(days=1, minutes=10))
    
    print("Getting video views...")
    records_with_views = get_batch_video_views(engine)

    print("Updating one day video views")
    update_one_day_video_views(engine, records_with_views)