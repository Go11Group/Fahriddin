CREATE TABLE Users(
    Id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    Surname varchar NOT NULL,
    Lastname varchar NOT NULL,
    Email varchar NOT NULL,
    Password varchar NOT NULL
);