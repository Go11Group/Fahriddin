CREATE TABLE Courses(
    course_id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    title varchar(100),
    description text,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

INSERT INTO Courses (title, description) VALUES
('Math', 'A comprehensive study of Math'),
('Science', 'An introductory course on Science'),
('History', 'A survey of History'),
('Literature', 'An in-depth exploration of Literature'),
('Art', 'The history and development of Art'),
('Music', 'Advanced topics in Music'),
('Biology', 'Principles of Biology'),
('Chemistry', 'A detailed examination of Chemistry'),
('Physics', 'A beginners guide to Physics'),
('Geography', 'Fundamental concepts in Geography'),
('Math', 'Advanced topics in Math'),
('Science', 'An in-depth exploration of Science'),
('History', 'A detailed examination of History'),
('Literature', 'A survey of Literature'),
('Art', 'A comprehensive study of Art'),
('Music', 'An introductory course on Music'),
('Biology', 'The history and development of Biology'),
('Chemistry', 'A beginners guide to Chemistry'),
('Physics', 'Principles of Physics'),
('Geography', 'Advanced topics in Geography'),
('Math', 'Fundamental concepts in Math'),
('Science', 'A detailed examination of Science'),
('History', 'The history and development of History'),
('Literature', 'A comprehensive study of Literature'),
('Art', 'An introductory course on Art'),
('Music', 'Principles of Music'),
('Biology', 'A survey of Biology'),
('Chemistry', 'Fundamental concepts in Chemistry'),
('Physics', 'An in-depth exploration of Physics'),
('Geography', 'Advanced topics in Geography'),
('Math', 'A detailed examination of Math'),
('Science', 'An introductory course on Science'),
('History', 'Advanced topics in History'),
('Literature', 'Principles of Literature'),
('Art', 'An in-depth exploration of Art'),
('Music', 'A comprehensive study of Music'),
('Biology', 'An introductory course on Biology'),
('Chemistry', 'Advanced topics in Chemistry'),
('Physics', 'A survey of Physics'),
('Geography', 'An in-depth exploration of Geography'),
('Math', 'The history and development of Math'),
('Science', 'Principles of Science'),
('History', 'A beginners guide to History'),
('Literature', 'Fundamental concepts in Literature'),
('Art', 'A detailed examination of Art'),
('Music', 'An introductory course on Music'),
('Biology', 'The history and development of Biology'),
('Chemistry', 'A comprehensive study of Chemistry'),
('Physics', 'A survey of Physics'),
('Geography', 'An in-depth exploration of Geography');
