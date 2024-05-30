Create table Course(
    Id uuid Primary Key Not Null Default gen_random_uuid(),
    Name varchar
    Field varchar
);
Create type gen AS ENUM ('M','F');

Create table Student(
    Id uuid Primary Key Not Null Default gen_random_uuid(),
    Name varchar,
    Age int,
    Gender gen,
    Course_id uuid References Course(id)
);