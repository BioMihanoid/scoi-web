CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL,
                                     first_name TEXT,
                                     second_name TEXT,
                                     email TEXT,
                                     hashed_password TEXT,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);