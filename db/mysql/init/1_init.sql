DROP DATABASE IF EXISTS skillset;
CREATE DATABASE IF NOT EXISTS skillset;
USE skillset;
CREATE TABLE skillset.members (
    id VARCHAR(255) PRIMARY KEY,
    image VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    message VARCHAR(255),
    slack VARCHAR(255) NOT NULL,
    twitter VARCHAR(255),
    github VARCHAR(255),
    grade_id INT NOT NULL,
    major_id INT NOT NULL,
    created_at DATETIME default current_timestamp,
    updated_at DATETIME default current_timestamp ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE skillset.groups_of_members (
    member_id VARCHAR(255) NOT NULL,
    group_id INT NOT NULL,
    PRIMARY KEY(member_id, group_id)
);
CREATE TABLE skillset.tech_areas_of_members (
    member_id VARCHAR(255) NOT NULL,
    tech_area_id INT NOT NULL,
    PRIMARY KEY(member_id, tech_area_id)
);
CREATE TABLE skillset.techs_of_members (
    member_id VARCHAR(255) NOT NULL,
    tech_id INT NOT NULL,
    PRIMARY KEY(member_id, tech_id)
);
CREATE TABLE skillset.products_of_members (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    member_id VARCHAR(255) NOT NULL,
    overview VARCHAR(255),
    url VARCHAR(255)
);