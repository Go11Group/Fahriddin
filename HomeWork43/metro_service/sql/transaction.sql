CREATE TYPE tr AS ENUM('credit','debit');

CREATE TABLE Transaction(
    Id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    Card_id uuid REFERENCES Card(id),
    Amount DECIMAL(10,2) NOT NULL,
    Terminal_id uuid DEFAULT NULL,
    Transaction_type tr NOT NULL DEFAULT 'debit'
);