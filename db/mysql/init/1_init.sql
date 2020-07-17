CREATE DATABASE IF NOT EXISTS skillset;
USE skillset;
CREATE TABLE skillset.members (
    id VARCHAR NOT NULL PRIMARY KEY,
    image VARCHAR,
    name VARCHAR NOT NULL,
    message VARCHAR,
    slack VARCHAR NOT NULL,
    twitter VARCHAR,
    github VARCHAR,
    grade_id INT NOT NULL,
    major_id INT NOT NULL,
    created_at DATETIME default current_timestamp,
    updated_at DATETIME default current_timestamp ON UPDATE CURRENT_TIMESTAMP,
);
INSERT INTO members (
        id,
        image,
        name,
        message,
        slack,
        twitter,
        github,
        grade_id,
        major_id,
    )
VALUES (
        "550e8400-e29b-41d4-a716-446655440000",
        "https://3.bp.blogspot.com/-A_AqujoUstw/Uj_279AEvZI/AAAAAAAAYEk/C7AU-QH0IiE/s800/monogatari_sanbikinokobuta.png"
        "三匹の子豚",
        "message",
        "ぶたさん",
        "butasan",
        "butasan",
        1,
        1,
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
    technology_id INT NOT NULL PRIMARY KEY,
    CONSTRAINT fk_member_id FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT fk_technology_id FOREIGN KEY (technology_id) REFERENCES technologies (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
CREATE TABLE skillset.products_of_members (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    member_id VARCHAR NOT NULL,
    overview VARCHAR,
    url VARCHAR,
    CONSTRAINT fk_member_id FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);