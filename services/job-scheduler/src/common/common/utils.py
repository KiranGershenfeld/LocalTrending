import time
from sqlalchemy import create_engine
from sqlalchemy.exc import OperationalError

def create_engine_with_retry(connection_string, max_retries=5, base_delay=1, max_delay=30):
    """
    Create a SQLAlchemy engine with retry logic in case of connection failure.

    Args:
        connection_string (str): The connection string for the database.
        max_retries (int): The maximum number of retries.
        base_delay (int): The base delay time (in seconds) for the initial retry.
        max_delay (int): The maximum delay time (in seconds) for the progressive backoff.

    Returns:
        sqlalchemy.engine.base.Engine: SQLAlchemy engine object.
    """
    retry_count = 0
    delay = base_delay

    while retry_count < max_retries:
        try:
            engine = create_engine(connection_string)
            engine.connect()
            print("Connection established successfully.")
            return engine
        except OperationalError as e:
            print(f"Connection failed. Retrying... ({retry_count+1}/{max_retries})")
            if retry_count == max_retries - 1:
                raise e
            retry_count += 1
            time.sleep(delay)
            delay = min(delay * 2, max_delay)
