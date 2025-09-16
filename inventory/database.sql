create table if not EXISTS Employee(
    employee_id INTEGER PRIMARY KEY AUTOINCREMENT,
    firstname varchar(30) not null,
    middlename varchar(30) not null,
    lastname varchar(30) not null,
    phonenumber int not null,
    email text not null,
    
);



create table if not EXISTS Item(
    item_id INTEGER PRIMARY KEY AUTOINCREMENT,
    employee_id INTEGER references Employee,
    item_name text not null,
    item_decrption text not null,
    item_price float not null,
    item_quantity int not null,
    category Text not null,
    location Text not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP 
);
