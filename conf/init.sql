-- create user
CREATE USER app WITH PASSWORD 'passwd';
-- create database
CREATE DATABASE dev OWNER app;
-- grant privileges
GRANT ALL PRIVILEGES ON DATABASE dev TO app;