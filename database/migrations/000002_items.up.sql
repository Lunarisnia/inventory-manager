CREATE TABLE items (
  id serial primary key,
  name text not null,
  image text not null,
  created_at bigint not null,
  updated_at bigint not null
);
