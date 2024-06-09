CREATE TYPE problem_status as enum('Easy', 'Medium', 'Hard');

CREATE TABLE 
    Problems(
        Id uuid primary key default gen_random_uuid() NOT NULL,
        Problem_num int unique NOT NULL,
        Title varchar NOT NULL,
        Status problem_status NOT NULL,
        Description text NOT NULL,
        Answer varchar NOT NULL
);
