create table product (
    id serial primary key,
    title text not null,
    description text not null
);

create table product_item (
    id serial primary key,
    sku text unique not null,
    material text not null,
    product_id serial not null, 
    FOREIGN KEY(product_id)
        REFERENCES product(id)
);

create table product_price (
    id serial primary key,
    product_id serial not null, 
    FOREIGN KEY(product_id)
        REFERENCES product(id),
    price serial not null
);