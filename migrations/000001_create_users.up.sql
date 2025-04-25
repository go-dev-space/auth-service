CREATE TABLE IF NOT EXISTS users(
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE CHECK(char_length(username)>2),
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL CHECK(char_length(password)>5),
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  refresh_token TEXT CHECK (char_length(refresh_token) > 20)
);