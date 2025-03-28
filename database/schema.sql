CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    nick_name TEXT UNIQUE NOT NULL,
    gender TEXT NOT NULL,
    age INTEGER NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    session_token TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    profile_image TEXT
);

CREATE TABLE IF NOT EXISTS posts (
  id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  post_title TEXT,
  body TEXT,
  parent_id INTEGER DEFAULT NULL,
  created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  post_status TEXT DEFAULT 'visible',
  media_url TEXT DEFAULT '', 
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (parent_id) REFERENCES posts (id)
);

CREATE TABLE IF NOT EXISTS post_categories (
  id INTEGER PRIMARY KEY,
  post_id INTEGER NOT NULL,
  category TEXT,
  FOREIGN KEY (post_id) REFERENCES posts (id)
);

CREATE TABLE IF NOT EXISTS reactions (
  id INTEGER PRIMARY KEY,
  reaction TEXT,
  reaction_status TEXT DEFAULT 'clicked',
  user_id INTEGER NOT NULL,
  post_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (post_id) REFERENCES posts (id)
);

CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT NOT NULL,
    sender TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
