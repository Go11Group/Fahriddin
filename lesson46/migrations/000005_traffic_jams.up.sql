CREATE TABLE traffic_jams (
    id SERIAL PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    reported_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);