CREATE TABLE
    Solved_problems(
        Id uuid primary KEY default gen_random_uuid() NOT NULL,
        User_id uuid REFERENCES Users(Id),
        Problem_id uuid REFERENCES Problems(Id),
        Answer_problem varchar,
        Result BOOLEAN
    );
