CREATE TABLE Terminal(
    Id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    Station_id uuid REFERENCES Station(id)
)