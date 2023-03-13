-- cart
DROP TABLE IF EXISTS cart;
CREATE TABLE IF NOT EXISTS cart (
    username VARCHAR(64),
    market_voucher_id INTEGER,
    qty INTEGER DEFAULT 0
);
INSERT INTO cart (username, market_voucher_id, qty)
VALUES ("kai", 1, 1),
    ("noel", 1, 2),
    ("noel", 2, 3),
    ("noel", 2, 2);