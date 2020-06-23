CREATE DATABASE IF NOT EXISTS skillset;
USE skillset;
CREATE TABLE skillset.users (
    user_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_email VARCHAR (255) NOT NULL,
    user_name VARCHAR (255) NOT NULL,
    user_password VARCHAR (255) NOT NULL
);
INSERT INTO users (user_email,user_name,user_password) VALUES ("a@gmail.com","たろう","pass");