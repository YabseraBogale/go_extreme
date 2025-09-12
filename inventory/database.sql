create table Employee(
    id INT PRIMARY KEY AUTO_INCREMENT,
    firstname varchar(30) not null,
    middlename varchar(30) not null,
    lastname varchar(30) not null,
    phonenumber int not null,
    email text not null
)

CREATE TABLE categories (
    category_id  INT PRIMARY KEY AUTO_INCREMENT,
    name         VARCHAR(100) NOT NULL
);

CREATE TABLE warehouses (
    warehouse_id   INT PRIMARY KEY AUTO_INCREMENT,
    name           VARCHAR(30) NOT NULL,
    location       VARCHAR(30)
);

create table Item(
    item_id INT PRIMARY KEY AUTO_INCREMENT,
    item_name text not null,
    item_decrption text not null,
    item_price float not null,
    item_quantity int not null,
    category_id references categories,
    warehouse_id references warehouses

)