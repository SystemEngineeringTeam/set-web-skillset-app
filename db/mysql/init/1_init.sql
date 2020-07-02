CREATE DATABASE IF NOT EXISTS skillset;
USE skillset;
CREATE TABLE skillset.users (
    user_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_email VARCHAR (255) NOT NULL,
    user_name VARCHAR (255) NOT NULL,
    user_password VARCHAR (255) NOT NULL
);
INSERT INTO users (user_email, user_name, user_password)
VALUES ("a@gmail.com", "たろう", "pass");
CREATE TABLE skillset.grades (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    text VARCHAR (255) NOT NULL
);
INSERT INTO grades (text)
VALUES ("1年");
CREATE TABLE skillset.majors (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    text VARCHAR (255) NOT NULL
);
CREATE TABLE skillset.groups (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    text VARCHAR (255) NOT NULL
);
CREATE TABLE skillset.technical_areas (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    text VARCHAR (255) NOT NULL
);
CREATE TABLE skillset.technologies (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    text VARCHAR (255) NOT NULL
);
CREATE TABLE skillset.members (
    id VARCHAR NOT NULL PRIMARY KEY,
    image VARCHAR,
    name VARCHAR,
    message VARCHAR,
    slack VARCHAR,
    twitter VARCHAR,
    github VARCHAR,
    grade_id INT,
    major_id INT,
    created_at DATETIME default current_timestamp,
    updated_at DATETIME default current_timestamp ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_grade_id FOREIGN KEY (grade_id) REFERENCES grades (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT fk_major_id FOREIGN KEY (major_id) REFERENCES majors (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
CREATE TABLE skillset.groups_of_members (
    member_id VARCHAR NOT NULL PRIMARY KEY,
    group_id INT NOT NULL PRIMARY KEY,
    CONSTRAINT fk_member_id FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT fk_group_id FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
CREATE TABLE skillset.technical_areas_of_members (
    member_id VARCHAR NOT NULL PRIMARY KEY,
    technical_area_id INT NOT NULL PRIMARY KEY,
    CONSTRAINT fk_member_id FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT fk_technical_area_id FOREIGN KEY (technical_area_id) REFERENCES technical_areas (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
CREATE TABLE skillset.technologies_of_members (
    member_id VARCHAR NOT NULL PRIMARY KEY,
    technologie_id INT NOT NULL PRIMARY KEY,
    CONSTRAINT fk_member_id FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT fk_technologie_id FOREIGN KEY (technologie_id) REFERENCES technologies (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
CREATE TABLE skillset.products_of_members (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    member_id VARCHAR NOT NULL,
    overview VARCHAR,
    url VARCHAR,
    CONSTRAINT fk_member_id FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);