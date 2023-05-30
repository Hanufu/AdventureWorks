create table products(
    Id                    serial primary key  ,
    ProductKey            int  ,
	ProductSubcategoryKey int,
	ProductSKU			  varchar,
	ProductName           varchar,
	ModelName             varchar ,
	ProductDescription    varchar ,
	ProductColor          varchar ,
	ProductSize           varchar ,
	ProductStyle          varchar ,
	ProductCost           float ,
	ProductPrice          float 
);

INSERT INTO products (productKey, productSubcategoryKey, productSKU, productName, modelName, productDescription, productColor, productSize, productStyle, productCost, productPrice) VALUES 
(214, 31, 'HL-U509-R', 'Sport-100 Helmet, Red', 'Sport-100', 'Universal fit, well-vented, lightweight, snap-on visor.', 'Red', 'U', NULL, 13.0863, 34.99),
(215, 31, 'HL-U509', 'Sport-100 Helmet, Black', 'Sport-100', 'Universal fit, well-vented, lightweight, snap-on visor.', 'Black', 'U', NULL, 12.0278, 33.6442),
(218, 23, 'SO-B909-M', 'Mountain Bike Socks, M', 'Mountain Bike Socks', 'Combination of natural and synthetic fibers stays dry and provides just the right cushioning.', 'White', 'M', 'U', 3.3963, 9.5),
(219, 23, 'SO-B909-L', 'Mountain Bike Socks, L', 'Mountain Bike Socks', 'Combination of natural and synthetic fibers stays dry and provides just the right cushioning.', 'White', 'L', 'U', 3.3963, 9.5),
(220, 31, 'HL-U509-B', 'Sport-100 Helmet, Blue', 'Sport-100', 'Universal fit, well-vented, lightweight, snap-on visor.', 'Blue', 'U', NULL, 12.0278, 33.6442),
(223, 19, 'CA-1098', 'AWC Logo Cap', 'Cycling Cap', 'Traditional style with a flip-up brim; one-size fits all.', 'Multi', 'U', NULL, 5.7052, 8.6442),
(226, 21, 'LJ-0192-S', 'Long-Sleeve Logo Jersey, S', 'Long-Sleeve Logo Jersey', 'Unisex long-sleeve AWC logo microfiber cycling jersey', 'Multi', 'S', 'U', 31.7244, 48.0673),
(229, 21, 'LJ-0192-M', 'Long-Sleeve Logo Jersey, M', 'Long-Sleeve Logo Jersey', 'Unisex long-sleeve AWC logo microfiber cycling jersey', 'Multi', 'M', 'U', 31.7244, 48.0673),
(232, 21, 'LJ-0192-L', 'Long-Sleeve Logo Jersey, L', 'Long-Sleeve Logo Jersey', 'Unisex long-sleeve AWC logo microfiber cycling jersey', 'Multi', 'L', 'U', 31.7244, 48.0673),
(235, 21, 'LJ-0192-X', 'Long-Sleeve Logo Jersey, XL', 'Long-Sleeve Logo Jersey', 'Unisex long-sleeve AWC logo microfiber cycling jersey', 'Multi', 'XL', 'U', 31.7244, 48.0673),
(238, 14, 'FR-R92R-62', 'HL Road Frame - Red, 62', 'HL Road Frame', 'Our lightest and best quality aluminum frame made from the newest alloy; it is welded and heat-treated for strength. Our innovative design results in maximum comfort and performance.', 'Red', 'U', 62, 747.9682, 1263.4598),
(241, 14, 'FR-R92R-44', 'HL Road Frame - Red, 44', 'HL Road Frame', 'Our lightest and best quality aluminum frame made from the newest alloy; it is welded and heat-treated for strength. Our innovative design results in maximum comfort and performance.', 'Red', 'U', 44, 747.9682, 1263.4598),
(244, 14, 'FR-R92R-48', 'HL Road Frame - Red, 48', 'HL Road Frame', 'Our lightest and best quality aluminum frame made from the newest alloy; it is welded and heat-treated for strength. Our innovative design results in maximum comfort and performance.', 'Red', 'U', 48, 747.9682, 1263.4598),
(247, 14, 'FR-R92R-52', 'HL Road Frame - Red, 52', 'HL Road Frame', 'Our lightest and best quality aluminum frame made from the newest alloy; it is welded and heat-treated for strength. Our innovative design results in maximum comfort and performance.', 'Red', 'U', 52, 747.9682, 1263.4598),
(250, 14, 'FR-R92R-56', 'HL Road Frame - Red, 56', 'HL Road Frame', 'Our lightest and best quality aluminum frame made from the newest alloy; it is welded and heat-treated for strength. Our innovative design results in maximum comfort and performance.', 'Red', 'U', 56, 747.9682, 1263.4598),
(253, 14, 'FR-R38B-58', 'LL Road Frame - Black, 58', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Black', 'U', 58, 176.1997, 297.6346),
(256, 14, 'FR-R38B-60', 'LL Road Frame - Black, 60', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Black', 'U', 60, 176.1997, 297.6346),
(259, 14, 'FR-R38B-62', 'LL Road Frame - Black, 62', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Black', 'U', 62, 176.1997, 297.6346),
(262, 14, 'FR-R38R-44', 'LL Road Frame - Red, 44', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Red', 'U', 44, 181.4857, 306.5636),
(264, 14, 'FR-R38R-48', 'LL Road Frame - Red, 48', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Red', 'U', 48, 181.4857, 306.5636),
(266, 14, 'FR-R38R-52', 'LL Road Frame - Red, 52', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Red', 'U', 52, 181.4857, 306.5636),
(268, 14, 'FR-R38R-58', 'LL Road Frame - Red, 58', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Red', 'U', 58, 181.4857, 306.5636),
(270, 14, 'FR-R38R-60', 'LL Road Frame - Red, 60', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Red', 'U', 60, 181.4857, 306.5636),
(272, 14, 'FR-R38R-62', 'LL Road Frame - Red, 62', 'LL Road Frame', 'The LL Frame provides a safe comfortable ride, while offering superior bump absorption in a value-priced aluminum frame.', 'Red', 'U', 62, 181.4857, 306.5636);

