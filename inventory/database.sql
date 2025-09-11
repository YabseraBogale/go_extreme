create table Employee(
    id INT PRIMARY KEY AUTO_INCREMENT,
    firstname varchar(30) not null,
    middlename varchar(30) not null,
    lastname varchar(30) not null,
    phonenumber int not null,
    email text not null
)

create table Item(
    item_id INT PRIMARY KEY AUTO_INCREMENT,
    item_name text not null,
    item_decrption text not null,
    item_price float not null,
    item_quantity int not null,
)