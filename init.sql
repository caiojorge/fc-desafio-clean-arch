CREATE DATABASE IF NOT EXISTS orders;
USE orders;

CREATE TABLE IF NOT EXISTS orders (
    id varchar(255) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,
    PRIMARY KEY (id)
);

-- Inserir alguns dados de exemplo
INSERT INTO orders (id, price, tax, final_price) VALUES
('order-1', 100.0, 10.0, 110.0),
('order-2', 200.0, 20.0, 220.0),
('order-3', 50.0, 5.0, 55.0)
ON DUPLICATE KEY UPDATE
    price = VALUES(price),
    tax = VALUES(tax),
    final_price = VALUES(final_price);
