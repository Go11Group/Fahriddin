CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);

CREATE TABLE UserProduct(
    id SERIAL PRIMARY KEY,
    user_id SERIAL REFERENCES Users(id),
    product_id SERIAL REFERENCES Products(id)
);

-- Users jadvalini to'ldirish
INSERT INTO Users (username, email, password) VALUES
('bekzod01', 'bekzod01@example.com', 'password1'),
('davron02', 'davron02@example.com', 'password2'),
('madina03', 'madina03@example.com', 'password3'),
('javohir04', 'javohir04@example.com', 'password4'),
('laziz05', 'laziz05@example.com', 'password5'),
('shahnoza06', 'shahnoza06@example.com', 'password6'),
('kamila07', 'kamila07@example.com', 'password7'),
('zulfiya08', 'zulfiya08@example.com', 'password8'),
('sarvar09', 'sarvar09@example.com', 'password9'),
('aziza10', 'aziza10@example.com', 'password10'),
('shoxrux11', 'shoxrux11@example.com', 'password11'),
('oybek12', 'oybek12@example.com', 'password12'),
('xushnud13', 'xushnud13@example.com', 'password13'),
('momin14', 'momin14@example.com', 'password14'),
('farzona15', 'farzona15@example.com', 'password15');

-- Products jadvalini to'ldirish
INSERT INTO Products (name, description, price, stock_quantity) VALUES
('Laptop', 'Yangi model laptop', 1200.00, 10),
('Smartfon', 'Songgi model smartfon', 800.00, 15),
('Planshet', 'Keng ekranli planshet', 400.00, 20),
('Televizor', 'Smart TV', 600.00, 5),
('Muzlatkich', 'Ikki kamerali muzlatkich', 1000.00, 7),
('Kir yuvish mashinasi', 'Avtomatik kir yuvish mashinasi', 700.00, 8),
('Mikrotolqinli pech', 'Yangi model mikrotolqinli pech', 150.00, 25),
('Konditsioner', 'Energiya tejamkor konditsioner', 500.00, 12),
('Quritgich', 'Kiyim quritgich', 300.00, 18),
('Qoshiq toplami', 'Zamonaviy dizayn qoshiq toplami', 50.00, 30),
('Telefon zaryadkachi', 'Tez zaryadlovchi', 20.00, 50),
('Quloqchin', 'Bluetooth quloqchin', 80.00, 40),
('Uy printeri', 'Kichik uy printeri', 200.00, 9),
('Kamera', 'Raqamli kamera', 350.00, 14),
('Elektron soat', 'Sport uchun elektron soat', 100.00, 35);

-- UserProduct jadvalini to'ldirish
INSERT INTO UserProduct (user_id, product_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(9, 9),
(10, 10),
(11, 11),
(12, 12),
(13, 13),
(14, 14),
(15, 15),
(1, 2),
(2, 3),
(3, 4),
(4, 5),
(5, 6),
(6, 7),
(7, 8),
(8, 9),
(9, 10),
(10, 11),
(11, 12),
(12, 13),
(13, 14),
(14, 15),
(15, 1);
