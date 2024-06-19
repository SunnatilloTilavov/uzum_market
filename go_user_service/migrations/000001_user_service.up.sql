  DO $$ 
  BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'language_enum') THEN
      CREATE TYPE language_enum AS ENUM ('uz', 'ru', 'en');
    END IF;
  END$$;

  DO $$ 
  BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
      CREATE TYPE gender_enum AS ENUM ('male', 'female', 'other');
    END IF;
  END$$;

  DO $$ 
  BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
      CREATE TYPE role_enum AS ENUM ('admin', 'courier');
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
      UNIQUE(phone,gmail)
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
      location polygon,
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
      UNIQUE(phone,gmail)
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
      UNIQUE(phone,gmail)
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


  ALTER TABLE branch
  ALTER COLUMN open_time TYPE INTEGER USING open_time::integer;

  ALTER TABLE branch
  ALTER COLUMN close_time TYPE INTEGER USING close_time::integer;

  CREATE EXTENSION IF NOT EXISTS postgis;
  
  ALTER TABLE branch 
  ALTER COLUMN location 
  TYPE GEOMETRY(POINT, 4326) 
  USING ST_SetSRID(ST_Centroid(location::geometry), 4326);

  UPDATE branch
  SET location = ST_GeomFromText(location::text, 4326);

  ALTER TABLE shop 
  ALTER COLUMN location 
  TYPE GEOMETRY(POINT, 4326) 
  USING ST_SetSRID(ST_Centroid(location::geometry), 4326);

  UPDATE shop
  SET location = ST_GeomFromText(location::text, 4326);