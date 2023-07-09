
-- Dummy data
INSERT INTO artists (name, genre, formation, created, expires) VALUES (
    'Slipknot',
    'Nu Metal',
    '1995-9-9',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO artists (name, genre, formation, created, expires) VALUES (
    'Korn',
    'Nu Metal',
    '1990-4-6',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO albums (artist_id, title, genre, released, created, expires) VALUES (
    1,
    'Slipknot - Self Titled',
    'Nu Metal',
    '1999-9-9',
    UTC_TIMESTAMP(),
    (SELECT created FROM artists WHERE artist_id = 1)
);

INSERT INTO tracks (artist_id, album_id, title, genre, duration, created, expires) VALUES (
    1,
    1,
    'Dead Memories',
    'Metal',
    245,
    UTC_TIMESTAMP(),
    (SELECT created FROM artists WHERE artist_id = 1)
);

INSERT INTO tracks (artist_id, album_id, title, genre, duration, created, expires) VALUES (
    1,
    NULL,
    'All Out Live',
    'Metal',
    190,
    UTC_TIMESTAMP(),
    (SELECT created FROM artists WHERE artist_id = 1)
);