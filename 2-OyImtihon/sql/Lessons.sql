CREATE TABLE Lessons(
    lesson_id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    course_id uuid REFERENCES Courses(course_id),
    title varchar(100),
    content text,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

INSERT INTO Lessons (course_id, title, content) VALUES
('1e023412-da90-47ea-8950-91773052b103', 'Lesson 1', 'Content for lesson 1'),
('1e023412-da90-47ea-8950-91773052b103', 'Lesson 2', 'Content for lesson 2'),
('0dac2ca2-e98c-460b-bf61-f0fb26d0e8de', 'Lesson 3', 'Content for lesson 3');
