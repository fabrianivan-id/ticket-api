CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    ticket_title VARCHAR(100) NOT NULL,
    ticket_msg TEXT NOT NULL,
    user_id INT NOT NULL,
    status VARCHAR(3) DEFAULT 'opn' CHECK (status IN ('opn', 'cld', 'asn')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Example Query
SELECT * FROM tickets WHERE status = 'opn' LIMIT 20 OFFSET 40;
