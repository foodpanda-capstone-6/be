-- market
DROP TABLE IF EXISTS incentives;
CREATE TABLE IF NOT EXISTS incentives (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username varchar(64) NOT NULL,
    incentive_code varchar(10) NOT NULL UNIQUE,
    transfer_code varchar(64) DEFAULT "",
    value int NOT NULL
);
INSERT INTO incentives (
        id,
        username,
        incentive_code,
        transfer_code,
        value
    )
VALUES (1001, "kai", "SAMP001", "", 5),
    (1002, "naz", "SAMP002", "", 5),
    (1003, "noel", "SAMP003", "", 5),
    (1004, "kai", "SAMP004", "", 5);