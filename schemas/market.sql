-- market
DROP TABLE IF EXISTS market;
CREATE TABLE IF NOT EXISTS market (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    description varchar(64) NOT NULL,
    amount int NOT NULL
);
INSERT INTO market (description, amount)
VALUES("indesc1", 10),
    ("indesc2", 20),
    ("indesc3", 30);