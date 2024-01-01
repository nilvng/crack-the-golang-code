CREATE SCHEMA IF NOT EXISTS trips;

CREATE TABLE IF NOT EXISTS pool_trips (
  trip_id BIGINT PRIMARY KEY,
  driver_id VARCHAR(255) NOT NULL REFERENCES drivers(driver_id),
  rider_id VARCHAR(255) NOT NULL REFERENCES riders(rider_id),
  pickup VARCHAR(255),
  dropoff VARCHAR(255),
  distance_traveled DECIMAL(10,2),
  fare_amount DECIMAL(10,2),
  payment_method VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS drivers (
  driver_id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  rating DECIMAL(2,1),
  vehicle_type VARCHAR(255),
  number_plate VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS riders (
  rider_id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  rating DECIMAL(2,1)
);

CREATE INDEX idx_trips_driver_number_plate ON drivers (number_plate);

CREATE ROLE github WITH 
  LOGIN
  NOSUPERUSER
  NOCREATEDB
  NOINHERIT
  PASSWORD '123456';
