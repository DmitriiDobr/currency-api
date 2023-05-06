CREATE TABLE currency_exchange(
    id serial not null unique,
    CurrencyFrom varchar(255) not null,
    CurrencyTo varchar(255)  not null,
    ExchangeRate float not null,
    UpdatedAt timestamp not null
)