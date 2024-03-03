import os
from typing import Type, Self
from dataclasses import dataclass

@dataclass
class Config:
    db_password: str
    db_username: str
    db_host: str
    db_port: str
    db_name: str
    db_sslmode: str
    youtube_api_key: str

    def valid(self) -> Self:
        for var_name, var_value in self.__dict__.items():
            if var_value is None:
                raise ValueError(f"Config {var_name} is None")
        return self

def InitConfig() -> Config:
    return Config(
        db_password=os.environ.get("DB_PASSWORD"),
        db_username=os.environ.get("DB_USERNAME"),
        db_host=os.environ.get("DB_HOST"),
        db_port=os.environ.get("DB_PORT"),
        db_name=os.environ.get("DB_NAME"),
        db_sslmode=os.environ.get("DB_SSLMODE"),
        youtube_api_key=os.environ.get("YOUTUBE_API_KEY")
    )