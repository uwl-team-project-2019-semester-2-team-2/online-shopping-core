SELECT product.id, product_line.id, product_line.name, product.colour, brand.name AS brand_name, product_line.description
	FROM product
		INNER JOIN product_line 
			ON product.product_line_id = product_line.id
		INNER JOIN brand
			ON product_line.brand_id = brand.id
		WHERE product.id LIKE 2
        LIMIT 1;