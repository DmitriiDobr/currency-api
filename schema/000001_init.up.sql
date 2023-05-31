CREATE TABLE currency_exchange(
    id serial not null unique,
    currency_from varchar(255) not null,
    currency_to varchar(255)  not null,
    exchange_rate float not null,
    updated_at timestamp not null
)