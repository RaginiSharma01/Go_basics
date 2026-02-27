CREATE TYPE transaction_status AS ENUM('pending' , 'completed' , 'failed');

CREATE TABLE IF NOT EXISTS transactions(
    transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id INT NOT NULL,
    amount NUMERIC(12,2) NOT NULL CHECK(amount>0),
    currency CHAR NOT NULL,
    status transaction_status DEFAULT 'pending'
    created_at TIMESTAMPTZ DEFAULT NOW()
);





