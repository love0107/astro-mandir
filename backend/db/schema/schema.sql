CREATE TABLE IF NOT EXISTS panchang (
    date TEXT PRIMARY KEY,
    vrat TEXT,
    tithi TEXT,
    nakshatra TEXT,
    sunrise TEXT,
    sunset TEXT,
    moonrise TEXT,
    yoga TEXT,
    muhurat TEXT,
    festival TEXT
);

CREATE TABLE IF NOT EXISTS bhajans (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    youtube_id TEXT NOT NULL,
    festival_type TEXT,
    rashi TEXT,
    scheduled_date TEXT
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    phone TEXT UNIQUE,
    rashi TEXT,
    name TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS kundali_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    dob TEXT,
    tob TEXT,
    place TEXT,
    rashi TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);