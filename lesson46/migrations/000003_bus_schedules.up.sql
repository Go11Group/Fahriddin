CREATE TABLE bus_schedules (
    id SERIAL PRIMARY KEY,
    bus_number VARCHAR(255) NOT NULL,
    schedule_time TIME NOT NULL
);