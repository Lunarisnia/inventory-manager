create table borrow_lists(
  id serial primary key,
  user_id int not null,
  item_id int not null,
  borrow_at bigint not null,
  returned_at bigint DEFAULT NULL,
  created_at bigint not null,
  updated_at bigint not null,
  foreign key(item_id) references items(id),
  foreign key(user_id) references users(id)
);
