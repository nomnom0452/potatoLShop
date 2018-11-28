DROP DATABASE absurd;

CREATE DATABASE absurd;

USE absurd

CREATE TABLE items (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Quantity int  NOT NULL,
    Price int NOT NULL,
    Story VARCHAR(255) NOT NULL
);

CREATE TABLE baskets (
    ID SERIAL PRIMARY KEY,
    ItemId INT NOT NULL REFERENCES items(ID),
    UserId INT NOT NULL REFERENCES customers(ID)
);

CREATE TABLE customers (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Age int NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL
);

CREATE TABLE sellers (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Age int NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL
);

CREATE TABLE transactions (
    ID SERIAL PRIMARY KEY,
    customerID INT REFERENCES customers(ID),
    sellerID INT REFERENCES sellers(ID),
    transactionDate DATE NOT NULL,
    status INT
);

CREATE TABLE transactionDetails (
    ID SERIAL PRIMARY KEY,
    transactionID INT REFERENCES transactions(ID),
    itemID INT REFERENCES items(ID),
    itemQuantity INT
);

CREATE TABLE sessions (
    ID SERIAL PRIMARY KEY,
    Uuid VARCHAR(255),
    Email VARCHAR(255),
    UserId INT,
    CreateAt TIMESTAMP
);



INSERT INTO items(Name, Quantity, Price, Story) VALUES
    ("Item 1", 30, 20000, "description 1"),
    ("item 2", 25, 15000, "description 2"),
    ("item 3", 10, 100000, "description 3");

INSERT INTO transactions(customerID, sellerID, transactionDate, status) VALUES
    (2, 1, "1998/09/18", 1),
    (2, 2, "1991/06/17", 1),
    (3, 1, "1992/04/14", 1),
    (3, 3, "1995/03/11", 1),
    (1, 2, "1994/02/28", 1),
    (1, 3, "1993/01/08", 0);

INSERT INTO transactionDetails(transactionID, itemID, itemQuantity) VALUES 
    (1, 1, 2),
    (1, 2, 2),
    (1, 3, 2),
    (2, 1, 3),
    (2, 1, 1),
    (2, 1, 2),
    (3, 2, 1),
    (3, 2, 1),
    (3, 2, 1);
