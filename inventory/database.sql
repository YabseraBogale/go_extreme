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
    category_id int references categories,
    warehouse_id int references warehouses

)

CREATE TABLE stock_movements (
    movement_id   INT PRIMARY KEY AUTO_INCREMENT,
    item_id    INT references Item,
    warehouse_id  INT references warehouses,
    change_qty    INT NOT NULL,  -- positive for add, negative for removal
    movement_type ENUM('purchase','sale','transfer','adjustment') NOT NULL,
    reference_id  INT NOT NULL, -- link to PO, SO, or manual adjustment
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   
);
