DROP TABLE IF EXISTS characters;

CREATE TABLE characters (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    species VARCHAR(50) NOT NULL,
    type VARCHAR(50),
    gender VARCHAR(20),
    image VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    created VARCHAR(100) NOT NULL

);
