INSERT INTO department (name) VALUES ("Fresh & Chilled");
INSERT INTO department (name) VALUES ("Bakery");
INSERT INTO department (name) VALUES ("Frozen");
INSERT INTO department (name) VALUES ("Beer, Wine & Spirits");
INSERT INTO department (name) VALUES ("Household");
INSERT INTO department (name) VALUES ("Toiletries");
INSERT INTO department (name) VALUES ("Health & Medicine");
INSERT INTO department (name) VALUES ("Pet");
INSERT INTO department (name) VALUES ("Drinks");
INSERT INTO department (name) VALUES ("New");

INSERT INTO brand (name, description) VALUES ("The Happy Egg Co.", "");
INSERT INTO product (name, brand_id, price, item_quantity, item_quantity_postfix, description) VALUES ("The Happy Egg Co. 6 Free Range Eggs Medium", 1, 1.30, 6, " per pack", "");
INSERT INTO product (name, brand_id, price, item_quantity, item_quantity_postfix, description) VALUES ("The Happy Egg Co. 12 Free Range Eggs Large", 1, 2.70, 12, " per pack", "");


INSERT INTO customer (first_name, second_name, email_address, phone_number) VALUES ("Sebastian", "Mills", "SebastianMills@jourrapide.com", "07922930005");
INSERT INTO address (customer_id, address_line_1, address_line_2, town_city, county, post_code) VALUES (1, "542 Tom Brook", "", "Lake Layla", "Lancastershire", "LA8 9RA");
