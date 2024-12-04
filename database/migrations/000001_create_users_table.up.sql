CREATE TABLE users (
  id    SERIAL PRIMARY KEY,
  name  text   NOT NULL,
  nis   text   NOT NULL,
  created_at bigint NOT NULL,
  updated_at bigint NOT NULL
);

