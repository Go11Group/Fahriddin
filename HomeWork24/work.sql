create table author
(
    id   int primary key,
    name varchar not null
);
create table book
(
    id          int primary key,
    name        varchar not null,
    page        int4,
    author_id   int 
    foreign key (author_id) references author(id)
);

INSERT INTO author (id, name) VALUES
(1, 'Abdulla Qodiriy'),
(2, 'Abdulla Oripov'),
(3, 'Erkin Vohidov'),
(4, 'Muhammad Yusuf'),
(5, 'Otkir Hoshimov');

INSERT INTO book (id, name, page, author_id) VALUES
(1, 'Otkan kunlar', 350, 1),
(2, 'Mehrobdan chayon', 320, 1),
(3, 'Obid ketmon', 250, 1),
(4, 'Hijron kunlari', 270, 2),
(5, 'Sarob', 310, 2),
(6, 'Topmoq va yoqotmoq', 230, 2),
(7, 'Ruhlar isyoni', 300, 3),
(8, 'Oltin devor', 340, 3),
(9, 'Qoshiqlarim sizga', 280, 3),
(10, 'Ishq savdosi', 220, 4),
(11, 'Sevgi iztirobi', 260, 4),
(12, 'Uchrashuv', 290, 4),
(13, 'Dunyoning ishlari', 400, 5),
(14, 'Ikki eshik orasi', 380, 5),
(15, 'Kohna dunyo', 350, 5);



