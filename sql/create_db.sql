CREATE TABLE department (
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(30),
	parent_id INT,
    url VARCHAR(30),
	PRIMARY KEY (id)
);

CREATE TABLE dietary (
	id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(30),
    url VARCHAR(30),
    filter boolean,
    PRIMARY KEY (id)
);

CREATE TABLE brand (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    description VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE product (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(60),
    brand_id INT,
    department_id INT,
	price DECIMAL(13, 2),
    item_quantity INT,
    item_quantity_postfix VARCHAR(10),
	description VARCHAR(5000),
    PRIMARY KEY (id),
    FOREIGN KEY (brand_id) REFERENCES brand (id),
    FOREIGN KEY (department_id) REFERENCES department (id)
);

CREATE TABLE product_dietary (
	id INT NOT NULL AUTO_INCREMENT,
    product_id INT,
    dietary_id INT,
	PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (dietary_id) REFERENCES dietary (id)
);

CREATE TABLE product_image (
	id INT NOT NULL AUTO_INCREMENT,
    product_id INT,
    url VARCHAR(255),
	PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES product (id)
);

CREATE TABLE product_image_cover (
	product_id INT,
    product_image_id INT,
    PRIMARY KEY (product_id, product_image_id),
	UNIQUE (product_id, product_image_id),
	FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (product_image_id) REFERENCES product_image (id)
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
);

CREATE TABLE basket (
	id INT NOT NULL AUTO_INCREMENT,
    customer_id INT,
    product_id INT,
    quantity INT,
	PRIMARY KEY (id),
	FOREIGN KEY (product_id) REFERENCES product (id),
	FOREIGN KEY (customer_id) REFERENCES customer (id)
);