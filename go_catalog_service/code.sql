sudo -u postgres psql -d uzum_market

INSERT INTO category (id, slug, name_uz, name_ru, name_en, active, order_no, parent_id, created_at, updated_at, deleted_at) VALUES
('a1a0234b-2f97-4e7b-bb07-c3c9d5b00111', 'electronics', 'Elektronika', 'Электроника', 'Electronics', true, 1, NULL, NOW(), NULL, NULL),
('a1a0234b-2f97-4e7b-bb07-c3c9d5b00222', 'clothing', 'Kiyim-kechak', 'Одежда', 'Clothing', true, 2,'a1a0234b-2f97-4e7b-bb07-c3c9d5b00111', NOW(), NULL, NULL),
('a1a0234b-2f97-4e7b-bb07-c3c9d5b00333', 'books', 'Kitoblar', 'Книги', 'Books', true, 3,'a1a0234b-2f97-4e7b-bb07-c3c9d5b00111', NOW(), NULL, NULL),
('a1a0234b-2f97-4e7b-bb07-c3c9d5b00444', 'furniture', 'Mebel', 'Мебель', 'Furniture', true, 4,'a1a0234b-2f97-4e7b-bb07-c3c9d5b00111', NOW(), NULL, NULL),
('a1a0234b-2f97-4e7b-bb07-c3c9d5b00555', 'toys', 'O‘yinchoqlar', 'Игрушки', 'Toys', true, 5,'a1a0234b-2f97-4e7b-bb07-c3c9d5b00111', NOW(), NULL, NULL);

INSERT INTO product (id, slug, name_uz, name_ru, name_en, description_uz, description_ru, description_en, active, order_no, in_price, out_price, left_count, discount_percent, image, created_at, updated_at, deleted_at) VALUES
('66666666-6666-6666-6666-666666666666', 'phone', 'Telefon', 'Телефон', 'Phone', 'Yangi model', 'Новая модель', 'New model', true, 1, 200, 250, 50, 10, ARRAY['img1.jpg', 'img2.jpg'], NOW(), NULL, NULL),
('77777777-7777-7777-7777-777777777777', 'shirt', 'Ko\ylak', 'Рубашка', 'Shirt', 'Yangi uslub', 'Новый стиль', 'New style', true, 2, 20, 30, 100, 5, ARRAY['img3.jpg', 'img4.jpg'], NOW(), NULL, NULL),
('88888888-8888-8888-8888-888888888888', 'book', 'Kitob', 'Книга', 'Book', 'Yangi adabiyot', 'Новая литература', 'New literature', true, 3, 10, 15, 200, 0, ARRAY['img5.jpg', 'img6.jpg'], NOW(), NULL, NULL),
('99999999-9999-9999-9999-999999999999', 'table', 'Stol', 'Стол', 'Table', 'Yangi model', 'Новая модель', 'New model', true, 4, 50, 70, 30, 15, ARRAY['img7.jpg', 'img8.jpg'], NOW(), NULL, NULL),
('00000000-0000-0000-0000-000000000000', 'teddy-bear', 'O\yinchoq ayiqlar', 'Игрушечные мишки', 'Teddy Bears', 'Yumshoq ayiqcha', 'Мягкая игрушка', 'Soft toy', true, 5, 15, 20, 150, 0, ARRAY['img9.jpg', 'img10.jpg'], NOW(), NULL, NULL);

INSERT INTO product_categories (id, product_id, category_id) VALUES
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '66666666-6666-6666-6666-666666666666', 'a1a0234b-2f97-4e7b-bb07-c3c9d5b00222'),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '77777777-7777-7777-7777-777777777777', 'a1a0234b-2f97-4e7b-bb07-c3c9d5b00222'),
('cccccccc-cccc-cccc-cccc-cccccccccccc', '88888888-8888-8888-8888-888888888888', 'a1a0234b-2f97-4e7b-bb07-c3c9d5b00333'),
('dddddddd-dddd-dddd-dddd-dddddddddddd', '99999999-9999-9999-9999-999999999999', 'a1a0234b-2f97-4e7b-bb07-c3c9d5b00444'),
('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '00000000-0000-0000-0000-000000000000', 'a1a0234b-2f97-4e7b-bb07-c3c9d5b00555');

INSERT INTO customers (id, phone, gmail, language, date_of_birth, gender, created_at, updated_at, deleted_at)
VALUES 
    ('123e4567-e89b-12d3-a456-426614174000', '+1234567890', 'example@gmail.com', 'en', '1990-01-01', 'Female', NOW(), NULL, NULL),
    ('223e4567-e89b-12d3-a456-426614174001', '+1987654321', 'test@gmail.com', 'en', '1985-05-15', 'Female', NOW(), NULL, NULL),
    ('323e4567-e89b-12d3-a456-426614174002', '+1122334455', 'demo@gmail.com', 'uz', '1982-11-30', 'Female', NOW(), NULL, NULL);




INSERT INTO orders (id, external_id, type, customer_phone, customer_name, customer_id, payment_type, status, to_address, to_location, discount_amount, amount, delivery_price, paid, created_at, updated_at, deleted_at)
VALUES
    ('1f643b3d-85a8-4c3e-bd4a-3812bde9f548', 'EXT12345', 'self_pickup', '+1234567890', 'John Doe', '123e4567-e89b-12d3-a456-426614174000', 'cash', 'swaiting_for_payment', '123 Main St, City, Country', '((1.0, 1.0), (1.0, 2.0), (2.0, 2.0), (2.0, 1.0))', 10.5, 100.0, 5.0, false, NOW(), NOW(), NULL),
    ('2f643b3d-85a8-4c3e-bd4a-3812bde9f549', 'EXT23456', 'self_pickup', '+9876543210', 'Jane Smith', '123e4567-e89b-12d3-a456-426614174000', 'cash', 'swaiting_for_payment', '456 Elm St, City, Country', '((1.0, 1.0), (1.0, 2.0), (2.0, 2.0), (2.0, 1.0))', 0.0, 50.0, 7.0, true, NOW(), NOW(), NULL),
    ('3f643b3d-85a8-4c3e-bd4a-3812bde9f550', 'EXT34567', 'self_pickup', '+5554443333', 'Sam Green', '123e4567-e89b-12d3-a456-426614174000', 'cash', 'swaiting_for_payment', '789 Oak St, City, Country', '((1.0, 1.0), (1.0, 2.0), (2.0, 2.0), (2.0, 1.0))', 15.75, 200.0, 10.0, true, NOW(), NOW(), NULL);
