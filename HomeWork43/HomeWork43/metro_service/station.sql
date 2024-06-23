CREATE TABLE Station(
    Id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    Name VARCHAR NOT NULL
);