Create table student(
    Id Serial Primary Key not null,
    Name Varchar not null,
    Age int
);

Create table course(
    Id Serial Primary Key not null,
    Name Varchar not null  
);

Create table Student_course(
    Id Serial Primary Key not null,
    student_id int References sutdent(id),
    course_id int References course(id)
);

Create table Grade(
    Id Serial Primary Key not null,
    Student_course_id int References Student_course(id),
    Grade_student int
);

INSERT INTO student (Name, Age) VALUES ('Ali Valiyev', 20);
INSERT INTO student (Name, Age) VALUES ('Zarina Karimova', 22);
INSERT INTO student (Name, Age) VALUES ('Bekzod Nurmatov', 19);
INSERT INTO student (Name, Age) VALUES ('Dilnoza Abduvahobova', 21);
INSERT INTO student (Name, Age) VALUES ('Sardor Rahimov', 23);
INSERT INTO student (Name, Age) VALUES ('Otabek Ismoilov', 24);
INSERT INTO student (Name, Age) VALUES ('Nodira Yusufova', 20);
INSERT INTO student (Name, Age) VALUES ('Shavkat Xudoyberdiev', 22);
INSERT INTO student (Name, Age) VALUES ('Laylo Mirzayeva', 18);
INSERT INTO student (Name, Age) VALUES ('Javohir Sattorov', 25);

INSERT INTO course (Name) VALUES ('Matematika');
INSERT INTO course (Name) VALUES ('Ingliz tili');
INSERT INTO course (Name) VALUES ('Dasturlash asoslari');
INSERT INTO course (Name) VALUES ('Fizika');
INSERT INTO course (Name) VALUES ('Biologiya');

INSERT INTO student_course (student_id, course_id) VALUES (1, 1);
INSERT INTO student_course (student_id, course_id) VALUES (1, 2);
INSERT INTO student_course (student_id, course_id) VALUES (2, 1);
INSERT INTO student_course (student_id, course_id) VALUES (2, 3);
INSERT INTO student_course (student_id, course_id) VALUES (3, 2);
INSERT INTO student_course (student_id, course_id) VALUES (3, 4);
INSERT INTO student_course (student_id, course_id) VALUES (4, 3);
INSERT INTO student_course (student_id, course_id) VALUES (4, 5);
INSERT INTO student_course (student_id, course_id) VALUES (5, 1);
INSERT INTO student_course (student_id, course_id) VALUES (5, 4);
INSERT INTO student_course (student_id, course_id) VALUES (6, 2);
INSERT INTO student_course (student_id, course_id) VALUES (6, 5);
INSERT INTO student_course (student_id, course_id) VALUES (7, 3);
INSERT INTO student_course (student_id, course_id) VALUES (7, 1);
INSERT INTO student_course (student_id, course_id) VALUES (8, 2);

INSERT INTO Grade (Student_course_id, Grade_student) VALUES (1, 5);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (2, 4);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (3, 3);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (4, 5);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (5, 4);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (6, 3);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (7, 5);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (8, 4);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (9, 5);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (10, 3);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (11, 4);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (12, 5);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (13, 3);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (14, 4);
INSERT INTO Grade (Student_course_id, Grade_student) VALUES (15, 5);

--3
Select course.Name,AVG(Grade.Grade_student) As Middle_Grade  from student_course
Join
student
On student_course.student_id = student.id
Join
course
On student_course.course_id = course.id
Join 
Grade
On student_course.id = Grade.Student_course_id
Group By course.Name
Order By Middle_Grade Desc;

--4
Select course.Name,array_agg(student.name),Min(student.Age) As The_youngs  from student_course
Join
student
On student_course.student_id = student.id
Join
course
On student_course.course_id = course.id
Join 
Grade
On student_course.id = Grade.Student_course_id
Group By course.Name

--5
Select course.Name,AVG(Grade.Grade_student) As Middle_Grade  from student_course
Join
student
On student_course.student_id = student.id
Join
course
On student_course.course_id = course.id
Join 
Grade
On student_course.id = Grade.Student_course_id
Group By course.Name
Order By Middle_Grade Desc Limit 1;

--2
Select course.Name,array_agg(student.name),Max(Grade.Grade_student) As Middle_Grade  from student_course
Join
student
On student_course.student_id = student.id
Join
course
On student_course.course_id = course.id
Join 
Grade
On student_course.id = Grade.Student_course_id
Group By course.Name
Order By Middle_Grade Desc;


Select * from student_course
Join
student
On student_course.student_id = student.id
Join
course
On student_course.course_id = course.id
Join 
Grade
On student_course.id = Grade.Student_course_id


