USE tugas_day15;

-- nomor 1
SELECT *
FROM transactions tr
WHERE tr.user_id=1
UNION
SELECT *
FROM transactions tr
WHERE tr.user_id=2;

-- nomor 2
SELECT SUM(tr.total_price)
FROM transactions tr
JOIN users us ON tr.user_id = us.id
WHERE us.id = 1;

-- nomor 3
SELECT td.product_id, COUNT(td.transaction_id)
FROM transactions tr
JOIN transaction_details td ON tr.id_transaction = td.transaction_id
WHERE td.product_id = 2
GROUP BY td.product_id;

-- nomor 4
SELECT pr.*, pt.nama_product_types
FROM products pr
JOIN product_types pt ON pr.product_type_id = pt.id_product_types;

-- nomor 5
SELECT tr.*, pr.`name`, us.nama
FROM transactions tr
RIGHT JOIN transaction_details dt ON tr.id_transaction = dt.transaction_id
RIGHT JOIN products pr ON pr.id_product = dt.product_id
RIGHT JOIN users us ON us.id = tr.user_id;

INSERT INTO products (id_product,product_type_id, operator_id,`code`,`name`,`status`,price,created_at, updated_at)
VALUES
	(NULL,'1',NULL,'prd776','Susu', 2, 7500, NULL, NULL),
	(NULL,'1',NULL,'prd778','Jus', 1, 7000, NULL, NULL),
	(NULL,'1',NULL,'prd777','Kopi', 2, 6000, NULL, NULL);
	
-- nomor 8
SELECT pr.`name`
FROM products pr
WHERE NOT EXISTS(
	SELECT td.transaction_id
	FROM transaction_details td
	JOIN products ON td.product_id = pr.id_product
);
