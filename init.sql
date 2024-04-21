-- Create buyer table
create table buyer (
    id uuid primary key default uuid_generate_v4(),
    email varchar(255) unique not null,
    name varchar(255) not null,
    password varchar(255) not null,
    alamat_pengiriman text not null
);

-- Create seller table
create table seller (
    id uuid primary key default uuid_generate_v4(),
    email varchar(255) unique not null,
    name varchar(255) not null,
    password varchar(255) not null,
    alamat_pickup text not null
);

-- Create product table
create table product (
    id uuid primary key default uuid_generate_v4(),
    product_name varchar(255) not null,
    description text,
    price decimal(10, 2) not null,
    seller_id uuid,
    foreign key (seller_id) references seller(id)
);

-- Create order table
create table "order" (
    id uuid primary key default uuid_generate_v4(),
    buyer_id uuid,
    seller_id uuid,
    delivery_source_address text not null,
    delivery_destination_address text not null,
    items text,
    quantity int,
    price decimal(10, 2),
    total_price decimal(10, 2),
    status varchar(50) default 'pending',
    foreign key (buyer_id) references buyer(id),
    foreign key (seller_id) references seller(id)
);

-- Create the pgcrypto extension if not exists
create extension if not exists pgcrypto;

-- Insert a seller into the table
insert into seller (email, name, password, alamat_pickup)
values ('rio@gmail.com', 'rio', crypt('123', gen_salt('bf')), '123 Main Street, Anytown, USA');

-- Insert a buyer into the table
insert into buyer (email, name, password, alamat_pengiriman)
values ('rich@gmail.com', 'richard', crypt('456', gen_salt('bf')), '456 Elm Street, Anytown, USA');