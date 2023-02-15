CREATE TABLE Movie{
    Id int PRIMARY KEY not null,
    Name varchar(20) not null,
    Genre varchar(20) not null,
    Rating float not null,
    ReleaseDate date not null,
    UpdatedAt timestamp not null,
    CreatedAt date time not null,
    Plot varchar(200) not null,
    DeletedAt date,
    released boolean,
    };