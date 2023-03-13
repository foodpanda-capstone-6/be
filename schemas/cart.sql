-- cart
DROP TABLE IF EXISTS cart;
CREATE TABLE IF NOT EXISTS cart (
    username VARCHAR(64),
    market_voucher_id INTEGER,
    qty INTEGER DEFAULT 0
);
INSERT INTO cart (username, market_voucher_id, qty)
VALUES ("kai", 12, 1),
    ("noel", 12, 2),
    ("noel", 22, 3),
    ("noel", 22, 2);