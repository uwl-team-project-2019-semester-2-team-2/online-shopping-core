CREATE TABLE brand (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    description VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE product_line (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    brand_id INT,
    description VARCHAR(5000),
    PRIMARY KEY (id),
    FOREIGN KEY (brand_id) REFERENCES brand (id)
);

CREATE TABLE product (
    id INT NOT NULL AUTO_INCREMENT,
    product_line_id INT,
    colour VARCHAR(255),
    PRIMARY KEY (id),
    FOREIGN KEY (product_line_id) REFERENCES product_line (id)
);

CREATE TABLE stock (
    id INT NOT NULL AUTO_INCREMENT,
    product_id INT,
    size VARCHAR(255),
    quantity INT,
    PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES product (id)
);

CREATE TABLE pictures (
    id INT NOT NULL AUTO_INCREMENT,
    url VARCHAR(255),
    product_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES product (id)
);

CREATE TABLE customer (
	id INT NOT NULL AUTO_INCREMENT,
	first_name VARCHAR(30),
	second_name VARCHAR(30),
	email_address VARCHAR(30),
	hash VARCHAR(255),
	phone_number VARCHAR(20),
	PRIMARY KEY (id)
);

CREATE TABLE address (
	id INT NOT NULL AUTO_INCREMENT,
	customer_id INT,
	address_line_1 VARCHAR(30),
	address_line_2 VARCHAR(30),
	town_city VARCHAR(30),
	county VARCHAR(30),
	post_code VARCHAR(10),
	PRIMARY KEY (id),
	FOREIGN KEY (customer_id) REFERENCES customer (id)
);

CREATE TABLE orders (
	id INT NOT NULL AUTO_INCREMENT,
	customer_id INT,
	date_placed VARCHAR(255),
	order_status VARCHAR(255),
	PRIMARY KEY (id),
	FOREIGN KEY (customer_id) REFERENCES customer (id)
);

CREATE TABLE order_items (
	id INT NOT NULL AUTO_INCREMENT,
	order_id INT,
	product_id INT,
	quantity INT,
	PRIMARY KEY (id),
	FOREIGN KEY (order_id) REFERENCES orders (id),
	FOREIGN KEY (product_id) REFERENCES product (id)
)