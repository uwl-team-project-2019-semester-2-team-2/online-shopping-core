SELECT product.id, product_line.name, product.colour FROM product 
	INNER JOIN product_line
		ON product.product_line_id = product_line.id
	WHERE product_line.name LIKE "%nike%"
    ORDER BY product.id
    LIMIT 0,25;