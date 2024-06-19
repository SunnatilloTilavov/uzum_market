DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'language_enum') THEN
    CREATE TYPE "language_enum" AS ENUM ('uz', 'ru', 'en');
  END IF;
END$$;

DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
    CREATE TYPE "gender_enum" AS ENUM ('Male', 'Female', 'Other');
  END IF;
END$$;

DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
    CREATE TYPE "role_enum" AS ENUM ('admin', 'courier');
  END IF;
END$$;

DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'currency_enum') THEN
    CREATE TYPE currency_enum AS ENUM ('usd', 'uzs');
  END IF;
END$$;

CREATE TABLE IF NOT EXISTS customers(
    id uuid PRIMARY KEY,
    phone varchar(20),
    gmail varchar(30),
    language language_enum,
    date_of_birth DATE,
    gender gender_enum,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp,
    UNIQUE(phone, gmail)
);

CREATE TABLE IF NOT EXISTS shop (
    id uuid PRIMARY KEY,
    slug varchar(20) UNIQUE,
    phone varchar(20) NOT NULL,
    name_uz varchar(20) DEFAULT '',
    name_ru varchar(20) DEFAULT '',
    name_en varchar(20) DEFAULT '',
    description_uz varchar(500) DEFAULT '',
    description_ru varchar(500) DEFAULT '',
    description_en varchar(500) DEFAULT '',
    location varchar,
    currency currency_enum,
    payment_types varchar[],
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS seller(
    id uuid PRIMARY KEY,
    phone varchar(20),
    gmail varchar(30),
    name varchar(20),
    date_of_birth DATE,
    shop_id uuid REFERENCES shop (id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp,
    UNIQUE(phone, gmail)
);

CREATE TABLE IF NOT EXISTS system_user(
    id uuid PRIMARY KEY,
    phone varchar(20),
    gmail varchar(30),
    name varchar(20),
    role role_enum,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp,
    UNIQUE(phone, gmail)
);

CREATE TABLE IF NOT EXISTS branch(
    id uuid PRIMARY KEY,
    phone varchar(20),
    name varchar(20),
    location polygon,
    addres varchar,
    open_time varchar,
    close_time varchar,
    active boolean NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp,
    UNIQUE(phone)
);

CREATE TABLE IF NOT EXISTS category (
    id uuid PRIMARY KEY,
    slug varchar(20) UNIQUE,
    name_uz varchar(20) DEFAULT '',
    name_ru varchar(20) DEFAULT '',
    name_en varchar(20) DEFAULT '',
    active boolean DEFAULT true,
    order_no integer DEFAULT 0,
    parent_id uuid REFERENCES category (id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS product (
    id uuid PRIMARY KEY,
    slug varchar(20) UNIQUE,
    name_uz varchar(20) DEFAULT '',
    name_ru varchar(20) DEFAULT '',
    name_en varchar(20) DEFAULT '',
    description_uz varchar(500) DEFAULT '',
    description_ru varchar(500) DEFAULT '',
    description_en varchar(500) DEFAULT '',
    active boolean DEFAULT true,
    order_no integer DEFAULT 0,
    in_price float,
    out_price float,
    left_count integer,
    discount_percent float DEFAULT 0,
    image varchar(200)[],
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS product_categories(
    id uuid PRIMARY KEY,
    product_id uuid REFERENCES product (id) NOT NULL,
    category_id uuid REFERENCES category (id) NOT NULL,
    UNIQUE (product_id, category_id)
);

CREATE TABLE IF NOT EXISTS product_reviews(
    id uuid PRIMARY KEY,
    customer_id uuid REFERENCES customers (id),
    product_id uuid REFERENCES product (id),
    text varchar(500),
    rating float,
    order_id uuid REFERENCES orders(id), 
    created_at timestamp DEFAULT NOW(),
    UNIQUE (product_id, category_id,order_id)
);
DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'type_enum') THEN
    CREATE TYPE type_enum AS ENUM ('self_pickup', 'delivery');
  END IF;
END$$;

DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN
    CREATE TYPE status_enum AS ENUM ('swaiting_for_payment', 'collecting', 'delivery', 'waiting_on_branch', 'finished', 'cancelled');
  END IF;
END$$;

DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_enum') THEN
    CREATE TYPE payment_enum AS ENUM ('uzum','cash','terminal');
  END IF;
END$$;

CREATE TABLE IF NOT EXISTS orders(
    id uuid PRIMARY KEY,
    external_id varchar,
    type type_enum,
    customer_phone varchar(20),
    customer_name varchar(20),
    customer_id uuid REFERENCES customers (id),
    payment_type payment_enum,
    status status_enum,
    to_address varchar,
    to_location polygon,
    discount_amount float,
    amount float,
    delivery_price float,
    paid boolean default false,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp,
    UNIQUE(customer_phone)
);


CREATE TABLE IF NOT EXISTS order_products(
    id uuid PRIMARY KEY,
    product_id uuid REFERENCES product (id),
    count integer,
    discount_price float,
    price float,
    order_id uuid REFERENCES orders (id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);


CREATE TABLE IF NOT EXISTS order_status_notes(
    id uuid PRIMARY KEY,
    order_id uuid REFERENCES orders (id),
    status status_enum,
    user_id uuid,
    reason varchar(100),
    created_at timestamp default NOW()
);