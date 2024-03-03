import subprocess
import time
import logging
from datetime import datetime

import utils
import config

import schedule


def seven_day_video_job():
    subprocess.run(['python', 'jobs/delayed_video_views/job.py', datetime.now().strftime('%Y-%m-%d %H:%M:%S'), db_conn, config.youtube_api_key])

if __name__ == '__main__':
    logger = logging.getLogger(__name__)
    logger.info("Job scheduler running")
    config = config.InitConfig().valid()
    db_conn = utils.CreatePostgresConnectionString(config)

    schedule.every(10).minutes.do(seven_day_video_job)

    while True:
        schedule.run_pending()
        time.sleep(5)
    