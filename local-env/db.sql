-- Create user for connector debezium
CREATE USER scinventory IDENTIFIED BY 'dbzpass';
GRANT RELOAD, SELECT, SHOW DATABASES, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO scinventory;

CREATE DATABASE inventorydb;
CREATE USER inventoryusr IDENTIFIED BY 'dbpass';
GRANT ALL ON inventorydb.* TO inventoryusr;

USE inventorydb;

CREATE TABLE menu (
    id BINARY(16),
    name VARCHAR(50),
    type CHAR(2),
    description TEXT,
    created_by VARCHAR(50),
    created_at TIMESTAMP,
    updated_by VARCHAR(50),
    updated_at TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE kitchen (
    id BINARY(16),
    menu_id BINARY(16),
    status CHAR(2),
    created_by VARCHAR(50),
    created_at TIMESTAMP,
    updated_by VARCHAR(50),
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    CONSTRAINT fk_kitchen_menu
                     FOREIGN KEY(menu_id) REFERENCES menu(id)
);

CREATE INDEX idx_kitchen_status ON kitchen(status);