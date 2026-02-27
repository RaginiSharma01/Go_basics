create table if not exists order(
    order_id serial primary key,
    user_id int not null,
    amount numeric(12,2) not null(amount>0),
    created_at timestamptz default now()

    forigen
);