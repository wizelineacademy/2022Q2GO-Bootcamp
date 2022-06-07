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
    created VARCHAR(100) NOT NULL );


INSERT INTO characters (id,name,status,species,"type",gender,image,url,created) VALUES
	 (1,'Rick Sanchez','Alive','Human','','Male','https://rickandmortyapi.com/api/character/avatar/1.jpeg','https://rickandmortyapi.com/api/character/1','2017-11-04T18:48:46.250Z'),
	 (2,'Morty Smith','Alive','Human','','Male','https://rickandmortyapi.com/api/character/avatar/2.jpeg','https://rickandmortyapi.com/api/character/2','2017-11-04T18:50:21.651Z'),
	 (3,'Summer Smith','Alive','Human','','Female','https://rickandmortyapi.com/api/character/avatar/3.jpeg','https://rickandmortyapi.com/api/character/3','2017-11-04T19:09:56.428Z'),
	 (4,'Beth Smith','Alive','Human','','Female','https://rickandmortyapi.com/api/character/avatar/4.jpeg','https://rickandmortyapi.com/api/character/4','2017-11-04T19:22:43.665Z'),
	 (5,'Jerry Smith','Alive','Human','','Male','https://rickandmortyapi.com/api/character/avatar/5.jpeg','https://rickandmortyapi.com/api/character/5','2017-11-04T19:26:56.301Z'),
	 (6,'Abadango Cluster Princess','Alive','Alien','','Female','https://rickandmortyapi.com/api/character/avatar/6.jpeg','https://rickandmortyapi.com/api/character/6','2017-11-04T19:50:28.250Z'),
	 (7,'Abradolf Lincler','unknown','Human','Genetic experiment','Male','https://rickandmortyapi.com/api/character/avatar/7.jpeg','https://rickandmortyapi.com/api/character/7','2017-11-04T19:59:20.523Z'),
	 (8,'Adjudicator Rick','Dead','Human','','Male','https://rickandmortyapi.com/api/character/avatar/8.jpeg','https://rickandmortyapi.com/api/character/8','2017-11-04T20:03:34.737Z'),
	 (9,'Agency Director','Dead','Human','','Male','https://rickandmortyapi.com/api/character/avatar/9.jpeg','https://rickandmortyapi.com/api/character/9','2017-11-04T20:06:54.976Z'),
	 (10,'Alan Rails','Dead','Human','Superhuman (Ghost trains summoner)','Male','https://rickandmortyapi.com/api/character/avatar/10.jpeg','https://rickandmortyapi.com/api/character/10','2017-11-04T20:19:09.017Z');

INSERT INTO characters (id,name,status,species,"type",gender,image,url,created) VALUES
	 (11,'Albert Einstein','Dead','Human','','Male','https://rickandmortyapi.com/api/character/avatar/11.jpeg','https://rickandmortyapi.com/api/character/11','2017-11-04T20:20:20.965Z'),
	 (12,'Alexander','Dead','Human','','Male','https://rickandmortyapi.com/api/character/avatar/12.jpeg','https://rickandmortyapi.com/api/character/12','2017-11-04T20:32:33.144Z'),
	 (13,'Alien Googah','unknown','Alien','','unknown','https://rickandmortyapi.com/api/character/avatar/13.jpeg','https://rickandmortyapi.com/api/character/13','2017-11-04T20:33:30.779Z'),
	 (14,'Alien Morty','unknown','Alien','','Male','https://rickandmortyapi.com/api/character/avatar/14.jpeg','https://rickandmortyapi.com/api/character/14','2017-11-04T20:51:31.373Z'),
	 (15,'Alien Rick','unknown','Alien','','Male','https://rickandmortyapi.com/api/character/avatar/15.jpeg','https://rickandmortyapi.com/api/character/15','2017-11-04T20:56:13.215Z'),
	 (16,'Amish Cyborg','Dead','Alien','Parasite','Male','https://rickandmortyapi.com/api/character/avatar/16.jpeg','https://rickandmortyapi.com/api/character/16','2017-11-04T21:12:45.235Z'),
	 (17,'Annie','Alive','Human','','Female','https://rickandmortyapi.com/api/character/avatar/17.jpeg','https://rickandmortyapi.com/api/character/17','2017-11-04T22:21:24.481Z'),
	 (18,'Antenna Morty','Alive','Human','Human with antennae','Male','https://rickandmortyapi.com/api/character/avatar/18.jpeg','https://rickandmortyapi.com/api/character/18','2017-11-04T22:25:29.008Z'),
	 (19,'Antenna Rick','unknown','Human','Human with antennae','Male','https://rickandmortyapi.com/api/character/avatar/19.jpeg','https://rickandmortyapi.com/api/character/19','2017-11-04T22:28:13.756Z'),
	 (20,'Ants in my Eyes Johnson','unknown','Human','Human with ants in his eyes','Male','https://rickandmortyapi.com/api/character/avatar/20.jpeg','https://rickandmortyapi.com/api/character/20','2017-11-04T22:34:53.659Z');