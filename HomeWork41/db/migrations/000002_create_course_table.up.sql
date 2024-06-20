CREATE TABLE Universty(
    Id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    User_id uuid REFERENCES Users(Id),
    Title varchar NOT NULL,
    Description varchar NOT NULL
);

migrate -database 'postgres://postgres:0412@localhost:5432/nt?sslmode=disable'