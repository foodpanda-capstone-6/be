CREATE TABLE IF NOT EXISTS market (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username varchar(64) NOT NULL,
    hashed_password varchar(64) NOT NULL
)