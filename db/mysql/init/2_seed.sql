USE skillset;
INSERT INTO members (
        id,
        image,
        name,
        message,
        slack,
        twitter,
        github,
        grade_id,
        major_id
    )
VALUES (
        "11111111-1111-1111-1111-111111111111",
        "https://3.bp.blogspot.com/-A_AqujoUstw/Uj_279AEvZI/AAAAAAAAYEk/C7AU-QH0IiE/s800/monogatari_sanbikinokobuta.png",
        "三匹の子豚たち",
        "message",
        "ぶたさん",
        "butasan",
        "butasan",
        1,
        1
    ),
    (
        "22222222-2222-2222-2222-222222222222",
        "https://2.bp.blogspot.com/-18rIpu5Yr7c/VnE3ZiXdvGI/AAAAAAAA11U/s0F8cwnqo1M/s800/character_hitsuji_ookami.png",
        "おおかみ",
        "message",
        "羊さんだよ",
        "gaogao",
        "harahetta",
        1,
        1
    );
INSERT INTO groups_of_members(member_id, group_id)
VALUES ("11111111-1111-1111-1111-111111111111", 0),
    ("11111111-1111-1111-1111-111111111111", 1),
    ("22222222-2222-2222-2222-222222222222", 0)