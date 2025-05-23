-- +goose Up
CREATE SCHEMA IF NOT EXISTS test;

CREATE TABLE IF NOT EXISTS test.test (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    test TEXT
);