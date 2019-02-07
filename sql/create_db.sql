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

SELECT * FROM product INNER JOIN product_line ON product.product_line_id = product.id WHERE id LIKE 1;