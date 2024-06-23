CREATE TABLE Card(
    Id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    Number VARCHAR UNIQUE,
    User_id uuid REFERENCES User(Id)
)