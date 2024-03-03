from config import Config

def CreatePostgresConnectionString(config: Config) -> str:
	return f"postgresql://{config.db_username}:{config.db_password}@{config.db_host}:{config.db_port}/{config.db_name}?sslmode={config.db_sslmode}"