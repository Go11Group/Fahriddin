CREATE TABLE 
    Users(
        Id uuid PRIMARY KEY default gen_random_uuid() NOT NULL,
        Name varchar,
        Age int,
        Email varchar,
        Password varchar
    );