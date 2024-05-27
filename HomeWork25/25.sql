Create table Users(
    Id UUID Primary Key Not Null default gen_random_uuid(),
    Username Varchar Not null,
    Email Varchar Not null
    );

Create Table Cars (
    Id UUID Primary Key Not null Default gen_random_uuid(),
    Brand varchar not null,
    Year int,
    Price int
);

CREATE TABLE users_cars (
    user_id UUID NOT NULL,
    car_id UUID NOT NULL,
    PRIMARY KEY (user_id, car_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    FOREIGN KEY (car_id) REFERENCES cars (id) ON DELETE CASCADE
);

INSERT INTO users (username, email) VALUES 
('user1', 'user1@example.com'),
('user2', 'user2@example.com'),
('user3', 'user3@example.com'),
('user4', 'user4@example.com'),
('user5', 'user5@example.com'),
('user6', 'user6@example.com'),
('user7', 'user7@example.com'),
('user8', 'user8@example.com'),
('user9', 'user9@example.com'),
('user10', 'user10@example.com');

INSERT INTO cars (brand, year, price) VALUES 
('Toyota', 2020, 20000),
('Honda', 2019, 18000),
('Ford', 2018, 22000),INSERT INTO cars (brand, year, price) VALUES 
('Toyota', 2020, 20000),
('Honda', 2019, 18000),
('Ford', 2018, 22000),
('Chevrolet', 2021, 25000),
('Nissan', 2020, 21000),
('BMW', 2019, 35000),
('Audi', 2021, 40000),
('Mercedes', 2020, 45000),
('Kia', 2018, 17000),
('Hyundai', 2019, 19000);
('Chevrolet', 2021, 25000),
('Nissan', 2020, 21000),
('BMW', 2019, 35000),
('Audi', 2021, 40000),
('Mercedes', 2020, 45000),
('Kia', 2018, 17000),
('Hyundai', 2019, 19000);

Insert Into users_cars Values
(
    '8c1919a5-156a-4638-9df0-0f2ec07d802e','6d1ad3dd-025a-4d6a-b6c0-6e9ff33c13b0',
    '533883d7-27e3-4e2c-aca9-6017e9ca9acb','6d1ad3dd-025a-4d6a-b6c0-6e9ff33c13b0',
    '5d7322cd-48ae-41ff-8285-8cfb1fff7011','0b2df790-0b90-48d9-b2f5-2e7c432f5b46',
    '5fdfebff-493e-4dc7-84e6-c2911e25227e','0b2df790-0b90-48d9-b2f5-2e7c432f5b46',
    'e8a4bbca-5433-4102-8e7c-65ca053a83d4','9204de16-c1d0-485e-a69e-d9bad4e5b1c9',
    '939fbfbb-94f9-4baa-bc62-99ebcf14444f','9204de16-c1d0-485e-a69e-d9bad4e5b1c9',
    'a1872763-d168-4bde-b50b-a6d0e63f918c','20ac0b40-6ab9-4626-a750-b868c82cf1a3',
    '401f398a-136f-4cea-a21a-5fc538515d8c','20ac0b40-6ab9-4626-a750-b868c82cf1a3',
    '7d75c6c5-17bc-4f3a-b539-7a3242eaa3e7','1fec887e-ce03-4a7c-8575-a2d29e2762e0',
    '14987d22-6f09-4fd5-9100-5f465bededac','1fec887e-ce03-4a7c-8575-a2d29e2762e0',
    '8c1919a5-156a-4638-9df0-0f2ec07d802e','033f2824-6f11-4551-8dd8-bc7518d91da6',
    '533883d7-27e3-4e2c-aca9-6017e9ca9acb','033f2824-6f11-4551-8dd8-bc7518d91da6',
    '5d7322cd-48ae-41ff-8285-8cfb1fff7011','e7a43b80-9c37-4eb0-91c5-e0b14b543ac4',
    '5fdfebff-493e-4dc7-84e6-c2911e25227e','e7a43b80-9c37-4eb0-91c5-e0b14b543ac4',
    'e8a4bbca-5433-4102-8e7c-65ca053a83d4','226932fc-5294-4eb6-b34b-da0542bf5623',
    '939fbfbb-94f9-4baa-bc62-99ebcf14444f','226932fc-5294-4eb6-b34b-da0542bf5623',
    'a1872763-d168-4bde-b50b-a6d0e63f918c','dd73f76e-136a-4e3a-bcd4-abb7a3d8b06d',
    '401f398a-136f-4cea-a21a-5fc538515d8c','dd73f76e-136a-4e3a-bcd4-abb7a3d8b06d',
    '7d75c6c5-17bc-4f3a-b539-7a3242eaa3e7','8526f9df-d6d2-420f-bd76-01a21a92f0bf',
    '14987d22-6f09-4fd5-9100-5f465bededac','8526f9df-d6d2-420f-bd76-01a21a92f0bf'

)



