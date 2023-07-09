CREATE TABLE artists (
    artist_id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    genre TEXT NOT NULL,
    formation DATETIME NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_artists_created ON artists(created);

CREATE TABLE albums (
    album_id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    artist_id INTEGER NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artists(artist_id),
    title VARCHAR(100) NOT NULL,
    genre TEXT NOT NULL,
    released DATETIME NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_albums_created ON albums(created);

CREATE TABLE tracks(
    track_id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    artist_id INTEGER NOT NULL,
    album_id INTEGER,
    FOREIGN KEY (artist_id) REFERENCES artists(artist_id),
    FOREIGN KEY (album_id) REFERENCES albums(album_id)
    title VARCHAR(100) NOT NULL,
    genre TEXT NOT NULL,
    duration INTEGER NOT NULL,
    created DATETIME NOT NULL
);

CREATE INDEX idx_tracks_created ON tracks(created);
