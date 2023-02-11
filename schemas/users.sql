CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username varchar(64) NOT NULL,
    hashed_password varchar(64) NOT NULL
)