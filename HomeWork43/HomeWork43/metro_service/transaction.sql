CREATE TYPE tr ENUM('credit','debit') 

CREATE TABLE Transaction(
    Id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    Card_id uuid REFERENCES Card(id),
    Amount real,
    Terminal_id uuid DEFAULT NULL,
    Transaction_type tr
)