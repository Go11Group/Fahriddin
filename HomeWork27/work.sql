Create Table Cars(Id uuid Primary Key not null default gen_random_uuid(),
    Brand varchar,
    Year int,
    Price int
    );

Create Table Users(Id uuid Primary Key not null default gen_random_uuid(),
Name varchar,
Car_id uuid not null ,
Foreign Key (Car_id) References Cars(id)
);

