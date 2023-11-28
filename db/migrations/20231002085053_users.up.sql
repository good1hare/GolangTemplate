CREATE TABLE IF NOT EXISTS users
(
    id         BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name  VARCHAR(255),
    phone INT,
    created_at DATETIME,
    updated_at DATETIME
);

