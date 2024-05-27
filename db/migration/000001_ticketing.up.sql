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

CREATE TABLE ticket (
    id SERIAL PRIMARY KEY,
    event_id INT NOT NULL,
    buyer_id INT NOT NULL,
    date TIMESTAMPTZ(6) NOT NULL,
    location VARCHAR(255) NOT NULL,
    qty INT NOT NULL,
    price NUMERIC(19,2) NOT NULL,
    created_at TIMESTAMPTZ(6) NOT NULL,
    updated_at TIMESTAMPTZ(6) NOT NULL,
    CONSTRAINT fk_event FOREIGN KEY(event_id) REFERENCES event(id),
    CONSTRAINT fk_buyer FOREIGN KEY(buyer_id) REFERENCES buyer(id)
);

CREATE TABLE history (
    id SERIAL PRIMARY KEY,
    buyer_id INT NOT NULL,
    number TIMESTAMPTZ(6) NOT NULL,
    location VARCHAR(255) NOT NULL,
    qty INT NOT NULL,
    price NUMERIC(19,2) NOT NULL,
    created_at TIMESTAMPTZ(6) NOT NULL,
    updated_at TIMESTAMPTZ(6) NOT NULL,
    CONSTRAINT fk_buyer FOREIGN KEY(buyer_id) REFERENCES buyer(id)
);

CREATE TABLE history_item (
    id SERIAL PRIMARY KEY,
    history_id INT NOT NULL,
    event_id INT NOT NULL,
    price NUMERIC(19,2) NOT NULL,
    qty INT NOT NULL,
    subtotal NUMERIC(19,2) NOT NULL,
    created_at TIMESTAMPTZ(6),
    updated_at TIMESTAMPTZ(6),
    CONSTRAINT fk_history FOREIGN KEY(history_id) REFERENCES history(id),
    CONSTRAINT fk_event FOREIGN KEY(event_id) REFERENCES event(id)
);
