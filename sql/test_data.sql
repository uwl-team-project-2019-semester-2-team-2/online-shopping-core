INSERT INTO department (name, parent_id, url) VALUES ("Fresh & Chilled", 0, "fresh-chilled");
INSERT INTO department (name, parent_id, url) VALUES ("Bakery", 0, "bakery");
INSERT INTO department (name, parent_id, url) VALUES ("Frozen", 0, "frozen");
INSERT INTO department (name, parent_id, url) VALUES ("Beer, Wine & Spirits", 0, "beer-wine-spirits");
INSERT INTO department (name, parent_id, url) VALUES ("Household", 0, "household");
INSERT INTO department (name, parent_id, url) VALUES ("Toiletries", 0, "toiletries");
INSERT INTO department (name, parent_id, url) VALUES ("Health & Medicine", 0, "health-medicine");
INSERT INTO department (name, parent_id, url) VALUES ("Pet", 0, "pet");
INSERT INTO department (name, parent_id, url) VALUES ("Drinks", 0, "drinks");
INSERT INTO department (name, parent_id, url) VALUES ("New", 0, "new");

INSERT INTO department (name, parent_id, url) VALUES ("Milk, Butter & Eggs", 1, "milk-butter-eggs");
INSERT INTO department (name, parent_id, url) VALUES ("Free Range Eggs", 11, "free-range-eggs");

INSERT INTO dietary (name, url) VALUES ("No Egg", "eggs");
INSERT INTO dietary (name, url) VALUES ("No Milk", "milk");
INSERT INTO dietary (name, url) VALUES ("No Gluten", "gluten");
INSERT INTO dietary (name, url) VALUES ("Vegeterian", "vegeterian");
INSERT INTO dietary (name, url) VALUES ("Vegan", "vegan");
INSERT INTO dietary (name, url) VALUES ("Organic", "organic");

INSERT INTO brand (name, description) VALUES ("The Happy Egg Co.", "");
INSERT INTO product (name, brand_id, department_id, price, item_quantity, item_quantity_postfix, description) VALUES ("The Happy Egg Co. 6 Free Range Eggs Medium", 1, 11, 1.30, 6, " per pack", "");
INSERT INTO product_image (product_id, url) VALUES (1, "/product_images/116487011_0_640x640.jpg");
INSERT INTO product_image_cover (product_id, product_image_id) VALUES (1, 1);
INSERT INTO product_dietary (product_id, dietary_id) VALUES (1, 1);

INSERT INTO product (name, brand_id, department_id, price, item_quantity, item_quantity_postfix, description) VALUES ("The Happy Egg Co. 12 Free Range Eggs Large", 1, 11, 2.70, 12, " per pack", "");
INSERT INTO product_image (product_id, url) VALUES (2, "/product_images/394000011_0_640x640.jpg");
INSERT INTO product_image_cover (product_id, product_image_id) VALUES (2, 2);
INSERT INTO product_dietary (product_id, dietary_id) VALUES (2, 1);

INSERT INTO customer (first_name, second_name, email_address, phone_number) VALUES ("Sebastian", "Mills", "SebastianMills@jourrapide.com", "07922930005");
INSERT INTO address (customer_id, address_line_1, address_line_2, town_city, county, post_code) VALUES (1, "542 Tom Brook", "", "Lake Layla", "Lancastershire", "LA8 9RA");
