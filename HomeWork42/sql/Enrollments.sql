CREATE TABLE Enrollments(
    enrollment_id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid REFERENCES Users(user_id),
    course_id uuid REFERENCES Courses(course_id),
    enrollment_date timestamp DEFAULT CURRENT_TIMESTAMP,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    UNIQUE (course_id,user_id)
);

INSERT INTO Enrollments (user_id, course_id) VALUES
('1fddb3e3-0e7b-4100-97e0-cde2f0fc9706', '1e023412-da90-47ea-8950-91773052b103'),
('1fddb3e3-0e7b-4100-97e0-cde2f0fc9706', '0dac2ca2-e98c-460b-bf61-f0fb26d0e8de'),
('ced7f583-b3ba-4558-a36a-d5086c7fe604', '08ceb17a-3f33-4081-85ae-047fef4efece');