CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);


INSERT INTO items (name) 
VALUES 
('clothes'), ('skincare'), ('lotion'),('vege'),('meat'),('medicine'),('utinsile'),('bottle'),('blah'),('xyz')