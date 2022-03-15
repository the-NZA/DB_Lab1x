CREATE TABLE IF NOT EXISTS users(
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  username text UNIQUE NOT NULL,
  passwd text NOT NULL,
  email text UNIQUE,
  is_admin boolean NOT NULL DEFAULT(false),
  deleted boolean NOT NULL DEFAULT(false),
);