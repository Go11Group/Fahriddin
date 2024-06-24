CREATE TABLE Users(
    Id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    Name VARCHAR NOT NULL,
    Phone VARCHAR NOT NULL,
    Age INT NOT NULL
);

