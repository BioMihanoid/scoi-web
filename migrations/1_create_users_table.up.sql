CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL,
                                     first_name TEXT default '',
                                     second_name TEXT default '',
                                     email TEXT UNIQUE NOT NULL,
                                     hashed_password TEXT NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);