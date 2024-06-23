CREATE TABLE buyer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ(6),
    updated_at TIMESTAMPTZ(6)
);

CREATE TABLE seller (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    balance NUMERIC(19, 2) NOT NULL,
    created_at TIMESTAMPTZ(6),
    updated_at TIMESTAMPTZ(6)
);

CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    seller_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    location VARCHAR(255) NOT NULL,
    qty INTEGER NOT NULL,
    category VARCHAR(255) NOT NULL,
    price NUMERIC(19, 2) NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_event_seller FOREIGN KEY (seller_id) REFERENCES seller(id)
);

CREATE TABLE ticket (
    id SERIAL PRIMARY KEY,
    event_id INT NOT NULL,
    buyer_id INT NOT NULL,
    created_at TIMESTAMPTZ(6) NOT NULL,
    updated_at TIMESTAMPTZ(6) NOT NULL,
    status VARCHAR(50) DEFAULT 'valid' CHECK (status IN ('valid', 'used', 'expired')),
    CONSTRAINT fk_event FOREIGN KEY(event_id) REFERENCES event(id),
    CONSTRAINT fk_buyer FOREIGN KEY(buyer_id) REFERENCES buyer(id)
);

CREATE TABLE history (
    id SERIAL PRIMARY KEY,
    buyer_id INT NOT NULL,
    number VARCHAR(255) NOT NULL,
    payment_status VARCHAR(255) NOT NULL,
    total NUMERIC(19, 2) NOT NULL,
    paid_at TIMESTAMPTZ(6),
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