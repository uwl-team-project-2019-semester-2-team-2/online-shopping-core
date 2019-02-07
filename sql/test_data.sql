INSERT INTO brand (name, description) VALUES ("Nike", "");

INSERT INTO product_line (name, brand_id, description) VALUES ("Nike Air Max 90", 1, "The Nike Air Max 90 released in 1990 and is considered to be the second flagship sneaker of the Air Max legacy. The most popular colorway is the 'Infrared', which was the original colorway in 1990. The Air Max 90 was designed by legendary Nike architect Tinker Hatfield.");
INSERT INTO product (product_line_id, colour) VALUES (1, "Black/White");

INSERT INTO stock (product_id, size, quantity) VALUES (1, "6", 10);
INSERT INTO stock (product_id, size, quantity) VALUES (1, "7", 3);
INSERT INTO stock (product_id, size, quantity) VALUES (1, "8", 12);
INSERT INTO stock (product_id, size, quantity) VALUES (1, "9", 16);
INSERT INTO stock (product_id, size, quantity) VALUES (1, "10", 5);
INSERT INTO stock (product_id, size, quantity) VALUES (1, "11", 2);
INSERT INTO stock (product_id, size, quantity) VALUES (1, "12", 0);


INSERT INTO product (product_line_id, colour) VALUES (1, "Grey");
INSERT INTO stock (product_id, size, quantity) VALUES (2, "6", 13);
INSERT INTO stock (product_id, size, quantity) VALUES (2, "7", 5);
INSERT INTO stock (product_id, size, quantity) VALUES (2, "8", 5);
INSERT INTO stock (product_id, size, quantity) VALUES (2, "9", 12);
INSERT INTO stock (product_id, size, quantity) VALUES (2, "10", 15);
INSERT INTO stock (product_id, size, quantity) VALUES (2, "11", 1);
INSERT INTO stock (product_id, size, quantity) VALUES (2, "12", 6);