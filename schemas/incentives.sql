-- market
DROP TABLE IF EXISTS incentives;
CREATE TABLE IF NOT EXISTS incentives (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username varchar(64) NOT NULL,
    incentive_code varchar(10) NOT NULL,
    transfer_code varchar(64),
    value int NOT NULL
);