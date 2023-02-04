CREATE TABLE sample5(
                       id INT NOT NULL AUTO_INCREMENT,
                       name  VARCHAR(10) NOT NULL,
                       genre VARCHAR(15) NOT NULL,
                       rating FLOAT NOT NULL,
                       releasedDate DATE NOT NULL,
                       updatedAt TIMESTAMP NOT NULL ,
                       createdAt TIMESTAMP NOT NULL ,
                       plot varchar(100) NOT NULL,
                       released TINYINT NOT NULL,
                       deletedAt DATETIME ,
                       primary key (id)
)
-- CREATE TABLE sample2{
--     id int PRIMARY KEY not null,
--     name varchar(20) not null,
--     genre varchar(20) not null,
--     rating float not null,
--     releaseDate date,
--     plot varchar(100) not null,
--     released boolean,
--     };