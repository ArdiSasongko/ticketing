CREATE TABLE buyer (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL

);

CREATE TABLE seller (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        email VARCHAR(255) UNIQUE NOT NULL,
                        password VARCHAR(255) NOT NULL
);

CREATE TABLE event (
                       id SERIAL PRIMARY KEY,
                       seller_id INTEGER NOT NULL,
                       name VARCHAR(255) UNIQUE NOT NULL,
                       date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       location VARCHAR(255) UNIQUE NOT NULL,
                       qty INTEGER NOT NULL,
                       category VARCHAR(255) UNIQUE NOT NULL,
                       price NUMERIC(19, 2) NOT NULL,
                       constraint seller_id FOREIGN KEY (seller_id) REFERENCES seller(id)
);
