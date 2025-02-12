CREATE DATABASE books_management;
USE books_management;

CREATE TABLE books (
    id             INT AUTO_INCREMENT PRIMARY KEY,
    name           VARCHAR(255) NOT NULL UNIQUE,
    author_name    VARCHAR(255) NOT NULL,
    price          DECIMAL(10,2) NOT NULL,
    available      INT NOT NULL,
    issued        INT NOT NULL DEFAULT 0,
    publisher      VARCHAR(255),
    published_year INT,
    description    TEXT,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
